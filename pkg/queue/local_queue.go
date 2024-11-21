/*
Copyright 2022 The Kubernetes Authors.

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

package queue

import (
	"fmt"

	kueue "sigs.k8s.io/kueue/apis/kueue/v1beta1"
	"sigs.k8s.io/kueue/pkg/workload"
)

type localQueueOption func(*LocalQueue)

// WithMetricsEnabled provides an option to enable metrics for this LocalQueue.
func withMetricsEnabled() localQueueOption {
	return func(lq *LocalQueue) {
		lq.enableMetrics = true
	}
}

// Key is the key used to index the queue.
func Key(q *kueue.LocalQueue) string {
	return fmt.Sprintf("%s/%s", q.Namespace, q.Name)
}

// LocalQueue is the internal implementation of kueue.LocalQueue.
type LocalQueue struct {
	Key           string
	ClusterQueue  string
	enableMetrics bool

	items map[string]*workload.Info
}

func newLocalQueue(q *kueue.LocalQueue, opts ...localQueueOption) *LocalQueue {
	qImpl := &LocalQueue{
		Key:   Key(q),
		items: make(map[string]*workload.Info),
	}
	for _, opt := range opts {
		opt(qImpl)
	}
	qImpl.update(q)
	return qImpl
}

func (q *LocalQueue) update(apiQueue *kueue.LocalQueue) {
	q.ClusterQueue = string(apiQueue.Spec.ClusterQueue)
}

func (q *LocalQueue) AddOrUpdate(info *workload.Info) {
	key := workload.Key(info.Obj)
	q.items[key] = info
}
