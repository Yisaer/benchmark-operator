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
	"fmt"
	"strconv"
	"time"

	"github.com/go-logr/logr"
	benchmarktidbpingcapcomv1alpha1 "github.com/yisaer/benchmark-operator/api/v1alpha1"
	batchv1 "k8s.io/api/batch/v1"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

const (
	jobImage = "longfangsong/benchmarksql:1582222183"
	protocol = "jdbc:mysql://"
)

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
	// your logic here
	//ctrl.CreateOrUpdate()
	var testRequest benchmarktidbpingcapcomv1alpha1.TpccBenchmark
	if err := r.Get(ctx, req.NamespacedName, &testRequest); err != nil {
		log.Error(err, "unable to fetch testRequest")
		return ctrl.Result{}, client.IgnoreNotFound(err)
	}

	var childJobs batchv1.JobList
	if err := r.List(ctx, &childJobs); err != nil {
		log.Error(err, "unable to list child Jobs")
		return ctrl.Result{}, err
	}

	conn := ""
	if testRequest.Spec.Conn != nil {
		conn = *testRequest.Spec.Conn
	} else {
		port := 4000
		c := testRequest.Spec.Cluster
		conn = fmt.Sprintf("%s%s-tidb.%s.svc:%d/%s", protocol, c.Name, c.Namespace, port, *testRequest.Spec.Database)
	}

	constructJob := func(request benchmarktidbpingcapcomv1alpha1.TpccBenchmark) (*batchv1.Job, error) {
		// We want job names for a given nominal start time to have a deterministic name to avoid the same job being created twice
		name := fmt.Sprintf("test")

		job := &batchv1.Job{
			ObjectMeta: metav1.ObjectMeta{
				Labels:      make(map[string]string),
				Annotations: make(map[string]string),
				Name:        name,
				Namespace:   "default",
			},
			Spec: batchv1.JobSpec{
				Parallelism:           nil,
				Completions:           nil,
				ActiveDeadlineSeconds: nil,
				BackoffLimit:          nil,
				Selector:              nil,
				ManualSelector:        nil,
				Template: corev1.PodTemplateSpec{
					ObjectMeta: metav1.ObjectMeta{},
					Spec: corev1.PodSpec{
						RestartPolicy: "OnFailure",
						Containers: []corev1.Container{{
							Name:  "test",
							Image: jobImage,
							Env: []corev1.EnvVar{
								{Name: "CONN", Value: conn},
								{Name: "WAREHOUSES", Value: strconv.Itoa(int(request.Spec.Warehouses))},
								{Name: "LOADWORKERS", Value: strconv.Itoa(int(request.Spec.LoadWorkers))},
								{Name: "TERMINALS", Value: strconv.Itoa(int(request.Spec.Terminals))},
							},
						}},
					},
				},
				TTLSecondsAfterFinished: nil,
			},
		}
		job.Annotations["createTime"] = time.Now().Format(time.RFC3339)
		return job, nil
	}
	job, err := constructJob(testRequest)
	if err != nil {
		log.Error(err, "unable to construct job from template")
		// don't bother requeuing until we get a change to the spec
		return ctrl.Result{}, nil
	}
	if err := r.Create(ctx, job); err != nil {
		log.Error(err, "unable to create Job for CronJob", "job", job)
		return ctrl.Result{}, err
	}
	return ctrl.Result{}, nil
}

func (r *TpccBenchmarkReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&benchmarktidbpingcapcomv1alpha1.TpccBenchmark{}).
		Complete(r)
}
