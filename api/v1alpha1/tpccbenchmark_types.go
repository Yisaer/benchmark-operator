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
	batchv1 "k8s.io/api/batch/v1"
	"k8s.io/api/core/v1"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"strconv"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// TpccBenchmarkSpec defines the desired state of TpccBenchmark
type TpccBenchmarkSpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "make" to regenerate code after modifying this file
	Image       string `json:"image,omitempty"`
	Conn        string `json:"conn,omitempty"`
	User        string `json:"user,omitempty"`
	Password    string `json:"password,omitempty"`
	Warehouses  uint32 `json:"warehouses"`
	Terminals   uint32 `json:"terminals"`
	LoadWorkers uint32 `json:"loadworkers"`
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

func (in *TpccBenchmark) CreateJob() (*batchv1.Job, error) {
	ttl := int32(60)
	result := batchv1.Job{
		ObjectMeta: metav1.ObjectMeta{
			Name:      in.Name,
			Namespace: in.Namespace,
			// todo: add labels and annotation support
			Labels:      nil,
			Annotations: nil,
		},
		Spec: batchv1.JobSpec{
			// todo: add selector support
			Selector: &metav1.LabelSelector{
				MatchLabels:      nil,
				MatchExpressions: nil,
			},
			ManualSelector: nil,
			Template: corev1.PodTemplateSpec{
				Spec: corev1.PodSpec{
					Containers: []v1.Container{{
						Name:  "benchmarksql",
						Image: in.Spec.Image,
						// todo: check whether empty values are handled here
						Env: []v1.EnvVar{
							{Name: "CONN", Value: in.Spec.Conn},
							{Name: "USER", Value: in.Spec.User},
							{Name: "PASSWORD", Value: in.Spec.Password},
						},
						// todo: use this to sync benchmark result?
						//Lifecycle: &corev1.Lifecycle{
						//	PreStop: &corev1.Handler{},
						//},
						// todo: or this?
						TerminationMessagePath:   "",
						TerminationMessagePolicy: "",
					}},
					RestartPolicy: "OnFailure",
				},
			},
			// I don't know why k8s use *int32 instead of plain int32 here
			TTLSecondsAfterFinished: &ttl,
		},
	}
	// todo: maybe refactor these three into one function
	if in.Spec.Terminals != 0 {
		result.Spec.Template.Spec.Containers[0].Env = append(result.Spec.Template.Spec.Containers[0].Env, v1.EnvVar{
			Name:  "TERMINALS",
			Value: strconv.Itoa(int(in.Spec.Terminals)),
		})
	}
	if in.Spec.LoadWorkers != 0 {
		result.Spec.Template.Spec.Containers[0].Env = append(result.Spec.Template.Spec.Containers[0].Env, v1.EnvVar{
			Name:  "LOADWORKERS",
			Value: strconv.Itoa(int(in.Spec.LoadWorkers)),
		})
	}
	if in.Spec.Warehouses != 0 {
		result.Spec.Template.Spec.Containers[0].Env = append(result.Spec.Template.Spec.Containers[0].Env, v1.EnvVar{
			Name:  "WAREHOUSES",
			Value: strconv.Itoa(int(in.Spec.Warehouses)),
		})
	}
	return &result, nil
}

// +kubebuilder:object:root=true

// TpccBenchmarkList contains a list of TpccBenchmark
type TpccBenchmarkList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []TpccBenchmark `json:"items"`
}

func init() {
	SchemeBuilder.Register(&TpccBenchmark{}, &TpccBenchmarkList{})
}
