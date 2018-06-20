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

package rest

import (
	"strings"
	"testing"

	"github.com/kubernetes-incubator/service-catalog/pkg/registry/servicecatalog/server"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apiserver/pkg/registry/generic"
	"k8s.io/apiserver/pkg/registry/rest"
	serverstorage "k8s.io/apiserver/pkg/server/storage"
	"k8s.io/apiserver/pkg/storage"
	"k8s.io/apiserver/pkg/storage/storagebackend"
	"k8s.io/apiserver/pkg/storage/storagebackend/factory"
)

type GetRESTOptionsHelper struct {
	retStorageInterface storage.Interface
	retDestroyFunc      func()
}

func (g GetRESTOptionsHelper) GetRESTOptions(resource schema.GroupResource) (generic.RESTOptions, error) {
	return generic.RESTOptions{
		ResourcePrefix: resource.Group + "/" + resource.Resource,
		StorageConfig:  &storagebackend.Config{},
		Decorator: generic.StorageDecorator(func(
			config *storagebackend.Config,
			objectType runtime.Object,
			resourcePrefix string,
			keyFunc func(obj runtime.Object) (string, error),
			newListFunc func() runtime.Object,
			getAttrsFunc storage.AttrFunc,
			trigger storage.TriggerPublisherFunc,
		) (storage.Interface, factory.DestroyFunc) {
			return g.retStorageInterface, g.retDestroyFunc
		})}, nil
}

func testRESTOptionsGetter(
	retStorageInterface storage.Interface,
	retDestroyFunc func(),
) generic.RESTOptionsGetter {
	return GetRESTOptionsHelper{retStorageInterface, retDestroyFunc}
}

func TestV1Beta1Storage(t *testing.T) {
	provider := StorageProvider{
		DefaultNamespace: "test-default",
		StorageType:      server.StorageTypeEtcd,
		RESTClient:       nil,
	}
	configSource := serverstorage.NewResourceConfig()
	roGetter := testRESTOptionsGetter(nil, func() {})
	storageMap, err := provider.v1beta1Storage(configSource, roGetter)
	if err != nil {
		t.Fatalf("error getting v1beta1 storage (%s)", err)
	}

	storages := [...]string{
		"clusterservicebrokers",
		"clusterservicebrokers/status",
		"clusterserviceclasses",
		"clusterserviceclasses/status",
		"clusterserviceplans",
		"clusterserviceplans/status",
		"serviceinstances",
		"serviceinstances/status",
		"serviceinstances/reference",
		"servicebindings",
		"servicebindings/status",
	}

	for _, s := range storages {
		storage, storageExists := storageMap[s]
		if !storageExists {
			t.Fatalf("no storage found for %q", s)
		}
		t.Logf("%q %+v", s, storage)
		if _, isStandardStorage := storage.(rest.Storage); !isStandardStorage {
			t.Errorf("%q isn't even a storage", s)
			continue
		}

		if strings.Contains(s, "status") {
			// check that status is only GET & UPDATE
			if _, isStandardStorage := storage.(rest.Getter); !isStandardStorage {
				t.Errorf("not compliant to getter interface for %q", s)
			}
			if _, isStandardStorage := storage.(rest.Updater); !isStandardStorage {
				t.Errorf("not compliant to updater interface for %q", s)
			}
			continue
		}

		if _, isStandardStorage := storage.(rest.Creater); !isStandardStorage {
			t.Errorf("not compliant to creater interface for %q", s)
		}
		if _, isStandardStorage := storage.(rest.Lister); !isStandardStorage {
			t.Errorf("not compliant to lister interface for %q", s)
		}
		if _, isStandardStorage := storage.(rest.GracefulDeleter); !isStandardStorage {
			t.Errorf("not compliant to graceful deleter interface for %q", s)
		}
		if _, isStandardStorage := storage.(rest.CollectionDeleter); !isStandardStorage {
			t.Errorf("not compliant to collection deleter interface for %q", s)
		}
		if _, isStandardStorage := storage.(rest.Watcher); !isStandardStorage {
			t.Errorf("not compliant to watcher interface for %q", s)
		}
		if _, isStandardStorage := storage.(rest.StandardStorage); !isStandardStorage {
			t.Errorf("not compliant to standardStorage interface for %q", s)
		}
	}

}
