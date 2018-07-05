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

ROOT="$(cd "$(dirname "${BASH_SOURCE[0]}")/../.." && pwd)"
export PATH=${ROOT}/contrib/hack:${PATH}

trap cleanup EXIT

function cleanup {
    rc=$?
	echo Cleaning up
	stop-server.sh
	exit $rc
}

start-server.sh

# Kubectl needs to be configured with the current cluster
# setup. Kubectl was initially configured in a different script and
# the port mapping may have changed by the time we get here.
PORT=$(docker port etcd-svc-cat 443 | sed "s/.*://")
echo $PORT
D_HOST=${DOCKER_HOST:-localhost}
D_HOST=${D_HOST#*//}   # remove leading proto://
D_HOST=${D_HOST%:*}    # remove trailing port #
NO_TTY=1 kubectl config set-cluster service-catalog-cluster --server=https://${D_HOST}:${PORT} --certificate-authority=/var/run/kubernetes-service-catalog/apiserver.crt


# create a few resources
set -x
NO_TTY=1 kubectl config view
# create the binding
NO_TTY=1 kubectl create -f contrib/examples/apiserver/binding.yaml
# make sure it's still there
NO_TTY=1 kubectl get servicebinding test-binding --namespace test-ns -o yaml
# call delete
NO_TTY=1 kubectl delete -v 9 -f contrib/examples/apiserver/binding.yaml
# check that we have no deletion timestamp
NO_TTY=1 kubectl get servicebinding test-binding --namespace test-ns -o yaml
curl -k -v -XGET  -H "Authorization: Basic YWRtaW46YWRtaW4=" -H "User-Agent: kubectl/v1.6.6 (linux/amd64) kubernetes/7fa1c17" -H "Accept: application/json" -H "Content-Type: application/json" https://localhost:${PORT}/apis/servicecatalog.k8s.io/v1beta1/namespaces/test-ns/servicebindings/test-binding/status
# do the actual delete on the special resource
curl -k -v -XDELETE  -H "Authorization: Basic YWRtaW46YWRtaW4=" -H "User-Agent: kubectl/v1.6.6 (linux/amd64) kubernetes/7fa1c17" -H "Accept: application/json" -H "Content-Type: application/json" https://localhost:${PORT}/apis/servicecatalog.k8s.io/v1beta1/namespaces/test-ns/servicebindings/test-binding/delete
# see that the deletion timestamp is set now
NO_TTY=1 kubectl get servicebinding test-binding --namespace test-ns -o yaml

set +x
