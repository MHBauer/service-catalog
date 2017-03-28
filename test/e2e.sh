#!/bin/bash

# Copyright 2017 The Kubernetes Authors.
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#     http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.

set -o errexit
set -o nounset
set -o pipefail

set -o xtrace

# this script resides in the `test/` folder at the root of the project
# realpath doesn't exist on trusty ubuntu which is what travis runs in...
#KUBE_ROOT=$(realpath $(dirname "${BASH_SOURCE}")/../vendor/k8s.io/kubernetes)
# script should be running in the repo root, so access vendor directly
KUBE_ROOT=$(pwd -P)/vendor/k8s.io/kubernetes
source "${KUBE_ROOT}/hack/lib/init.sh"

cleanup () {
    docker rm --force svc-cat-lk
    docker rm --force $(docker ps -q)
}

KUBE_CLIENT_VERSION='v1.6.0-beta.4'
kubectl () {
    docker run --rm -it --net=host --volume $(pwd):$(pwd) \
           gcr.io/google_containers/hyperkube-amd64:${KUBE_CLIENT_VERSION} \
           /hyperkube kubectl "$@"
}


# uncomment this at the end to clean everything up
#trap cleanup EXIT

sudo mkdir -p /var/lib/kubelet
sudo mount
sudo mount --make-shared /
#start the localkube in docker
docker run --name svc-cat-lk -d -it --privileged --volume=/:/rootfs:ro --volume=/sys:/sys:rw     --volume=/var/lib/docker/:/var/lib/docker:rw --volume=/var/lib/dockershim/sandbox:/var/lib/dockershim/sandbox:rw     --volume=/var/lib/kubelet/:/var/lib/kubelet:rw,shared     --volume=/var/run:/var/run:rw     --net=host -P brahmaroutu/localkube ./localkube --containerized

kube::util::wait_for_url http://localhost:8080/healthz

kubectl version
kubectl get all --all-namespaces
kubectl create -f https://raw.githubusercontent.com/kubernetes/minikube/master/deploy/addons/kube-dns/kube-dns-rc.yaml
kubectl create -f https://raw.githubusercontent.com/kubernetes/minikube/master/deploy/addons/kube-dns/kube-dns-svc.yaml

kubectl get all --all-namespaces

kubectl create -f $(pwd)/test/e2e/install-broker.yaml
kubectl create -f $(pwd)/test/e2e/install-apiserver.yaml
kubectl create -f $(pwd)/test/e2e/install-controller.yaml

kubectl create ns cb-ns
kubectl create -f $(pwd)/test/e2e/install-secret.yaml

kubectl get all --all-namespaces

apiPort=$(kubectl get  svc/apiserver-service -o jsonpath='{.spec.ports[0].nodePort}') 

kube::util::wait_for_url http://localhost:${apiPort}/healthz

kubectl --server=localhost:${apiPort} create -f $(pwd)/test/e2e/cbbroker.yaml
kubectl --server=localhost:${apiPort} create -f $(pwd)/test/e2e/cbinstance.yaml
kubectl --server=localhost:${apiPort} create -f $(pwd)/test/e2e/cbbinding.yaml
kubectl --server=localhost:${apiPort} get broker,serviceclass -o yaml
kubectl --server=localhost:${apiPort} -n cb-ns get instance,binding -o yaml

kubectl get all --all-namespaces -o yaml

# by the end this should have a secret in it
kubectl get secret -n cb-ns 


