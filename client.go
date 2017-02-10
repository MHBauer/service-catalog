package main

import (
	"flag"
	"fmt"
	"time"

	"github.com/golang/glog"

	// metav1 "k8s.io/kubernetes/pkg/apis/meta/v1"

	"k8s.io/kubernetes/pkg/api/v1"
	"k8s.io/kubernetes/pkg/client/unversioned/clientcmd"

	_ "github.com/kubernetes-incubator/service-catalog/pkg/apis/servicecatalog/install"
	_ "k8s.io/kubernetes/pkg/api/install"

	servicecatalog "github.com/kubernetes-incubator/service-catalog/pkg/client/clientset_generated/clientset"
)

var (
	kubeconfig = flag.String("kubeconfig", ".kube/config", "absolute path to the kubeconfig file")
)

func main() {
	flag.Parse()
	// uses the current context in kubeconfig
	config, err := clientcmd.BuildConfigFromFlags("", *kubeconfig)
	// generates a k8s.io/client-go/1.5/rest.Config
	if err != nil {
		panic(err.Error())
	}
	// creates the clientset
	// do I need to fix the generator?
	// wants a k8s.io/kubernetes/pkg/client/restclient/rest.ClientConfig
	//config = restclient.RESTClientFor(config)
	clientset, err := servicecatalog.NewForConfig(config)
	if err != nil {
		glog.Fatal(err)
	}

	for {

		brokers, err := clientset.Servicecatalog().Brokers().List(v1.ListOptions{})
		if err != nil {
			glog.Flush()
			glog.Fatal(err)
		}
		fmt.Printf("There are %d brokers in the cluster\n", len(brokers.Items))
		time.Sleep(10 * time.Second)
	}
}
