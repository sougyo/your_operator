/*
Copyright 2023.

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

	"k8s.io/apimachinery/pkg/runtime"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/log"

	appv1alpha1 "github.com/sougyo/your-operator/api/v1alpha1"
)

// YourCustomResourceReconciler reconciles a YourCustomResource object
type YourCustomResourceReconciler struct {
	client.Client
	Scheme *runtime.Scheme
}

//+kubebuilder:rbac:groups=app.example.com,resources=yourcustomresources,verbs=get;list;watch;create;update;patch;delete
//+kubebuilder:rbac:groups=app.example.com,resources=yourcustomresources/status,verbs=get;update;patch
//+kubebuilder:rbac:groups=app.example.com,resources=yourcustomresources/finalizers,verbs=update

// Reconcile is part of the main kubernetes reconciliation loop which aims to
// move the current state of the cluster closer to the desired state.
// TODO(user): Modify the Reconcile function to compare the state specified by
// the YourCustomResource object against the actual cluster state, and then
// perform operations to make the cluster state reflect the state specified by
// the user.
//
// For more details, check Reconcile and its Result here:
// - https://pkg.go.dev/sigs.k8s.io/controller-runtime@v0.14.1/pkg/reconcile
func (r *YourCustomResourceReconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
	logger := log.FromContext(ctx)

	logger.Info("Reconciling YourCustomResourceXX")

	yourCustomResource := appv1alpha1.YourCustomResource{}
	if err := r.Get(ctx, req.NamespacedName, &yourCustomResource); err != nil {
		logger.Error(err, "unable to fetch YourCustomResource")
		return ctrl.Result{}, client.IgnoreNotFound(err)
	}
	//log Foo spec
	logger.Info("Foo: " + yourCustomResource.Spec.Foo)

	//set Foo status to Running
	//yourCustomResource.Status.State = "Running"

	return ctrl.Result{}, nil
}

// SetupWithManager sets up the controller with the Manager.
func (r *YourCustomResourceReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&appv1alpha1.YourCustomResource{}).
		Complete(r)
}
