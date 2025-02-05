/*
Copyright 2023 The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package minwait

import (
	"context"
	"fmt"
	"time"

	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/types"
	autoscaling "k8s.io/autoscaler/cluster-autoscaler/apis/provisioningrequest/autoscaling.x-k8s.io/v1beta1"
	"k8s.io/client-go/tools/record"
	"k8s.io/client-go/util/workqueue"
	"k8s.io/klog/v2"
	"k8s.io/utils/clock"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/event"
	"sigs.k8s.io/controller-runtime/pkg/handler"
	"sigs.k8s.io/controller-runtime/pkg/reconcile"

	kueue "sigs.k8s.io/kueue/apis/kueue/v1beta1"
	"sigs.k8s.io/kueue/pkg/util/admissioncheck"
	"sigs.k8s.io/kueue/pkg/util/api"
	"sigs.k8s.io/kueue/pkg/util/slices"
	"sigs.k8s.io/kueue/pkg/workload"
)

const (
	RequestsOwnedByWorkloadKey     = "metadata.ownedByWorkload"
	WorkloadsWithAdmissionCheckKey = "status.admissionChecks"
	AdmissionCheckUsingConfigKey   = "spec.provisioningRequestConfig"
)

var (
	realClock = clock.RealClock{}
)

type minWaitConfigHelper = admissioncheck.ConfigHelper[*kueue.MinWaitConfig, kueue.MinWaitConfig]

func newProvisioningConfigHelper(c client.Client) (*minWaitConfigHelper, error) {
	return admissioncheck.NewConfigHelper[*kueue.MinWaitConfig](c)
}

type Controller struct {
	client client.Client
	record record.EventRecorder
	helper *minWaitConfigHelper
	clock  clock.Clock
}

type workloadInfo struct {
	checkStates  []kueue.AdmissionCheckState
	requeueState *kueue.RequeueState
}

var _ reconcile.Reconciler = (*Controller)(nil)

// +kubebuilder:rbac:groups="",resources=events,verbs=create;watch;update
// +kubebuilder:rbac:groups=kueue.x-k8s.io,resources=workloads,verbs=get;list;watch;update;patch;delete
// +kubebuilder:rbac:groups=kueue.x-k8s.io,resources=workloads/status,verbs=get;update;patch
// +kubebuilder:rbac:groups=kueue.x-k8s.io,resources=admissionchecks,verbs=get;list;watch

func NewController(client client.Client, record record.EventRecorder) (*Controller, error) {
	helper, err := newProvisioningConfigHelper(client)
	if err != nil {
		return nil, err
	}
	return &Controller{
		client: client,
		record: record,
		helper: helper,
		clock:  realClock,
	}, nil
}

// Reconcile performs a full reconciliation for the object referred to by the Request.
// The Controller will requeue the Request to be processed again if an error is non-nil or
// Result.Requeue is true, otherwise upon completion it will remove the work from the queue.
func (c *Controller) Reconcile(ctx context.Context, req reconcile.Request) (reconcile.Result, error) {
	wl := &kueue.Workload{}
	err := c.client.Get(ctx, req.NamespacedName, wl)
	if err != nil {
		return reconcile.Result{}, client.IgnoreNotFound(err)
	}
	if !workload.HasQuotaReservation(wl) || workload.IsFinished(wl) || workload.IsEvicted(wl) {
		return reconcile.Result{}, nil
	}

	provisioningRequestList := &autoscaling.ProvisioningRequestList{}
	if err := c.client.List(ctx, provisioningRequestList, client.InNamespace(wl.Namespace), client.MatchingFields{RequestsOwnedByWorkloadKey: wl.Name}); client.IgnoreNotFound(err) != nil {
		return reconcile.Result{}, err
	}

	// get the lists of relevant checks
	relevantChecks, err := admissioncheck.FilterForController(ctx, c.client, wl.Status.AdmissionChecks, kueue.MinWaitControllerName)
	if err != nil {
		return reconcile.Result{}, err
	}

	checkConfig := make(map[string]*kueue.MinWaitConfig, len(relevantChecks))
	for _, checkName := range relevantChecks {
		prc, err := c.helper.ConfigForAdmissionCheck(ctx, checkName)
		if client.IgnoreNotFound(err) != nil {
			return reconcile.Result{}, err
		}
		checkConfig[checkName] = prc
	}

	waitTime := int(time.Since(wl.CreationTimestamp.Time).Seconds())

	minWait := 0
	for _, check := range checkConfig {
		if waitTime < check.Spec.TimeSeconds {
			minWait = max(minWait, check.Spec.TimeSeconds-waitTime)
		}
	}

	wlInfo := workloadInfo{
		checkStates: make([]kueue.AdmissionCheckState, 0),
	}
	err = c.syncCheckStates(ctx, wl, &wlInfo, checkConfig, waitTime)
	if err != nil {
		return reconcile.Result{}, err
	}

	if minWait > 0 {
		return reconcile.Result{RequeueAfter: time.Duration(minWait)}, nil
	}
	return reconcile.Result{}, nil
}

func updateCheckMessage(checkState *kueue.AdmissionCheckState, message string) bool {
	if message == "" || checkState.Message == message {
		return false
	}
	checkState.Message = message
	return true
}

func updateCheckState(checkState *kueue.AdmissionCheckState, state kueue.CheckState) bool {
	if checkState.State == state {
		return false
	}
	checkState.State = state
	return true
}

func (wlInfo *workloadInfo) update(wl *kueue.Workload, c clock.Clock) {
	for _, check := range wl.Status.AdmissionChecks {
		workload.SetAdmissionCheckState(&wlInfo.checkStates, check, c)
	}
	wlInfo.requeueState = wl.Status.RequeueState
}

func (c *Controller) syncCheckStates(
	ctx context.Context, wl *kueue.Workload,
	wlInfo *workloadInfo,
	checkConfig map[string]*kueue.MinWaitConfig,
	waitTime int,
) error {
	wlInfo.update(wl, c.clock)
	checksMap := slices.ToRefMap(wl.Status.AdmissionChecks, func(c *kueue.AdmissionCheckState) string { return c.Name })
	wlPatch := workload.BaseSSAWorkload(wl)
	recorderMessages := make([]string, 0, len(checkConfig))
	updated := false
	for check, prc := range checkConfig {
		checkState := *checksMap[check]
		//nolint:gocritic
		if prc == nil {
			// the check is not active
			updated = updateCheckState(&checkState, kueue.CheckStatePending) || updated
			updated = updateCheckMessage(&checkState, CheckInactiveMessage) || updated
		} else if waitTime >= prc.Spec.TimeSeconds {
			if updateCheckState(&checkState, kueue.CheckStateReady) {
				updated = true
				checkState.Message = "Min wait time exceeded"
			}
		} else {
			if updateCheckState(&checkState, kueue.CheckStateRetry) {
				updated = true
				checkState.Message = "Min wait time not exceed will retry"
			}
		}

		existingCondition := workload.FindAdmissionCheck(wlPatch.Status.AdmissionChecks, checkState.Name)
		if existingCondition != nil && existingCondition.State != checkState.State {
			message := fmt.Sprintf("Admission check %s updated state from %s to %s", checkState.Name, existingCondition.State, checkState.State)
			if checkState.Message != "" {
				message += fmt.Sprintf(" with message %s", checkState.Message)
			}
			recorderMessages = append(recorderMessages, message)
		}
		workload.SetAdmissionCheckState(&wlPatch.Status.AdmissionChecks, checkState, c.clock)
	}
	if updated {
		if err := c.client.Status().Patch(ctx, wlPatch, client.Apply, client.FieldOwner(kueue.ProvisioningRequestControllerName), client.ForceOwnership); err != nil {
			return err
		}
		for i := range recorderMessages {
			c.record.Event(wl, corev1.EventTypeNormal, "AdmissionCheckUpdated", api.TruncateEventMessage(recorderMessages[i]))
		}
	}
	wlInfo.update(wlPatch, c.clock)
	return nil
}

type acHandler struct {
	client client.Client
}

var _ handler.EventHandler = (*acHandler)(nil)

func (a *acHandler) Create(ctx context.Context, event event.CreateEvent, q workqueue.TypedRateLimitingInterface[reconcile.Request]) {
	ac, isAc := event.Object.(*kueue.AdmissionCheck)
	if !isAc {
		return
	}

	if ac.Spec.ControllerName == kueue.ProvisioningRequestControllerName {
		err := a.reconcileWorkloadsUsing(ctx, ac.Name, q)
		if err != nil {
			ctrl.LoggerFrom(ctx).V(5).Error(err, "Failure on create event", "admissionCheck", klog.KObj(ac))
		}
	}
}

func (a *acHandler) Update(ctx context.Context, event event.UpdateEvent, q workqueue.TypedRateLimitingInterface[reconcile.Request]) {
	oldAc, isOldAc := event.ObjectOld.(*kueue.AdmissionCheck)
	newAc, isNewAc := event.ObjectNew.(*kueue.AdmissionCheck)
	if !isNewAc || !isOldAc {
		return
	}

	if oldAc.Spec.ControllerName == kueue.ProvisioningRequestControllerName || newAc.Spec.ControllerName == kueue.ProvisioningRequestControllerName {
		err := a.reconcileWorkloadsUsing(ctx, oldAc.Name, q)
		if err != nil {
			ctrl.LoggerFrom(ctx).V(5).Error(err, "Failure on update event", "admissionCheck", klog.KObj(oldAc))
		}
	}
}

func (a *acHandler) Delete(ctx context.Context, event event.DeleteEvent, q workqueue.TypedRateLimitingInterface[reconcile.Request]) {
	ac, isAc := event.Object.(*kueue.AdmissionCheck)
	if !isAc {
		return
	}

	if ac.Spec.ControllerName == kueue.MinWaitControllerName {
		err := a.reconcileWorkloadsUsing(ctx, ac.Name, q)
		if err != nil {
			ctrl.LoggerFrom(ctx).V(5).Error(err, "Failure on delete event", "admissionCheck", klog.KObj(ac))
		}
	}
}

func (a *acHandler) Generic(_ context.Context, _ event.GenericEvent, _ workqueue.TypedRateLimitingInterface[reconcile.Request]) {
	// nothing to do for now
}

func (a *acHandler) reconcileWorkloadsUsing(ctx context.Context, check string, q workqueue.TypedRateLimitingInterface[reconcile.Request]) error {
	list := &kueue.WorkloadList{}
	if err := a.client.List(ctx, list, client.MatchingFields{WorkloadsWithAdmissionCheckKey: check}); client.IgnoreNotFound(err) != nil {
		return err
	}

	for i := range list.Items {
		wl := &list.Items[i]
		req := reconcile.Request{
			NamespacedName: types.NamespacedName{
				Name:      wl.Name,
				Namespace: wl.Namespace,
			},
		}
		q.Add(req)
	}

	return nil
}

type mwcHandler struct {
	client            client.Client
	acHandlerOverride func(ctx context.Context, config string, q workqueue.TypedRateLimitingInterface[reconcile.Request]) error
}

var _ handler.EventHandler = (*mwcHandler)(nil)

func (h *mwcHandler) Create(ctx context.Context, event event.CreateEvent, q workqueue.TypedRateLimitingInterface[reconcile.Request]) {
	mwc, isMWC := event.Object.(*kueue.MinWaitConfig)
	if !isMWC {
		return
	}
	err := h.reconcileWorkloadsUsing(ctx, mwc.Name, q)
	if err != nil {
		ctrl.LoggerFrom(ctx).V(5).Error(err, "Failure on create event", "provisioningRequestConfig", klog.KObj(mwc))
	}
}

func (h *mwcHandler) Update(ctx context.Context, event event.UpdateEvent, q workqueue.TypedRateLimitingInterface[reconcile.Request]) {
	oldMWC, isOldMWC := event.ObjectOld.(*kueue.MinWaitConfig)
	newMWC, isNewMWC := event.ObjectNew.(*kueue.MinWaitConfig)
	if !isNewMWC || !isOldMWC {
		return
	}

	if oldMWC.Spec.TimeSeconds < newMWC.Spec.TimeSeconds {
		err := h.reconcileWorkloadsUsing(ctx, oldMWC.Name, q)
		if err != nil {
			ctrl.LoggerFrom(ctx).V(5).Error(err, "Failure on update event", "provisioningRequestConfig", klog.KObj(oldMWC))
		}
	}
}

func (h *mwcHandler) Delete(ctx context.Context, event event.DeleteEvent, q workqueue.TypedRateLimitingInterface[reconcile.Request]) {
	mwc, isMWC := event.Object.(*kueue.MinWaitConfig)
	if !isMWC {
		return
	}
	err := h.reconcileWorkloadsUsing(ctx, mwc.Name, q)
	if err != nil {
		ctrl.LoggerFrom(ctx).V(5).Error(err, "Failure on delete event", "provisioningRequestConfig", klog.KObj(mwc))
	}
}

func (h *mwcHandler) Generic(_ context.Context, _ event.GenericEvent, _ workqueue.TypedRateLimitingInterface[reconcile.Request]) {
	// nothing to do for now
}

func (h *mwcHandler) reconcileWorkloadsUsing(ctx context.Context, config string, q workqueue.TypedRateLimitingInterface[reconcile.Request]) error {
	list := &kueue.AdmissionCheckList{}
	if err := h.client.List(ctx, list, client.MatchingFields{AdmissionCheckUsingConfigKey: config}); client.IgnoreNotFound(err) != nil {
		return err
	}
	users := slices.Map(list.Items, func(ac *kueue.AdmissionCheck) string { return ac.Name })
	for _, user := range users {
		if h.acHandlerOverride != nil {
			if err := h.acHandlerOverride(ctx, user, q); err != nil {
				return err
			}
		} else {
			req := reconcile.Request{
				NamespacedName: types.NamespacedName{
					Name: user,
				},
			}
			q.Add(req)
		}
	}
	return nil
}

func (c *Controller) SetupWithManager(mgr ctrl.Manager) error {
	ach := &acHandler{
		client: c.client,
	}
	prch := &mwcHandler{
		client:            c.client,
		acHandlerOverride: ach.reconcileWorkloadsUsing,
	}
	err := ctrl.NewControllerManagedBy(mgr).
		Named("minwait-workload").
		For(&kueue.Workload{}).
		Owns(&autoscaling.ProvisioningRequest{}).
		Watches(&kueue.AdmissionCheck{}, ach).
		Watches(&kueue.ProvisioningRequestConfig{}, prch).
		Complete(c)
	if err != nil {
		return err
	}

	mwcACh := &mwcHandler{
		client: c.client,
	}
	acReconciler := &acReconciler{
		client: c.client,
		helper: c.helper,
	}

	return ctrl.NewControllerManagedBy(mgr).
		Named("minwait-admissioncheck").
		For(&kueue.AdmissionCheck{}).
		Watches(&kueue.MinWaitConfig{}, mwcACh).
		Complete(acReconciler)
}
