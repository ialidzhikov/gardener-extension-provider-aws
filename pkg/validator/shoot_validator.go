// Copyright (c) 2019 SAP SE or an SAP affiliate company. All rights reserved. This file is licensed under the Apache Software License, v. 2 except as noted otherwise in the LICENSE file
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package validator

import (
	"context"
	"errors"
	"fmt"
	"reflect"

	apiaws "github.com/gardener/gardener-extension-provider-aws/pkg/apis/aws"
	awsvalidation "github.com/gardener/gardener-extension-provider-aws/pkg/apis/aws/validation"
	"github.com/gardener/gardener-extension-provider-aws/pkg/aws"

	extensionscontroller "github.com/gardener/gardener/extensions/pkg/controller"
	"github.com/gardener/gardener/pkg/apis/core"
	gardencorev1beta1 "github.com/gardener/gardener/pkg/apis/core/v1beta1"
	kutil "github.com/gardener/gardener/pkg/utils/kubernetes"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/util/validation/field"
)

func (v *Shoot) validateShoot(ctx context.Context, shoot *core.Shoot) error {
	// Network validation
	networkPath := field.NewPath("spec", "networking")

	if errList := awsvalidation.ValidateNetworking(shoot.Spec.Networking, networkPath); len(errList) != 0 {
		return errList.ToAggregate()
	}

	// Provider validation
	fldPath := field.NewPath("spec", "provider")

	// InfrastructureConfig
	infraConfigFldPath := fldPath.Child("infrastructureConfig")

	if shoot.Spec.Provider.InfrastructureConfig == nil {
		return field.Required(infraConfigFldPath, "InfrastructureConfig must be set for AWS shoots")
	}

	infraConfig, err := decodeInfrastructureConfig(v.decoder, shoot.Spec.Provider.InfrastructureConfig, infraConfigFldPath)
	if err != nil {
		return err
	}

	if errList := awsvalidation.ValidateInfrastructureConfig(infraConfig, shoot.Spec.Networking.Nodes, shoot.Spec.Networking.Pods, shoot.Spec.Networking.Services); len(errList) != 0 {
		return errList.ToAggregate()
	}

	// ControlPlaneConfig
	if shoot.Spec.Provider.ControlPlaneConfig != nil {
		if _, err := decodeControlPlaneConfig(v.decoder, shoot.Spec.Provider.ControlPlaneConfig, fldPath.Child("controlPlaneConfig")); err != nil {
			return err
		}
	}

	// WorkerConfig
	fldPath = fldPath.Child("workers")
	for i, worker := range shoot.Spec.Provider.Workers {
		if worker.ProviderConfig != nil {
			workerConfig, err := decodeWorkerConfig(v.decoder, worker.ProviderConfig, fldPath.Index(i).Child("providerConfig"))
			if err != nil {
				return err
			}

			var volumeType *string
			if worker.Volume != nil {
				volumeType = worker.Volume.Type
			}

			if errList := awsvalidation.ValidateWorkerConfig(workerConfig, volumeType); len(errList) != 0 {
				return errList.ToAggregate()
			}
		}
	}

	// Shoot workers
	if errList := awsvalidation.ValidateWorkers(shoot.Spec.Provider.Workers, infraConfig.Networks.Zones, fldPath); len(errList) != 0 {
		return errList.ToAggregate()
	}

	return nil
}

func (v *Shoot) validateShootUpdate(ctx context.Context, oldShoot, shoot *core.Shoot) error {
	var (
		fldPath            = field.NewPath("spec", "provider")
		infraConfigFldPath = fldPath.Child("infrastructureConfig")
	)

	// InfrastructureConfig update
	if shoot.Spec.Provider.InfrastructureConfig == nil {
		return field.Required(fldPath.Child("infrastructureConfig"), "InfrastructureConfig must be set for AWS shoots")
	}

	infraConfig, err := decodeInfrastructureConfig(v.decoder, shoot.Spec.Provider.InfrastructureConfig, infraConfigFldPath)
	if err != nil {
		return err
	}

	if oldShoot.Spec.Provider.InfrastructureConfig == nil {
		return field.InternalError(infraConfigFldPath, errors.New("InfrastructureConfig is not available on old shoot"))
	}

	oldInfraConfig, err := decodeInfrastructureConfig(v.decoder, oldShoot.Spec.Provider.InfrastructureConfig, infraConfigFldPath)
	if err != nil {
		return err
	}

	if !reflect.DeepEqual(oldInfraConfig, infraConfig) {
		if errList := awsvalidation.ValidateInfrastructureConfigUpdate(oldInfraConfig, infraConfig); len(errList) != 0 {
			return errList.ToAggregate()
		}
	}

	// Only validate against cloud profile when zones are updated.
	// This ensures that already running shoots won't break after a zone is removed from a cloud profile.
	if len(infraConfig.Networks.Zones) > len(oldInfraConfig.Networks.Zones) {
		if err := v.validateAgainstCloudProfile(ctx, shoot, infraConfig, infraConfigFldPath); err != nil {
			return err
		}
	}

	if errList := awsvalidation.ValidateWorkersUpdate(oldShoot.Spec.Provider.Workers, shoot.Spec.Provider.Workers, fldPath.Child("workers")); len(errList) != 0 {
		return errList.ToAggregate()
	}

	return v.validateShoot(ctx, shoot)
}

func (v *Shoot) validateShootCreation(ctx context.Context, shoot *core.Shoot) error {
	fldPath := field.NewPath("spec", "provider")
	infraConfig, err := decodeInfrastructureConfig(v.decoder, shoot.Spec.Provider.InfrastructureConfig, fldPath.Child("infrastructureConfig"))
	if err != nil {
		return err
	}

	if err := v.validateAgainstCloudProfile(ctx, shoot, infraConfig, fldPath.Child("infrastructureConfig")); err != nil {
		return err
	}

	if err := v.validateCloudProviderCredentials(ctx, shoot); err != nil {
		return err
	}

	return v.validateShoot(ctx, shoot)
}

func (v *Shoot) validateAgainstCloudProfile(ctx context.Context, shoot *core.Shoot, infraConfig *apiaws.InfrastructureConfig, fldPath *field.Path) error {
	cloudProfile := &gardencorev1beta1.CloudProfile{}
	if err := v.client.Get(ctx, kutil.Key(shoot.Spec.CloudProfileName), cloudProfile); err != nil {
		return err
	}

	if errList := awsvalidation.ValidateInfrastructureConfigAgainstCloudProfile(infraConfig, shoot, cloudProfile, fldPath); len(errList) != 0 {
		return errList.ToAggregate()
	}

	return nil
}

func (v *Shoot) validateCloudProviderCredentials(ctx context.Context, shoot *core.Shoot) error {
	secretBinding := &gardencorev1beta1.SecretBinding{}
	if err := v.client.Get(ctx, kutil.Key(shoot.Namespace, shoot.Spec.SecretBindingName), secretBinding); err != nil {
		return err
	}

	secret, err := extensionscontroller.GetSecretByReference(ctx, v.client, &secretBinding.SecretRef)
	if err != nil {
		return err
	}

	if err := validateCloudProviderSecret(secret); err != nil {
		return err
	}

	return nil
}

func validateCloudProviderSecret(secret *corev1.Secret) error {
	_, ok := secret.Data[aws.AccessKeyID]
	if !ok {
		return fmt.Errorf("missing %q field in secret", aws.AccessKeyID)
	}

	_, ok = secret.Data[aws.SecretAccessKey]
	if !ok {
		return fmt.Errorf("missing %q field in secret", aws.SecretAccessKey)
	}

	return nil
}
