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
	batchv1 "k8s.io/api/batch/v1"

	"github.com/go-logr/logr"
	"k8s.io/apimachinery/pkg/runtime"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"

	v1alpha1 "github.com/yisaer/benchmark-operator/api/v1alpha1"
)

// DataBaseBenchmarkPrepareReconciler reconciles a DataBaseBenchmarkPrepare object
type DataBaseBenchmarkPrepareReconciler struct {
	client.Client
	Log    logr.Logger
	Scheme *runtime.Scheme
}

// +kubebuilder:rbac:groups=benchmark.cloud.shuosc.org,resources=databasebenchmarkprepares,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=benchmark.cloud.shuosc.org,resources=databasebenchmarkprepares/status,verbs=get;update;patch

func (r *DataBaseBenchmarkPrepareReconciler) Reconcile(req ctrl.Request) (ctrl.Result, error) {
	ctx := context.Background()
	log := r.Log.WithValues("databasebenchmarkprepare", req.NamespacedName)
	constructJob := func(request *v1alpha1.DataBaseBenchmarkPrepare) (*batchv1.Job, error) {
		job, err := request.Spec.Prepares[0].CreateJob(request.Spec.Host, request.Spec.Host, request.Spec.User, request.Spec.Password)
		if err := ctrl.SetControllerReference(request, job, r.Scheme); err != nil {
			return nil, err
		}
		return job, err
	}
	var benchmark v1alpha1.DataBaseBenchmarkPrepare
	if err := r.Get(ctx, req.NamespacedName, &benchmark); err != nil {
		log.Error(err, "unable to fetch benchmark")
		return ctrl.Result{}, client.IgnoreNotFound(err)
	}
	job, err := constructJob(&benchmark)
	if err != nil {
		log.Error(err, "unable to create Job", "job", job)
		return ctrl.Result{}, err
	}
	if err := r.Create(ctx, job); err != nil {
		log.Error(err, "unable to create Job", "job", job)
		return ctrl.Result{}, err
	}
	return ctrl.Result{}, nil
}

func (r *DataBaseBenchmarkPrepareReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		Owns(&batchv1.Job{}).
		For(&v1alpha1.DataBaseBenchmarkPrepare{}).
		Complete(r)
}
