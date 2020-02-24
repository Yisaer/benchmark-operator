#!/usr/bin/env bash

# Copyright 2020 PingCAP, Inc.
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#     http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# See the License for the specific language governing permissions and
# limitations under the License.

set -e

ROOT=$(unset CDPATH && cd $(dirname "${BASH_SOURCE[0]}")/.. && pwd)
cd $ROOT

source $ROOT/hack/lib.sh

hack::ensure_kubectl
hack::ensure_kind

usage() {
  cat <<EOF
This script use kind to create Kubernetes cluster,about kind please refer: https://kind.sigs.k8s.io/
Before run this script,please ensure that docker have been installed.

Note that kind will merge config to your KUBECONFIG file, you could use a new KUBECONFIG file if you do not
want to mutate the defaulting one:

    export KUBECONFIG=${HOME}/.kube/kind-config-<cluster>

Options:
       -h,--help               prints the usage message
       -n,--name               name of the Kubernetes cluster,default value: kind
       -c,--nodeNum            the count of the cluster nodes,default value: 2
       -k,--k8sVersion         version of the Kubernetes cluster,default value: v1.16.3
       -v,--volumeNum          the volumes number of each kubernetes node,default value: 2

Environments:

    KUBECONFIG      kubectl config file that the kind config merged to
    HELM_VERSION    version of helm
    KUBECTL_VERSION version of kubectl
    KIND_VERSION    version of kind

Usage:
    $0 --name testCluster --nodeNum 4 --k8sVersion v1.12.9
EOF
}

while [[ $# -gt 0 ]]; do
  key="$1"

  case $key in
  -n | --name)
    clusterName="$2"
    shift
    shift
    ;;
  -c | --nodeNum)
    nodeNum="$2"
    shift
    shift
    ;;
  -k | --k8sVersion)
    k8sVersion="$2"
    shift
    shift
    ;;
  -v | --volumeNum)
    volumeNum="$2"
    shift
    shift
    ;;
  -h | --help)
    usage
    exit 0
    ;;
  *)
    echo "unknown option: $key"
    usage
    exit 1
    ;;
  esac
done

clusterName=${clusterName:-kind}
nodeNum=${nodeNum:-2}
k8sVersion=${k8sVersion:-v1.16.3}
volumeNum=${volumeNum:-2}

echo "clusterName: ${clusterName}"
echo "nodeNum: ${nodeNum}"
echo "k8sVersion: ${k8sVersion}"
echo "volumeNum: ${volumeNum}"

# check requirements
for requirement in $KIND_BIN docker; do
  echo "############ check ${requirement} ##############"
  if hash ${requirement} 2>/dev/null; then
    echo "${requirement} have installed"
  else
    echo "this script needs ${requirement}, please install ${requirement} first."
    exit 1
  fi
done

echo "############# start create cluster:[${clusterName}] #############"
workDir=${HOME}/kind/${clusterName}
mkdir -p ${workDir}

data_dir=${workDir}/data

echo "clean data dir: ${data_dir}"
if [ -d ${data_dir} ]; then
  rm -rf ${data_dir}
fi

configFile=${workDir}/kind-config.yaml

cat <<EOF >${configFile}
kind: Cluster
apiVersion: kind.x-k8s.io/v1alpha4
nodes:
- role: control-plane
EOF

for ((i = 0; i < ${nodeNum}; i++)); do
  mkdir -p ${data_dir}/worker${i}
  cat <<EOF >>${configFile}
- role: worker
  extraMounts:
EOF
  for ((k = 1; k <= ${volumeNum}; k++)); do
    mkdir -p ${data_dir}/worker${i}/vol${k}
    cat <<EOF >>${configFile}
  - containerPath: /mnt/disks/vol${k}
    hostPath: ${data_dir}/worker${i}/vol${k}
EOF
  done
done

echo "start to create k8s cluster"
$KIND_BIN create cluster --config ${configFile} --image kindest/node:${k8sVersion} --name=${clusterName}
echo "switch kube context to kind"
$KUBECTL_BIN config use-context kind-${clusterName}

echo "############# success create cluster:[${clusterName}] #############"

echo "To start using your cluster, run:"
echo "    kubectl config use-context kind-${clusterName}"
echo ""
echo <<EOF
NOTE: In kind, nodes run docker network and cannot access host network.
If you configured local HTTP proxy in your docker, images may cannot be pulled
because http proxy is inaccessible.

If you cannot remove http proxy settings, you can either whitelist image
domains in NO_PROXY environment or use 'docker pull <image> && $KIND_BIN load
docker-image <image>' command to load images into nodes.
EOF
