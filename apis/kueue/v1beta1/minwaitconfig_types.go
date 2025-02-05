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

package v1beta1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

const (
	// MinWaitControllerName is the name used by the Provisioning
	// Request admission check controller.
	MinWaitControllerName = "kueue.x-k8s.io/minwait"
)

// MinWaitConfigSpec defines the desired state of ProvisioningRequestConfig
type MinWaitConfigSpec struct {
	// Time in seconds before the admission check will pass
	TimeSeconds int `json:"parameters,omitempty"`
}

// +genclient
// +genclient:nonNamespaced
// +kubebuilder:object:root=true
// +kubebuilder:storageversion
// +kubebuilder:resource:scope=Cluster

// MinWaitConfig is the Schema for the provisioningrequestconfig API
type MinWaitConfig struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec MinWaitConfigSpec `json:"spec,omitempty"`
}

// +kubebuilder:object:root=true

// MinWaitConfigList contains a list of ProvisioningRequestConfig
type MinWaitConfigList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []MinWaitConfig `json:"items"`
}

func init() {
	// SchemeBuilder.Register(&MinWaitConfig{}, &MinWaitConfigList{})
}
