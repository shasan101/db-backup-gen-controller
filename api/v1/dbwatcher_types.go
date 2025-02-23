/*
Copyright 2025.

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

package v1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

type DbSpecs struct {
	Name             string `json:"Name,omitempty"`
	ConnectionString string `json:"ConnectionString,omitempty"`
}

// DbWatcherSpec defines the desired state of DbWatcher.
type DbWatcherSpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "make" to regenerate code after modifying this file

	DbSpec         DbSpecs `json:"DbSpec,omitempty"`
	CronExpression string  `json:"CronExpression,omitempty"`
}

// DbWatcherStatus defines the observed state of DbWatcher.
type DbWatcherStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "make" to regenerate code after modifying this file
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status

// DbWatcher is the Schema for the dbwatchers API.
type DbWatcher struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   DbWatcherSpec   `json:"spec,omitempty"`
	Status DbWatcherStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// DbWatcherList contains a list of DbWatcher.
type DbWatcherList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []DbWatcher `json:"items"`
}

func init() {
	SchemeBuilder.Register(&DbWatcher{}, &DbWatcherList{})
}
