#!/bin/sh

# delete stuff so it gets regenerated
# rm -rf bin pkg/client
# rebuild necessary
make build images

#start and stop infraastructure (etcd and apiserver)
contrib/hack/stop-server.sh
contrib/hack/start-server.sh


# set up the kubeconfig file
#mkdir -p /var/run/kubernetes
#docker cp apiserver:/var/run/kubernetes/apiserver.crt .apiserver.crt
contrib/hack/setup-kubectl.sh
contrib/hack/kubectl config set-credentials service-catalog-creds --username=admin --password=admin
contrib/hack/kubectl config set-cluster service-catalog-cluster --server=https://localhost:6443 --certificate-authority=$(pwd)/.var/run/kubernetes-service-catalog/apiserver.crt
contrib/hack/kubectl config set-context service-catalog-ctx --cluster=service-catalog-cluster --user=service-catalog-creds
contrib/hack/kubectl config use-context service-catalog-ctx

#compile and then run the client
bin/client  -v 10 -stderrthreshold 10 -logtostderr -kubeconfig .kube/config
