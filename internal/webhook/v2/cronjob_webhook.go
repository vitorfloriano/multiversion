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

package v2

import (
	"context"
	"fmt"

	"k8s.io/apimachinery/pkg/runtime"
	ctrl "sigs.k8s.io/controller-runtime"
	logf "sigs.k8s.io/controller-runtime/pkg/log"
	"sigs.k8s.io/controller-runtime/pkg/webhook"
	"sigs.k8s.io/controller-runtime/pkg/webhook/admission"

	batchv2 "github.com/vitorfloriano/testproject/api/v2"
)

// nolint:unused
// log is for logging in this package.
var cronjoblog = logf.Log.WithName("cronjob-resource")

// SetupCronJobWebhookWithManager registers the webhook for CronJob in the manager.
func SetupCronJobWebhookWithManager(mgr ctrl.Manager) error {
	return ctrl.NewWebhookManagedBy(mgr).For(&batchv2.CronJob{}).
		WithValidator(&CronJobCustomValidator{}).
		WithDefaulter(&CronJobCustomDefaulter{}).
		Complete()
}

// TODO(user): EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!

// +kubebuilder:webhook:path=/mutate-batch-tutorial-kubebuilder-io-v2-cronjob,mutating=true,failurePolicy=fail,sideEffects=None,groups=batch.tutorial.kubebuilder.io,resources=cronjobs,verbs=create;update,versions=v2,name=mcronjob-v2.kb.io,admissionReviewVersions=v1

// CronJobCustomDefaulter struct is responsible for setting default values on the custom resource of the
// Kind CronJob when those are created or updated.
//
// NOTE: The +kubebuilder:object:generate=false marker prevents controller-gen from generating DeepCopy methods,
// as it is used only for temporary operations and does not need to be deeply copied.
type CronJobCustomDefaulter struct {
	// TODO(user): Add more fields as needed for defaulting
}

var _ webhook.CustomDefaulter = &CronJobCustomDefaulter{}

// Default implements webhook.CustomDefaulter so a webhook will be registered for the Kind CronJob.
func (d *CronJobCustomDefaulter) Default(ctx context.Context, obj runtime.Object) error {
	cronjob, ok := obj.(*batchv2.CronJob)

	if !ok {
		return fmt.Errorf("expected an CronJob object but got %T", obj)
	}
	cronjoblog.Info("Defaulting for CronJob", "name", cronjob.GetName())

	// TODO(user): fill in your defaulting logic.

	return nil
}

// TODO(user): change verbs to "verbs=create;update;delete" if you want to enable deletion validation.
// NOTE: The 'path' attribute must follow a specific pattern and should not be modified directly here.
// Modifying the path for an invalid path can cause API server errors; failing to locate the webhook.
// +kubebuilder:webhook:path=/validate-batch-tutorial-kubebuilder-io-v2-cronjob,mutating=false,failurePolicy=fail,sideEffects=None,groups=batch.tutorial.kubebuilder.io,resources=cronjobs,verbs=create;update,versions=v2,name=vcronjob-v2.kb.io,admissionReviewVersions=v1

// CronJobCustomValidator struct is responsible for validating the CronJob resource
// when it is created, updated, or deleted.
//
// NOTE: The +kubebuilder:object:generate=false marker prevents controller-gen from generating DeepCopy methods,
// as this struct is used only for temporary operations and does not need to be deeply copied.
type CronJobCustomValidator struct {
	// TODO(user): Add more fields as needed for validation
}

var _ webhook.CustomValidator = &CronJobCustomValidator{}

// ValidateCreate implements webhook.CustomValidator so a webhook will be registered for the type CronJob.
func (v *CronJobCustomValidator) ValidateCreate(ctx context.Context, obj runtime.Object) (admission.Warnings, error) {
	cronjob, ok := obj.(*batchv2.CronJob)
	if !ok {
		return nil, fmt.Errorf("expected a CronJob object but got %T", obj)
	}
	cronjoblog.Info("Validation for CronJob upon creation", "name", cronjob.GetName())

	// TODO(user): fill in your validation logic upon object creation.

	return nil, nil
}

// ValidateUpdate implements webhook.CustomValidator so a webhook will be registered for the type CronJob.
func (v *CronJobCustomValidator) ValidateUpdate(ctx context.Context, oldObj, newObj runtime.Object) (admission.Warnings, error) {
	cronjob, ok := newObj.(*batchv2.CronJob)
	if !ok {
		return nil, fmt.Errorf("expected a CronJob object for the newObj but got %T", newObj)
	}
	cronjoblog.Info("Validation for CronJob upon update", "name", cronjob.GetName())

	// TODO(user): fill in your validation logic upon object update.

	return nil, nil
}

// ValidateDelete implements webhook.CustomValidator so a webhook will be registered for the type CronJob.
func (v *CronJobCustomValidator) ValidateDelete(ctx context.Context, obj runtime.Object) (admission.Warnings, error) {
	cronjob, ok := obj.(*batchv2.CronJob)
	if !ok {
		return nil, fmt.Errorf("expected a CronJob object but got %T", obj)
	}
	cronjoblog.Info("Validation for CronJob upon deletion", "name", cronjob.GetName())

	// TODO(user): fill in your validation logic upon object deletion.

	return nil, nil
}
