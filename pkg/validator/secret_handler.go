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
	"net/http"

	"github.com/gardener/gardener-extension-provider-aws/pkg/aws"
	gardencorev1beta1 "github.com/gardener/gardener/pkg/apis/core/v1beta1"
	kutil "github.com/gardener/gardener/pkg/utils/kubernetes"
	"github.com/go-logr/logr"
	admissionv1beta1 "k8s.io/api/admission/v1beta1"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/serializer"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/webhook/admission"
)

// Shoot validates shoots
type Secret struct {
	client  client.Client
	decoder runtime.Decoder
	Logger  logr.Logger
}

// Handle implements Handler.Handle
func (v *Secret) Handle(ctx context.Context, req admission.Request) admission.Response {
	// continue from des. check controller-runtime
	secret := &corev1.Secret{}

	if err := runtime.DecodeInto(v.decoder, req.Object.Raw, secret); err != nil {
		v.Logger.Error(err, "failed to decode secret")
		return admission.Errored(http.StatusBadRequest, err)
	}

	v.Logger.Info("Inside2", "secret", secret.Data)

	switch req.Operation {
	case admissionv1beta1.Update:
		// TODO: use lister
		// Check if secret used by Shoot of the respective cloud provider

		shoots := &gardencorev1beta1.ShootList{}
		if err := v.client.List(ctx, shoots); err != nil {
			return admission.Errored(http.StatusBadRequest, err)
		}

		// TODO: think about optimization
		for _, shoot := range shoots.Items {
			if shoot.Spec.Provider.Type != aws.Type {
				continue
			}

			secretBinding := &gardencorev1beta1.SecretBinding{}
			if err := v.client.Get(ctx, kutil.Key(shoot.Namespace, shoot.Spec.SecretBindingName), secretBinding); err != nil {
				continue
			}

			if secretBinding.SecretRef.Name == secret.Name && secretBinding.SecretRef.Namespace == secret.Namespace {
				// there is a Shoot from this provider type using the secret
				if err := validateCloudProviderSecret(secret); err != nil {
					return admission.Errored(http.StatusBadRequest, err)
				}
			}
		}

		/*if err := v.validateShootUpdate(ctx, oldShoot, shoot); err != nil {
			v.Logger.Error(err, "denied request")
			return admission.Errored(http.StatusBadRequest, err)
		}*/
	default:
		v.Logger.Info("Webhook not responsible", "operation", req.Operation)
	}

	return admission.Allowed("validations succeeded")
}

// InjectClient injects the client.
func (v *Secret) InjectClient(c client.Client) error {
	v.client = c
	return nil
}

// InjectScheme injects the scheme.
func (v *Secret) InjectScheme(s *runtime.Scheme) error {
	v.decoder = serializer.NewCodecFactory(s).UniversalDecoder()
	return nil
}
