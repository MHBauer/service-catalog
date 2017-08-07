/*
Copyright 2017 The Kubernetes Authors.

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

// This file was automatically generated by informer-gen

package internalversion

import (
	internalinterfaces "github.com/kubernetes-incubator/service-catalog/pkg/client/informers_generated/internalversion/internalinterfaces"
)

// Interface provides access to all the informers in this group version.
type Interface interface {
	// Bindings returns a BindingInformer.
	Bindings() BindingInformer
	// Brokers returns a BrokerInformer.
	Brokers() BrokerInformer
	// Instances returns a InstanceInformer.
	Instances() InstanceInformer
	// Plans returns a PlanInformer.
	Plans() PlanInformer
	// ServiceClasses returns a ServiceClassInformer.
	ServiceClasses() ServiceClassInformer
}

type version struct {
	internalinterfaces.SharedInformerFactory
}

// New returns a new Interface.
func New(f internalinterfaces.SharedInformerFactory) Interface {
	return &version{f}
}

// Bindings returns a BindingInformer.
func (v *version) Bindings() BindingInformer {
	return &bindingInformer{factory: v.SharedInformerFactory}
}

// Brokers returns a BrokerInformer.
func (v *version) Brokers() BrokerInformer {
	return &brokerInformer{factory: v.SharedInformerFactory}
}

// Instances returns a InstanceInformer.
func (v *version) Instances() InstanceInformer {
	return &instanceInformer{factory: v.SharedInformerFactory}
}

// Plans returns a PlanInformer.
func (v *version) Plans() PlanInformer {
	return &planInformer{factory: v.SharedInformerFactory}
}

// ServiceClasses returns a ServiceClassInformer.
func (v *version) ServiceClasses() ServiceClassInformer {
	return &serviceClassInformer{factory: v.SharedInformerFactory}
}
