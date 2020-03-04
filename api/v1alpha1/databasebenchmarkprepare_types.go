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

type Prepare struct {
	Type     string            `json:"type"`
	Database string            `json:"database"`
	Image    string            `json:"image"`
	Params   map[string]string `json:"params"`
}

// DataBaseBenchmarkPrepareSpec defines the desired state of DataBaseBenchmarkPrepare
type DataBaseBenchmarkPrepareSpec struct {
	Host     string    `json:"host"`
	Port     uint16    `json:"port"`
	User     string    `json:"user"`
	Password string    `json:"password"`
	Prepares []Prepare `json:"prepares"`
}

// DataBaseBenchmarkPrepareStatus defines the observed state of DataBaseBenchmarkPrepare
type DataBaseBenchmarkPrepareStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "make" to regenerate code after modifying this file
}

// +kubebuilder:object:root=true

// DataBaseBenchmarkPrepare is the Schema for the databasebenchmarkprepares API
type DataBaseBenchmarkPrepare struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   DataBaseBenchmarkPrepareSpec   `json:"spec,omitempty"`
	Status DataBaseBenchmarkPrepareStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// DataBaseBenchmarkPrepareList contains a list of DataBaseBenchmarkPrepare
type DataBaseBenchmarkPrepareList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []DataBaseBenchmarkPrepare `json:"items"`
}

func init() {
	SchemeBuilder.Register(&DataBaseBenchmarkPrepare{}, &DataBaseBenchmarkPrepareList{})
}
