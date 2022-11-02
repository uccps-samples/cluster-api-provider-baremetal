/*
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 *
 * Copyright 2019 Red Hat, Inc.
 *
 */
// Code generated by informer-gen. DO NOT EDIT.

package v1beta1

import (
	internalinterfaces "github.com/uccps-samples/machine-api-operator/pkg/generated/informers/externalversions/internalinterfaces"
)

// Interface provides access to all the informers in this group version.
type Interface interface {
	// Machines returns a MachineInformer.
	Machines() MachineInformer
	// MachineHealthChecks returns a MachineHealthCheckInformer.
	MachineHealthChecks() MachineHealthCheckInformer
	// MachineSets returns a MachineSetInformer.
	MachineSets() MachineSetInformer
}

type version struct {
	factory          internalinterfaces.SharedInformerFactory
	namespace        string
	tweakListOptions internalinterfaces.TweakListOptionsFunc
}

// New returns a new Interface.
func New(f internalinterfaces.SharedInformerFactory, namespace string, tweakListOptions internalinterfaces.TweakListOptionsFunc) Interface {
	return &version{factory: f, namespace: namespace, tweakListOptions: tweakListOptions}
}

// Machines returns a MachineInformer.
func (v *version) Machines() MachineInformer {
	return &machineInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// MachineHealthChecks returns a MachineHealthCheckInformer.
func (v *version) MachineHealthChecks() MachineHealthCheckInformer {
	return &machineHealthCheckInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// MachineSets returns a MachineSetInformer.
func (v *version) MachineSets() MachineSetInformer {
	return &machineSetInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}