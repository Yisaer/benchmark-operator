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
	"github.com/go-logr/logr"
	batchv1 "k8s.io/api/batch/v1"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"

	"github.com/yisaer/benchmark-operator/api/v1alpha1"
)

// SysBenchmarkReconciler reconciles a SysBenchmark object
type SysBenchmarkReconciler struct {
	client.Client
	Log    logr.Logger
	Scheme *runtime.Scheme
}

const (
	jobOwnerLabel = ".metadata.controller"
)
const (
	PrepareStage = 0
	RunStage     = 1
	CleanStage   = 2 // not used now
)

func commandsForJob(sysBenchmark *v1alpha1.SysBenchmark, stage uint32) []string {
	commands := []string{
		"sysbench",
		"--db-driver=mysql",
		fmt.Sprint("--oltp-table-size=", sysBenchmark.Spec.TableSize),
		fmt.Sprint("--oltp-tables-count=", sysBenchmark.Spec.TablesCount),
		fmt.Sprint("--threads=", sysBenchmark.Spec.Threads),
		fmt.Sprint("--mysql-host=", sysBenchmark.Spec.Host),
		fmt.Sprint("--mysql-port=", sysBenchmark.Spec.Port),
		fmt.Sprint("--mysql-user=", sysBenchmark.Spec.Username),
		fmt.Sprint("--mysql-password=", sysBenchmark.Spec.Password),
	}
	switch stage {
	case PrepareStage:
		commands = append(commands, "/usr/share/sysbench/tests/include/oltp_legacy/parallel_prepare.lua")
	case RunStage:
		commands = append(commands,
			fmt.Sprint("--time=", sysBenchmark.Spec.Time),
			fmt.Sprint("--report-interval=", sysBenchmark.Spec.ReportInterval),
			"/usr/share/sysbench/tests/include/oltp_legacy/oltp.lua")
	}
	commands = append(commands, "run")
	return commands
}

func nameForJob(sysBenchmark *v1alpha1.SysBenchmark, stage uint32) string {
	switch stage {
	case PrepareStage:
		return sysBenchmark.Name + "-prepare"
	case RunStage:
		return sysBenchmark.Name
	}
	panic("invalid argument!")
}

func createSysBenchmarkJob(in *v1alpha1.SysBenchmark, stage uint32) (*batchv1.Job, error) {
	result := batchv1.Job{
		ObjectMeta: metav1.ObjectMeta{
			Name:      nameForJob(in, stage),
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
					Containers: []corev1.Container{{
						Name:    "sysbench",
						Image:   "severalnines/sysbench",
						Command: commandsForJob(in, stage),
					}},
					RestartPolicy: "Never",
				},
			},
		},
	}
	return &result, nil
}

// +kubebuilder:rbac:groups=benchmark.cloud.shuosc.org,resources=sysbenchmarks,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=benchmark.cloud.shuosc.org,resources=sysbenchmarks/status,verbs=get;update;patch
func (r *SysBenchmarkReconciler) Reconcile(req ctrl.Request) (ctrl.Result, error) {
	ctx := context.Background()
	log := r.Log.WithValues("sysbenchmark", req.NamespacedName)
	constructJob := func(request *v1alpha1.SysBenchmark, stage uint32) (*batchv1.Job, error) {
		job, err := createSysBenchmarkJob(request, stage)
		if err := ctrl.SetControllerReference(request, job, r.Scheme); err != nil {
			return nil, err
		}
		return job, err
	}
	var benchmark v1alpha1.SysBenchmark
	if err := r.Get(ctx, req.NamespacedName, &benchmark); err != nil {
		log.Error(err, "unable to fetch benchmark")
		return ctrl.Result{}, client.IgnoreNotFound(err)
	}
	var childJobs batchv1.JobList
	if err := r.List(ctx, &childJobs, client.InNamespace(req.Namespace), client.MatchingLabels{jobOwnerLabel: req.Name}); err != nil {
		log.Error(err, "unable to list child Jobs")
		return ctrl.Result{}, err
	}
	var job *batchv1.Job
	var err error
	for _, job := range childJobs.Items {
		log.Info("Child", "job", job.Name)
	}
	// no child Jobs, start prepare
	if len(childJobs.Items) == 0 {
		job, err = constructJob(&benchmark, PrepareStage)
		if err != nil {
			log.Error(err, "unable to construct prepare job")
			return ctrl.Result{}, err
		}
		log.Info("PrepareStage Job constructed")
	} else if len(childJobs.Items) == 1 && childJobs.Items[0].Status.Succeeded == 1 {
		job, err = constructJob(&benchmark, RunStage)
		if err != nil {
			log.Error(err, "unable to construct run job")
			return ctrl.Result{}, err
		}
		log.Info("RunStage Job constructed")
	} else {
		return ctrl.Result{}, nil
	}
	if err := r.Create(ctx, job); err != nil {
		log.Error(err, "unable to create Job", "job", job)
		return ctrl.Result{}, err
	}
	return ctrl.Result{}, nil
}

func (r *SysBenchmarkReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		Owns(&corev1.Pod{}).
		Owns(&batchv1.Job{}).
		For(&v1alpha1.SysBenchmark{}).
		Complete(r)
}
