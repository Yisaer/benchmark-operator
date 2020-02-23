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

package controllers

import (
	"context"
	"github.com/go-logr/logr"
	"github.com/yisaer/benchmark-operator/api/v1alpha1"
	batchv1 "k8s.io/api/batch/v1"
	corev1 "k8s.io/api/core/v1"
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"strconv"
)

const (
	// todo: make this configable in user cr's yaml
	jobImage = "longfangsong/benchmarksql:v0"
	protocol = "jdbc:mysql://"
)

func createJob(in *v1alpha1.TpccBenchmark) (*batchv1.Job, error) {
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
						Image: jobImage,
						// todo: check whether empty values are handled here
						Env: []v1.EnvVar{
							{Name: "CONN", Value: *in.Spec.Conn},
							{Name: "USER", Value: in.Spec.Username},
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

// TpccBenchmarkReconciler reconciles a TpccBenchmark object
type TpccBenchmarkReconciler struct {
	client.Client
	Log    logr.Logger
	Scheme *runtime.Scheme
}

// +kubebuilder:rbac:groups=benchmark.tidb.pingcap.com.benchmark.pingcap.com,resources=tpccbenchmarks,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=benchmark.tidb.pingcap.com.benchmark.pingcap.com,resources=tpccbenchmarks/status,verbs=get;update;patch

func (r *TpccBenchmarkReconciler) Reconcile(req ctrl.Request) (ctrl.Result, error) {
	ctx := context.Background()
	log := r.Log.WithValues("benchmarksql", req.NamespacedName)
	var benchmark v1alpha1.TpccBenchmark
	if err := r.Get(ctx, req.NamespacedName, &benchmark); err != nil {
		log.Error(err, "unable to fetch benchmark")
		return ctrl.Result{}, client.IgnoreNotFound(err)
	}
	constructJob := func(request *v1alpha1.TpccBenchmark) (*batchv1.Job, error) {
		job, err := createJob(request)
		if err := ctrl.SetControllerReference(request, job, r.Scheme); err != nil {
			return nil, err
		}
		return job, err
	}
	job, err := constructJob(&benchmark)
	if err != nil {
		log.Error(err, "unable to create job")
		return ctrl.Result{}, client.IgnoreNotFound(err)
	}
	if err := r.Create(ctx, job); err != nil {
		log.Error(err, "unable to create Job for CronJob", "job", job)
		return ctrl.Result{}, err
	}
	return ctrl.Result{}, err
}

func (r *TpccBenchmarkReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&v1alpha1.TpccBenchmark{}).
		Complete(r)
}
