/*
Copyright 2019 LitmusChaos Authors

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

package service

import (
	"errors"
	"fmt"

	corev1 "k8s.io/api/core/v1"
)

// Builder is the builder object for Service
type Builder struct {
	service *Service
	errs    []error
}

// NewBuilder returns new instance of Builder
func NewBuilder() *Builder {
	return &Builder{service: &Service{object: &corev1.Service{}}}
}

// WithName sets the Name field of Service with provided value.
func (b *Builder) WithName(name string) *Builder {
	if len(name) == 0 {
		b.errs = append(
			b.errs,
			errors.
				New("failed to build service object: missing name"),
		)
		return b
	}
	b.service.object.Name = name
	return b
}

// WithNamespace sets the Namespace field of Service provided arguments
func (b *Builder) WithNamespace(namespace string) *Builder {
	if len(namespace) == 0 {
		b.errs = append(
			b.errs,
			errors.
				New("failed to build service object: missing namespace"),
		)
		return b
	}
	b.service.object.Namespace = namespace
	return b
}

// WithLabels sets the labels field of Service provided arguments
func (b *Builder) WithLabels(labels map[string]string) *Builder {
	if len(labels) == 0 {
		b.errs = append(
			b.errs,
			errors.New("failed to build service object: missing labels"),
		)
		return b
	}

	if b.service.object.Labels == nil {
		b.service.object.Labels = map[string]string{}
	}

	for key, value := range labels {
		b.service.object.Labels[key] = value
	}
	return b
}

// WithPorts sets the ports field of Service provided arguments
func (b *Builder) WithPorts(ports []corev1.ServicePort) *Builder {
	if len(ports) == 0 {
		b.errs = append(
			b.errs,
			errors.New("failed to build service object: missing ports"),
		)
		return b
	}

	// copy of original slice
	newports := []corev1.ServicePort{}
	newports = append(newports, ports...)

	// override
	b.service.object.Spec.Ports = newports
	return b
}

// WithSelectorsNew resets existing selectors if any with
// ones that are provided here
func (b *Builder) WithSelectorsNew(selectors map[string]string) *Builder {
	if len(selectors) == 0 {
		b.errs = append(
			b.errs,
			errors.New("failed to build service object: no new selectors"),
		)
		return b
	}

	// copy of original map
	newslctrs := map[string]string{}
	for key, value := range selectors {
		newslctrs[key] = value
	}

	// override
	b.service.object.Spec.Selector = newslctrs
	return b
}

// Build returns the Service API instance
func (b *Builder) Build() (*corev1.Service, error) {
	if len(b.errs) > 0 {
		return nil, fmt.Errorf("%+v", b.errs)
	}
	return b.service.object, nil
}
