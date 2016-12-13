package install

import (
	"k8s.io/kubernetes/pkg/apimachinery/announced"
	"k8s.io/kubernetes/pkg/util/sets"

	"github.com/kubernetes-incubator/service-catalog/pkg/apis/servicecatalog"
	"github.com/kubernetes-incubator/service-catalog/pkg/apis/servicecatalog/v1alpha1"
)

func init() {
	if err := announced.NewGroupMetaFactory(
		&announced.GroupMetaFactoryArgs{
			GroupName:                  servicecatalog.GroupName,
			VersionPreferenceOrder:     []string{v1alpha1.SchemeGroupVersion.Version},
			ImportPrefix:               "k8s.io/kubernetes/pkg/apis/extensions",
			RootScopedKinds:            sets.NewString("service-catalog"),
			AddInternalObjectsToScheme: servicecatalog.AddToScheme,
		},
		announced.VersionToSchemeFunc{
			v1alpha1.SchemeGroupVersion.Version: v1alpha1.AddToScheme,
		},
	).Announce().RegisterAndEnable(); err != nil {
		panic(err)
	}
}
