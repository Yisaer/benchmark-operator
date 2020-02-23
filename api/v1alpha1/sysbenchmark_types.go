/*

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

package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// SysBenchmarkSpec defines the desired state of SysBenchmark
type SysBenchmarkSpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "make" to regenerate code after modifying this file
	Host           string `json:"host"`
	Port           uint16 `json:"port"`
	Username       string `json:"user,omitempty"`
	Password       string `json:"password,omitempty"`
	TableSize      uint32 `json:"tablesize"`
	TablesCount    uint32 `json:"tablescount"`
	Threads        uint32 `json:"threads"`
	Time           uint32 `json:"time"`
	ReportInterval uint32 `json:"reportinterval"`
}

// SysBenchmarkStatus defines the observed state of SysBenchmark
type SysBenchmarkStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "make" to regenerate code after modifying this file
}

// +kubebuilder:object:root=true

// SysBenchmark is the Schema for the sysbenchmarks API
type SysBenchmark struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   SysBenchmarkSpec   `json:"spec,omitempty"`
	Status SysBenchmarkStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// SysBenchmarkList contains a list of SysBenchmark
type SysBenchmarkList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []SysBenchmark `json:"items"`
}

func init() {
	SchemeBuilder.Register(&SysBenchmark{}, &SysBenchmarkList{})
}
