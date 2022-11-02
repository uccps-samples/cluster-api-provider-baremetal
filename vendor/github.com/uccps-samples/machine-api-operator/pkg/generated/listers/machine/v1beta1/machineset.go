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
// Code generated by lister-gen. DO NOT EDIT.

package v1beta1

import (
	v1beta1 "github.com/uccps-samples/machine-api-operator/pkg/apis/machine/v1beta1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// MachineSetLister helps list MachineSets.
// All objects returned here must be treated as read-only.
type MachineSetLister interface {
	// List lists all MachineSets in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1beta1.MachineSet, err error)
	// MachineSets returns an object that can list and get MachineSets.
	MachineSets(namespace string) MachineSetNamespaceLister
	MachineSetListerExpansion
}

// machineSetLister implements the MachineSetLister interface.
type machineSetLister struct {
	indexer cache.Indexer
}

// NewMachineSetLister returns a new MachineSetLister.
func NewMachineSetLister(indexer cache.Indexer) MachineSetLister {
	return &machineSetLister{indexer: indexer}
}

// List lists all MachineSets in the indexer.
func (s *machineSetLister) List(selector labels.Selector) (ret []*v1beta1.MachineSet, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1beta1.MachineSet))
	})
	return ret, err
}

// MachineSets returns an object that can list and get MachineSets.
func (s *machineSetLister) MachineSets(namespace string) MachineSetNamespaceLister {
	return machineSetNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// MachineSetNamespaceLister helps list and get MachineSets.
// All objects returned here must be treated as read-only.
type MachineSetNamespaceLister interface {
	// List lists all MachineSets in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1beta1.MachineSet, err error)
	// Get retrieves the MachineSet from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1beta1.MachineSet, error)
	MachineSetNamespaceListerExpansion
}

// machineSetNamespaceLister implements the MachineSetNamespaceLister
// interface.
type machineSetNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all MachineSets in the indexer for a given namespace.
func (s machineSetNamespaceLister) List(selector labels.Selector) (ret []*v1beta1.MachineSet, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1beta1.MachineSet))
	})
	return ret, err
}

// Get retrieves the MachineSet from the indexer for a given namespace and name.
func (s machineSetNamespaceLister) Get(name string) (*v1beta1.MachineSet, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1beta1.Resource("machineset"), name)
	}
	return obj.(*v1beta1.MachineSet), nil
}
