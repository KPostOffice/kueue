//go:build !ignore_autogenerated

/*
Copyright The Kubernetes Authors.

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

// Code generated by controller-gen. DO NOT EDIT.

package v1beta1

import (
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	"k8s.io/component-base/config/v1alpha1"
	timex "time"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClientConnection) DeepCopyInto(out *ClientConnection) {
	*out = *in
	if in.QPS != nil {
		in, out := &in.QPS, &out.QPS
		*out = new(float32)
		**out = **in
	}
	if in.Burst != nil {
		in, out := &in.Burst, &out.Burst
		*out = new(int32)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClientConnection.
func (in *ClientConnection) DeepCopy() *ClientConnection {
	if in == nil {
		return nil
	}
	out := new(ClientConnection)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterQueueVisibility) DeepCopyInto(out *ClusterQueueVisibility) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterQueueVisibility.
func (in *ClusterQueueVisibility) DeepCopy() *ClusterQueueVisibility {
	if in == nil {
		return nil
	}
	out := new(ClusterQueueVisibility)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Configuration) DeepCopyInto(out *Configuration) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	if in.Namespace != nil {
		in, out := &in.Namespace, &out.Namespace
		*out = new(string)
		**out = **in
	}
	in.ControllerManager.DeepCopyInto(&out.ControllerManager)
	if in.ManagedJobsNamespaceSelector != nil {
		in, out := &in.ManagedJobsNamespaceSelector, &out.ManagedJobsNamespaceSelector
		*out = new(v1.LabelSelector)
		(*in).DeepCopyInto(*out)
	}
	if in.InternalCertManagement != nil {
		in, out := &in.InternalCertManagement, &out.InternalCertManagement
		*out = new(InternalCertManagement)
		(*in).DeepCopyInto(*out)
	}
	if in.WaitForPodsReady != nil {
		in, out := &in.WaitForPodsReady, &out.WaitForPodsReady
		*out = new(WaitForPodsReady)
		(*in).DeepCopyInto(*out)
	}
	if in.ClientConnection != nil {
		in, out := &in.ClientConnection, &out.ClientConnection
		*out = new(ClientConnection)
		(*in).DeepCopyInto(*out)
	}
	if in.Integrations != nil {
		in, out := &in.Integrations, &out.Integrations
		*out = new(Integrations)
		(*in).DeepCopyInto(*out)
	}
	if in.QueueVisibility != nil {
		in, out := &in.QueueVisibility, &out.QueueVisibility
		*out = new(QueueVisibility)
		(*in).DeepCopyInto(*out)
	}
	if in.MultiKueue != nil {
		in, out := &in.MultiKueue, &out.MultiKueue
		*out = new(MultiKueue)
		(*in).DeepCopyInto(*out)
	}
	if in.FairSharing != nil {
		in, out := &in.FairSharing, &out.FairSharing
		*out = new(FairSharing)
		(*in).DeepCopyInto(*out)
	}
	if in.Resources != nil {
		in, out := &in.Resources, &out.Resources
		*out = new(Resources)
		(*in).DeepCopyInto(*out)
	}
	if in.FeatureGates != nil {
		in, out := &in.FeatureGates, &out.FeatureGates
		*out = make(map[string]bool, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Configuration.
func (in *Configuration) DeepCopy() *Configuration {
	if in == nil {
		return nil
	}
	out := new(Configuration)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Configuration) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ControllerConfigurationSpec) DeepCopyInto(out *ControllerConfigurationSpec) {
	*out = *in
	if in.GroupKindConcurrency != nil {
		in, out := &in.GroupKindConcurrency, &out.GroupKindConcurrency
		*out = make(map[string]int, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.CacheSyncTimeout != nil {
		in, out := &in.CacheSyncTimeout, &out.CacheSyncTimeout
		*out = new(timex.Duration)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ControllerConfigurationSpec.
func (in *ControllerConfigurationSpec) DeepCopy() *ControllerConfigurationSpec {
	if in == nil {
		return nil
	}
	out := new(ControllerConfigurationSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ControllerHealth) DeepCopyInto(out *ControllerHealth) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ControllerHealth.
func (in *ControllerHealth) DeepCopy() *ControllerHealth {
	if in == nil {
		return nil
	}
	out := new(ControllerHealth)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ControllerManager) DeepCopyInto(out *ControllerManager) {
	*out = *in
	in.Webhook.DeepCopyInto(&out.Webhook)
	if in.LeaderElection != nil {
		in, out := &in.LeaderElection, &out.LeaderElection
		*out = new(v1alpha1.LeaderElectionConfiguration)
		(*in).DeepCopyInto(*out)
	}
	in.Metrics.DeepCopyInto(&out.Metrics)
	out.Health = in.Health
	if in.Controller != nil {
		in, out := &in.Controller, &out.Controller
		*out = new(ControllerConfigurationSpec)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ControllerManager.
func (in *ControllerManager) DeepCopy() *ControllerManager {
	if in == nil {
		return nil
	}
	out := new(ControllerManager)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ControllerMetrics) DeepCopyInto(out *ControllerMetrics) {
	*out = *in
	if in.LocalQueueMetrics != nil {
		in, out := &in.LocalQueueMetrics, &out.LocalQueueMetrics
		*out = new(LocalQueueMetrics)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ControllerMetrics.
func (in *ControllerMetrics) DeepCopy() *ControllerMetrics {
	if in == nil {
		return nil
	}
	out := new(ControllerMetrics)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ControllerWebhook) DeepCopyInto(out *ControllerWebhook) {
	*out = *in
	if in.Port != nil {
		in, out := &in.Port, &out.Port
		*out = new(int)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ControllerWebhook.
func (in *ControllerWebhook) DeepCopy() *ControllerWebhook {
	if in == nil {
		return nil
	}
	out := new(ControllerWebhook)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FairSharing) DeepCopyInto(out *FairSharing) {
	*out = *in
	if in.PreemptionStrategies != nil {
		in, out := &in.PreemptionStrategies, &out.PreemptionStrategies
		*out = make([]PreemptionStrategy, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FairSharing.
func (in *FairSharing) DeepCopy() *FairSharing {
	if in == nil {
		return nil
	}
	out := new(FairSharing)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Integrations) DeepCopyInto(out *Integrations) {
	*out = *in
	if in.Frameworks != nil {
		in, out := &in.Frameworks, &out.Frameworks
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.ExternalFrameworks != nil {
		in, out := &in.ExternalFrameworks, &out.ExternalFrameworks
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.PodOptions != nil {
		in, out := &in.PodOptions, &out.PodOptions
		*out = new(PodIntegrationOptions)
		(*in).DeepCopyInto(*out)
	}
	if in.LabelKeysToCopy != nil {
		in, out := &in.LabelKeysToCopy, &out.LabelKeysToCopy
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Integrations.
func (in *Integrations) DeepCopy() *Integrations {
	if in == nil {
		return nil
	}
	out := new(Integrations)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *InternalCertManagement) DeepCopyInto(out *InternalCertManagement) {
	*out = *in
	if in.Enable != nil {
		in, out := &in.Enable, &out.Enable
		*out = new(bool)
		**out = **in
	}
	if in.WebhookServiceName != nil {
		in, out := &in.WebhookServiceName, &out.WebhookServiceName
		*out = new(string)
		**out = **in
	}
	if in.WebhookSecretName != nil {
		in, out := &in.WebhookSecretName, &out.WebhookSecretName
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new InternalCertManagement.
func (in *InternalCertManagement) DeepCopy() *InternalCertManagement {
	if in == nil {
		return nil
	}
	out := new(InternalCertManagement)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LocalQueueMetrics) DeepCopyInto(out *LocalQueueMetrics) {
	*out = *in
	if in.Enabled != nil {
		in, out := &in.Enabled, &out.Enabled
		*out = new(bool)
		**out = **in
	}
	if in.LocalQueueSelector != nil {
		in, out := &in.LocalQueueSelector, &out.LocalQueueSelector
		*out = new(v1.LabelSelector)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LocalQueueMetrics.
func (in *LocalQueueMetrics) DeepCopy() *LocalQueueMetrics {
	if in == nil {
		return nil
	}
	out := new(LocalQueueMetrics)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MultiKueue) DeepCopyInto(out *MultiKueue) {
	*out = *in
	if in.GCInterval != nil {
		in, out := &in.GCInterval, &out.GCInterval
		*out = new(v1.Duration)
		**out = **in
	}
	if in.Origin != nil {
		in, out := &in.Origin, &out.Origin
		*out = new(string)
		**out = **in
	}
	if in.WorkerLostTimeout != nil {
		in, out := &in.WorkerLostTimeout, &out.WorkerLostTimeout
		*out = new(v1.Duration)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MultiKueue.
func (in *MultiKueue) DeepCopy() *MultiKueue {
	if in == nil {
		return nil
	}
	out := new(MultiKueue)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PodIntegrationOptions) DeepCopyInto(out *PodIntegrationOptions) {
	*out = *in
	if in.NamespaceSelector != nil {
		in, out := &in.NamespaceSelector, &out.NamespaceSelector
		*out = new(v1.LabelSelector)
		(*in).DeepCopyInto(*out)
	}
	if in.PodSelector != nil {
		in, out := &in.PodSelector, &out.PodSelector
		*out = new(v1.LabelSelector)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PodIntegrationOptions.
func (in *PodIntegrationOptions) DeepCopy() *PodIntegrationOptions {
	if in == nil {
		return nil
	}
	out := new(PodIntegrationOptions)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *QueueVisibility) DeepCopyInto(out *QueueVisibility) {
	*out = *in
	if in.ClusterQueues != nil {
		in, out := &in.ClusterQueues, &out.ClusterQueues
		*out = new(ClusterQueueVisibility)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new QueueVisibility.
func (in *QueueVisibility) DeepCopy() *QueueVisibility {
	if in == nil {
		return nil
	}
	out := new(QueueVisibility)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RequeuingStrategy) DeepCopyInto(out *RequeuingStrategy) {
	*out = *in
	if in.Timestamp != nil {
		in, out := &in.Timestamp, &out.Timestamp
		*out = new(RequeuingTimestamp)
		**out = **in
	}
	if in.BackoffLimitCount != nil {
		in, out := &in.BackoffLimitCount, &out.BackoffLimitCount
		*out = new(int32)
		**out = **in
	}
	if in.BackoffBaseSeconds != nil {
		in, out := &in.BackoffBaseSeconds, &out.BackoffBaseSeconds
		*out = new(int32)
		**out = **in
	}
	if in.BackoffMaxSeconds != nil {
		in, out := &in.BackoffMaxSeconds, &out.BackoffMaxSeconds
		*out = new(int32)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RequeuingStrategy.
func (in *RequeuingStrategy) DeepCopy() *RequeuingStrategy {
	if in == nil {
		return nil
	}
	out := new(RequeuingStrategy)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ResourceTransformation) DeepCopyInto(out *ResourceTransformation) {
	*out = *in
	if in.Strategy != nil {
		in, out := &in.Strategy, &out.Strategy
		*out = new(ResourceTransformationStrategy)
		**out = **in
	}
	if in.Outputs != nil {
		in, out := &in.Outputs, &out.Outputs
		*out = make(corev1.ResourceList, len(*in))
		for key, val := range *in {
			(*out)[key] = val.DeepCopy()
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ResourceTransformation.
func (in *ResourceTransformation) DeepCopy() *ResourceTransformation {
	if in == nil {
		return nil
	}
	out := new(ResourceTransformation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Resources) DeepCopyInto(out *Resources) {
	*out = *in
	if in.ExcludeResourcePrefixes != nil {
		in, out := &in.ExcludeResourcePrefixes, &out.ExcludeResourcePrefixes
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.Transformations != nil {
		in, out := &in.Transformations, &out.Transformations
		*out = make([]ResourceTransformation, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Resources.
func (in *Resources) DeepCopy() *Resources {
	if in == nil {
		return nil
	}
	out := new(Resources)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *WaitForPodsReady) DeepCopyInto(out *WaitForPodsReady) {
	*out = *in
	if in.Timeout != nil {
		in, out := &in.Timeout, &out.Timeout
		*out = new(v1.Duration)
		**out = **in
	}
	if in.BlockAdmission != nil {
		in, out := &in.BlockAdmission, &out.BlockAdmission
		*out = new(bool)
		**out = **in
	}
	if in.RequeuingStrategy != nil {
		in, out := &in.RequeuingStrategy, &out.RequeuingStrategy
		*out = new(RequeuingStrategy)
		(*in).DeepCopyInto(*out)
	}
	if in.RecoveryTimeout != nil {
		in, out := &in.RecoveryTimeout, &out.RecoveryTimeout
		*out = new(v1.Duration)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new WaitForPodsReady.
func (in *WaitForPodsReady) DeepCopy() *WaitForPodsReady {
	if in == nil {
		return nil
	}
	out := new(WaitForPodsReady)
	in.DeepCopyInto(out)
	return out
}
