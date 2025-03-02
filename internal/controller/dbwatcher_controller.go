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

package controller

import (
	"context"
	"encoding/json"
	"fmt"

	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/types"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/log"
	"sigs.k8s.io/controller-runtime/pkg/reconcile"

	dbv1 "github.com/shasan101/db-backup-gen/api/v1"
	batchv1 "k8s.io/api/batch/v1"
)

// DbWatcherReconciler reconciles a DbWatcher object
type DbWatcherReconciler struct {
	client.Client
	Scheme *runtime.Scheme
}

// +kubebuilder:rbac:groups=db.shasan.com,resources=dbwatchers,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=db.shasan.com,resources=dbwatchers/status,verbs=get;update;patch
// +kubebuilder:rbac:groups=db.shasan.com,resources=dbwatchers/finalizers,verbs=update

// +kubebuilder:rbac:groups=batch,resources=jobs,verbs=create;get;list;watch;update;patch;delete



// Reconcile is part of the main kubernetes reconciliation loop which aims to
// move the current state of the cluster closer to the desired state.
// TODO(user): Modify the Reconcile function to compare the state specified by
// the DbWatcher object against the actual cluster state, and then
// perform operations to make the cluster state reflect the state specified by
// the user.
//
// For more details, check Reconcile and its Result here:
// - https://pkg.go.dev/sigs.k8s.io/controller-runtime@v0.20.0/pkg/reconcile
func (r *DbWatcherReconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
	logger := log.FromContext(ctx)
	logger.Info("Reconcer is Ready!!!")
	var watcherObject dbv1.DbWatcher
	err := r.Get(ctx, req.NamespacedName, &watcherObject)
	if err != nil {
		logger.Info("error occurred: " + err.Error())
	}
	marshalledObj, _ := json.Marshal(watcherObject)
	logger.Info("deployed resource: " + string(marshalledObj))
	r.startBackupJob(ctx, &watcherObject)

	// TODO(user): your logic here

	return ctrl.Result{}, nil
}

func (r *DbWatcherReconciler) startBackupJob(ctx context.Context, jobInput *dbv1.DbWatcher) {
	logger := log.FromContext(ctx)
	logger.Info("Creating the backup job!!!")

	jobEnvVars := []corev1.EnvVar{
		{Name: "BACKUP_SOURCE_NAME", Value: jobInput.Spec.DatabaseName},
		{Name: "BACKUP_DB_NAME", Value: jobInput.Spec.DatabaseName},
		{Name: "BACKUP_USERNAME", Value: jobInput.Spec.Username},
		{Name: "BACKUP_PASSWORD", Value: jobInput.Spec.Password},
		{Name: "BACKUP_HOST", Value: jobInput.Spec.Host},
		{Name: "BACKUP_PORT", Value: jobInput.Spec.Port},
		{Name: "BACKUP_DEST_TYPE", Value: jobInput.Spec.BackupDestType},
		{Name: "BACKUP_DEST_NAME", Value: jobInput.Spec.BackupDestName},
		{Name: "BACKUP_DEST_PATH", Value: jobInput.Spec.BackupDestPath},
	}

	jobToCreate := batchv1.Job{
		ObjectMeta: metav1.ObjectMeta{
			Name: fmt.Sprintf("backup-operator-%s", jobInput.Name),
			Namespace: "db-backup-gen-system",
		},
		Spec: batchv1.JobSpec{
			Template: corev1.PodTemplateSpec{
				Spec: corev1.PodSpec{
					Containers: []corev1.Container{
						{
							Name:    "backup",
							Image:   "db-backup:latest",
							// Command: []string{""},
							Env:     jobEnvVars,
							ImagePullPolicy: corev1.PullNever,
						},
					},
					RestartPolicy: corev1.RestartPolicyNever,
				},
			},
		},
	}

	err := r.Client.Create(ctx, &jobToCreate)
	if err != nil {
		logger.Error(err, "failed to create job")
	}

	logger.Info(fmt.Sprintf("created job successfully: %v", jobInput.Name))
}

func (r *DbWatcherReconciler) HandlePodEvents(pod client.Object) []reconcile.Request {
	if pod.GetNamespace() != "default" {
		return []reconcile.Request{}
	}

	namespacedName := types.NamespacedName{
		Namespace: pod.GetNamespace(),
		Name:      pod.GetName(),
	}

	var podObject corev1.Pod
	err := r.Get(context.Background(), namespacedName, &podObject)
	if err != nil {
		return []reconcile.Request{}
	}

	podObject.SetAnnotations(map[string]string{
		"exampleAnnotation": "annotate.from.controller",
	})

	if err := r.Update(context.TODO(), &podObject); err != nil {
		log.Log.V(1).Info("error trying to annotate pod: " + err.Error())
	}
	return []reconcile.Request{}

}

// SetupWithManager sets up the controller with the Manager.
func (r *DbWatcherReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&dbv1.DbWatcher{}).
		Named("dbwatcher").
		Complete(r)
}
