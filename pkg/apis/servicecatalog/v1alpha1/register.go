package v1alpha1

import (
	"k8s.io/kubernetes/pkg/api/v1"
	metav1 "k8s.io/kubernetes/pkg/apis/meta/v1"
	"k8s.io/kubernetes/pkg/runtime"

	"github.com/kubernetes-incubator/service-catalog/pkg/apis/servicecatalog"
)

const GroupName = servicecatalog.GroupName

var (
	SchemeBuilder = runtime.NewSchemeBuilder(addKnownTypes)
	AddToScheme   = SchemeBuilder.AddToScheme
)

// Adds the list of known types to api.Scheme.
func addKnownTypes(scheme *runtime.Scheme) error {
	scheme.AddKnownTypes(SchemeGroupVersion,
		&v1.ListOptions{},
		&v1.DeleteOptions{},
		&metav1.ExportOptions{},
		&metav1.GetOptions{},

		&Broker{},
	)
	return nil
}
