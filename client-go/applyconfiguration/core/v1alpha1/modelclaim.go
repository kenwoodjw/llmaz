/*
Copyright 2025 The InftyAI Team.

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

// Code generated by applyconfiguration-gen. DO NOT EDIT.

package v1alpha1

import (
	corev1alpha1 "github.com/inftyai/llmaz/api/core/v1alpha1"
)

// ModelClaimApplyConfiguration represents a declarative configuration of the ModelClaim type for use
// with apply.
type ModelClaimApplyConfiguration struct {
	ModelName        *corev1alpha1.ModelName   `json:"modelName,omitempty"`
	InferenceFlavors []corev1alpha1.FlavorName `json:"inferenceFlavors,omitempty"`
}

// ModelClaimApplyConfiguration constructs a declarative configuration of the ModelClaim type for use with
// apply.
func ModelClaim() *ModelClaimApplyConfiguration {
	return &ModelClaimApplyConfiguration{}
}

// WithModelName sets the ModelName field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the ModelName field is set to the value of the last call.
func (b *ModelClaimApplyConfiguration) WithModelName(value corev1alpha1.ModelName) *ModelClaimApplyConfiguration {
	b.ModelName = &value
	return b
}

// WithInferenceFlavors adds the given value to the InferenceFlavors field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, values provided by each call will be appended to the InferenceFlavors field.
func (b *ModelClaimApplyConfiguration) WithInferenceFlavors(values ...corev1alpha1.FlavorName) *ModelClaimApplyConfiguration {
	for i := range values {
		b.InferenceFlavors = append(b.InferenceFlavors, values[i])
	}
	return b
}
