package main

import (
	"os"
	// "runtime"

	// commented out until I know what this does
	// "k8s.io/kubernetes/pkg/util/logs"

	"github.com/kubernetes-incubator/service-catalog/pkg/cmd/server"
	// commented out until I know what this does
	// install all APIs
	// _ "github.com/openshift/kube-aggregator/pkg/apis/apifederation/install"
	// _ "k8s.io/kubernetes/pkg/api/install"
)

func main() {
	// commented out until I know what this does
	// logs.InitLogs()
	// defer logs

	// if len(os.Getenv("GOMAXPROCS")) == 0 {
	// 	runtime.GOMAXPROCS(runtime.NumCPU())
	// }

	cmd := server.NewCommandServer(os.Stdout)
	if err := cmd.Execute(); err != nil {
		os.Exit(1)
	}
}
