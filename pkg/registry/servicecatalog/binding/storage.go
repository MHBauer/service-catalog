/*
Copyright 2016 The Kubernetes Authors.

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

package binding

import (
	"context"
	"errors"
	"fmt"

	"github.com/golang/glog"
	scmeta "github.com/kubernetes-incubator/service-catalog/pkg/api/meta"
	"github.com/kubernetes-incubator/service-catalog/pkg/apis/servicecatalog"
	"github.com/kubernetes-incubator/service-catalog/pkg/registry/servicecatalog/server"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/fields"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
	genericapirequest "k8s.io/apiserver/pkg/endpoints/request"
	"k8s.io/apiserver/pkg/registry/generic"
	"k8s.io/apiserver/pkg/registry/generic/registry"
	"k8s.io/apiserver/pkg/registry/rest"
	"k8s.io/apiserver/pkg/storage"
	storeerr "k8s.io/apiserver/pkg/storage/errors"
)

var (
	errNotAServiceBinding = errors.New("not a binding")
)

// NewSingular returns a new shell of a service binding, according to the given namespace and
// name
func NewSingular(ns, name string) runtime.Object {
	return &servicecatalog.ServiceBinding{
		TypeMeta: metav1.TypeMeta{
			Kind: "ServiceBinding",
		},
		ObjectMeta: metav1.ObjectMeta{
			Namespace: ns,
			Name:      name,
		},
	}
}

// EmptyObject returns an empty binding
func EmptyObject() runtime.Object {
	return &servicecatalog.ServiceBinding{}
}

// NewList returns a new shell of a binding list
func NewList() runtime.Object {
	return &servicecatalog.ServiceBindingList{
		TypeMeta: metav1.TypeMeta{
			Kind: "ServiceBindingList",
		},
		Items: []servicecatalog.ServiceBinding{},
	}
}

// CheckObject returns a non-nil error if obj is not a binding object
func CheckObject(obj runtime.Object) error {
	_, ok := obj.(*servicecatalog.ServiceBinding)
	if !ok {
		return errNotAServiceBinding
	}
	return nil
}

// Match determines whether an ServiceInstance matches a field and label
// selector.
func Match(label labels.Selector, field fields.Selector) storage.SelectionPredicate {
	return storage.SelectionPredicate{
		Label:    label,
		Field:    field,
		GetAttrs: GetAttrs,
	}
}

// toSelectableFields returns a field set that represents the object for matching purposes.
func toSelectableFields(binding *servicecatalog.ServiceBinding) fields.Set {
	// If you add a new selectable field, you also need to modify
	// pkg/apis/servicecatalog/v1beta1/conversion[_test].go
	objectMetaFieldsSet := generic.ObjectMetaFieldsSet(&binding.ObjectMeta, true)

	specFieldSet := make(fields.Set, 1)

	if binding.Spec.ExternalID != "" {
		specFieldSet["spec.externalID"] = binding.Spec.ExternalID
	}

	return generic.MergeFieldsSets(objectMetaFieldsSet, specFieldSet)
}

// GetAttrs returns labels and fields of a given object for filtering purposes.
func GetAttrs(obj runtime.Object) (labels.Set, fields.Set, bool, error) {
	binding, ok := obj.(*servicecatalog.ServiceBinding)
	if !ok {
		return nil, nil, false, fmt.Errorf("given object is not a ServiceBinding")
	}
	return labels.Set(binding.ObjectMeta.Labels), toSelectableFields(binding), binding.Initializers != nil, nil
}

// NewStorage creates a new rest.Storage responsible for accessing ServiceBinding
// resources
func NewStorage(opts server.Options) (rest.Storage, rest.Storage, rest.Storage, error) {
	prefix := "/" + opts.ResourcePrefix()

	storageInterface, dFunc := opts.GetStorage(
		&servicecatalog.ServiceBinding{},
		prefix,
		bindingRESTStrategies,
		NewList,
		nil,
		storage.NoTriggerPublisher,
	)

	store := registry.Store{
		NewFunc: EmptyObject,
		// NewListFunc returns an object capable of storing results of an etcd list.
		NewListFunc: NewList,
		KeyRootFunc: opts.KeyRootFunc(),
		KeyFunc:     opts.KeyFunc(true),
		// Retrieve the name field of the resource.
		ObjectNameFunc: func(obj runtime.Object) (string, error) {
			return scmeta.GetAccessor().Name(obj)
		},
		// Used to match objects based on labels/fields for list.
		PredicateFunc: Match,
		// DefaultQualifiedResource should always be plural
		DefaultQualifiedResource: servicecatalog.Resource("servicebindings"),

		CreateStrategy:          bindingRESTStrategies,
		UpdateStrategy:          bindingRESTStrategies,
		DeleteStrategy:          bindingRESTStrategies,
		EnableGarbageCollection: true,

		Storage:     storageInterface,
		DestroyFunc: dFunc,
	}

	options := &generic.StoreOptions{RESTOptions: opts.EtcdOptions.RESTOptions, AttrFunc: GetAttrs}
	if err := store.CompleteWithOptions(options); err != nil {
		panic(err) // TODO: Propagate error up
	}

	interceptedStore := DeleteInterceptREST{&store}
	deleteStore := DeleteREST{&store}

	statusStore := store
	statusStore.UpdateStrategy = bindingStatusUpdateStrategy

	return &interceptedStore, &StatusREST{&statusStore}, &deleteStore, nil
}

// StatusREST defines the REST operations for the status subresource via
// implementation of various rest interfaces.  It supports the http verbs GET,
// PATCH, and PUT.
type StatusREST struct {
	store *registry.Store
}

var (
	_ rest.Storage = &StatusREST{}
	_ rest.Getter  = &StatusREST{}
	_ rest.Updater = &StatusREST{}
)

// New returns a new ServiceBinding.
func (r *StatusREST) New() runtime.Object {
	return EmptyObject()
}

// Get retrieves the object from the storage. It is required to support Patch
// and to implement the rest.Getter interface.
func (r *StatusREST) Get(ctx context.Context, name string, options *metav1.GetOptions) (runtime.Object, error) {
	return r.store.Get(ctx, name, options)
}

// Update alters the status subset of an object and implements the rest.Updater
// interface.
func (r *StatusREST) Update(ctx context.Context, name string, objInfo rest.UpdatedObjectInfo, createValidation rest.ValidateObjectFunc, updateValidation rest.ValidateObjectUpdateFunc) (runtime.Object, bool, error) {
	return r.store.Update(ctx, name, objInfo, createValidation, updateValidation)
}

type DeleteInterceptREST struct {
	*registry.Store
}

var (
	_ rest.Storage         = &DeleteInterceptREST{}
	_ rest.StandardStorage = &DeleteInterceptREST{}
	_ rest.GracefulDeleter = &DeleteInterceptREST{}
)

func (r *DeleteInterceptREST) New() runtime.Object {
	return EmptyObject()
}

// Delete avoids standard kube delete functionality
func (r *DeleteInterceptREST) Delete(ctx context.Context, name string, options *metav1.DeleteOptions) (runtime.Object, bool, error) {
	key, err := r.KeyFunc(ctx, name)
	if err != nil {
		return nil, false, err
	}
	obj := r.NewFunc()
	// seems bizzare to pull this out of the context by the time we're here
	var qualifiedResource schema.GroupResource
	if info, ok := genericapirequest.RequestInfoFrom(ctx); ok {
		qualifiedResource = schema.GroupResource{Group: info.APIGroup, Resource: info.Resource}
		// these kind of have to be the same...
		// TODO investigate upstream code history
		// glog.Infof("is this ever set %+v", qualifiedResource)
		// glog.Infof("default could is %+v", r.DefaultQualifiedResource)
	} else {
		qualifiedResource = r.DefaultQualifiedResource
	}

	if err := r.Storage.Get(ctx, key, "", obj, false); err != nil {
		return nil, false, storeerr.InterpretDeleteError(err, qualifiedResource, name)
	}

	if binding, ok := obj.(*servicecatalog.ServiceBinding); ok {
		// write object with delete field back to storage

		var preconditions storage.Preconditions
		var lastExisting runtime.Object
		err = r.Storage.GuaranteedUpdate(
			ctx,
			key,
			binding,
			false, // ignoreNotFound
			&preconditions,
			storage.SimpleUpdate(func(existing runtime.Object) (runtime.Object, error) {

				if binding, ok := existing.(*servicecatalog.ServiceBinding); ok {
					glog.Errorf("binding %+v", binding.Spec.SecretName)
				}
				lastExisting = existing

				binding.Spec.SecretName = servicecatalog.SecretNameKey

				return binding, nil
			}),
		)
		switch err {
		case nil:
			// If we are here, the registry supports grace period mechanism and
			// we are intentionally delete gracelessly. In this case, we may
			// enter a race with other k8s components. If other component wins
			// the race, the object will not be found, and we should tolerate
			// the NotFound error. See
			// https://github.com/kubernetes/kubernetes/issues/19403 for
			// details.
			return lastExisting, false, nil
		default:
			return lastExisting, false, storeerr.InterpretUpdateError(err, qualifiedResource, name)
		}
	}
	glog.Errorf("we got a non-binding to delete somehow %+v", obj)
	return nil, false, fmt.Errorf("we didn't get a binding to delete")
}

// DeleteREST is type that only supports deleting
type DeleteREST struct {
	delete *registry.Store
}

var (
	_ rest.Storage         = &DeleteREST{}
	_ rest.GracefulDeleter = &DeleteREST{}
	// we're explicitly not a standard storage and only support DELETE
	// _ rest.StandardStorage = &DeleteREST{}
)

// New is necessary to implement storage
func (r *DeleteREST) New() runtime.Object {
	return EmptyObject()
}

// Delete is a passthrough to the underlying delete logic
func (r *DeleteREST) Delete(ctx context.Context, name string, options *metav1.DeleteOptions) (runtime.Object, bool, error) {
	return r.delete.Delete(ctx, name, options)
}
