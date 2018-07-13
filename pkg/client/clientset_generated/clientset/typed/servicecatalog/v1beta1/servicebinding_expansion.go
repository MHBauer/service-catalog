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
	"github.com/golang/glog"
	"github.com/kubernetes-incubator/service-catalog/pkg/apis/servicecatalog/v1beta1"
)

// The ReferencesExtension interface allows setting the References
// to ServiceClasses and ServicePlans.
type ServiceBindingExpansion interface {
	SpecialDelete(serviceInstance *v1beta1.ServiceBinding) (*v1beta1.ServiceBinding, error)
}

func (c *serviceBindings) SpecialDelete(serviceBinding *v1beta1.ServiceBinding) (result *v1beta1.ServiceBinding, err error) {
	result = &v1beta1.ServiceBinding{}
	glog.Infof("request ns %q", c.ns)
	err = c.client.Delete().
		Namespace(c.ns).
		Resource("servicebindings").
		Name(serviceBinding.Name).
		SubResource("delete").
		Body(result).
		Do().
		Into(result)
	return
}
