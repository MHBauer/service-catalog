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

package v1beta1

import (
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// The ReferencesExtension interface allows setting the References to
// ServiceClasses and ServicePlans.
type ServiceBindingExpansion interface {
	SpecialDelete(name string, options *v1.DeleteOptions) error
}

// Delete takes name of the serviceBinding and deletes it. Returns an error if
// one occurs.
func (c *serviceBindings) SpecialDelete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("servicebindings").
		Name(name).
		SubResource("delete").
		Body(options).
		Do().
		Error()
}
