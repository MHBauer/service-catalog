package main

import (
	"os"

	"github.com/golang/glog"

	"k8s.io/kubernetes/pkg/util/logs"

	"github.com/kubernetes-incubator/service-catalog/pkg/cmd/server"
	// commented out until I know what this does
	// install all APIs
	// force REGISTRATION of packages we'll later rely upon
	_ "k8s.io/kubernetes/cmd/kubernetes-discovery/pkg/apis/apiregistration/install"
	_ "k8s.io/kubernetes/cmd/kubernetes-discovery/pkg/apis/apiregistration/validation"
	_ "k8s.io/kubernetes/cmd/kubernetes-discovery/pkg/client/clientset_generated/internalclientset"
	_ "k8s.io/kubernetes/cmd/kubernetes-discovery/pkg/client/listers/apiregistration/internalversion"
	_ "k8s.io/kubernetes/cmd/kubernetes-discovery/pkg/client/listers/apiregistration/v1alpha1"

	_ "k8s.io/kubernetes/pkg/api/install"
)

func main() {
	logs.InitLogs()
	// make sure we print all the logs while shutting down.
	defer logs.FlushLogs()

	cmd := server.NewCommandServer(os.Stdout)
	if err := cmd.Execute(); err != nil {
		glog.Errorln(err)
		os.Exit(1)
	}
}
