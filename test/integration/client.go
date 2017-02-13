package main

import (
	"flag"
	"fmt"
	"path/filepath"
	"time"

	"github.com/golang/glog"

	"k8s.io/kubernetes/pkg/api/v1"
	"k8s.io/kubernetes/pkg/client/unversioned/clientcmd"

	// TODO: fix this upstream
	// we shouldn't have to install things to use our own generated client.

	// avoid error `servicecatalog/v1alpha1 is not enabled`
	_ "github.com/kubernetes-incubator/service-catalog/pkg/apis/servicecatalog/install"
	// avoid error `no kind is registered for the type v1.ListOptions`
	_ "k8s.io/kubernetes/pkg/api/install"

	// our versioned types
	"github.com/kubernetes-incubator/service-catalog/pkg/apis/servicecatalog/v1alpha1"
	// our versioned client
	servicecatalog "github.com/kubernetes-incubator/service-catalog/pkg/client/clientset_generated/clientset"
)

var (
	kubeconfig = flag.String("kubeconfig", ".kube/config", "absolute path to the kubeconfig file")
)

func main() {
	flag.Parse()

	absKConfig, err := filepath.Abs(*kubeconfig)
	if nil != err {
		glog.Fatal("could not load kubeconfig")
	}

	// uses the current context in kubeconfig
	config, err := clientcmd.BuildConfigFromFlags("", absKConfig)
	if err != nil {
		panic(err.Error())
	}
	// creates the clientset
	clientset, err := servicecatalog.NewForConfig(config)
	if err != nil {
		glog.Fatal(err)
	}

	// create a broker
	broker, err := clientset.Servicecatalog().Brokers().Create(
		&v1alpha1.Broker{ObjectMeta: v1.ObjectMeta{Name: "test-broker"},
			Spec: v1alpha1.BrokerSpec{URL: "broker needs a url"}})
	glog.Info(broker, err)

	// list
	for {
		brokers, err := clientset.Servicecatalog().Brokers().List(v1.ListOptions{})
		if err != nil {
			glog.Flush()
			glog.Fatal(err)
		}
		fmt.Printf("There are %d brokers in the cluster\n", len(brokers.Items))
		glog.Infof("There are %d brokers in the cluster\n", len(brokers.Items))
		time.Sleep(10 * time.Second)
		if len(brokers.Items) > 0 {
			break
		}
	}
}
