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

// TpccBenchmarkSpec defines the desired state of TpccBenchmark
type TpccBenchmarkSpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "make" to regenerate code after modifying this file

	// Conn describe the mysql host connection manually
	// +optional
	Conn        *string `json:"conn,omitempty"`
	Warehouses  uint32  `json:"warehouses"`
	Terminals   uint32  `json:"terminals"`
	LoadWorkers uint32  `json:"loadworkers"`

	// Cluster describe the TidbCluster Ref
	// +optional
	Cluster TidbClusterRef `json:"cluster,omitempty"`

	// Database describe the Target Database
	// +optional
	Database *string `json:"database,omitempty"`

	// Username describe the Username to connect the database
	// If not set, the default is root
	// +optional
	Username string `json:"user,omitempty"`

	// Password describe the password to connect the database
	// If not set, the default is empty
	Password string `json:"password,omitempty"`
}

// TpccBenchmarkStatus defines the observed state of TpccBenchmark
type TpccBenchmarkStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "make" to regenerate code after modifying this file
}

// +kubebuilder:object:root=true

// TpccBenchmark is the Schema for the tpccbenchmarks API
type TpccBenchmark struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   TpccBenchmarkSpec   `json:"spec,omitempty"`
	Status TpccBenchmarkStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// TpccBenchmarkList contains a list of TpccBenchmark
type TpccBenchmarkList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []TpccBenchmark `json:"items"`
}

// TidbClusterRef describe the target location of tidbcluster
type TidbClusterRef struct {
	// Namespace is the namespace that TidbCluster object locates,
	// default to the same namespace where the obj created
	// +optional
	Namespace string `json:"namespace,omitempty"`

	// Name is the name of TidbCluster object
	Name string `json:"name"`
}

func init() {
	SchemeBuilder.Register(&TpccBenchmark{}, &TpccBenchmarkList{})
}
