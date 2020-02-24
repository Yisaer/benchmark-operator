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

if [ -z "$ROOT" ]; then
  echo "error: ROOT should be initialized"
  exit 1
fi

OS=$(go env GOOS)
ARCH=$(go env GOARCH)
OUTPUT=${ROOT}/output
OUTPUT_BIN=${OUTPUT}/bin
KUBECTL_VERSION=${KUBECTL_VERSION:-1.12.10}
KUBECTL_BIN=$OUTPUT_BIN/kubectl
KIND_VERSION=${KIND_VERSION:-0.7.0}
KIND_BIN=$OUTPUT_BIN/kind
KUSTOMIZE_BIN=$OUTPUT_BIN/kustomize


test -d "$OUTPUT_BIN" || mkdir -p "$OUTPUT_BIN"


function hack::verify_kubectl() {
  if test -x "$KUBECTL_BIN"; then
    [[ "$($KUBECTL_BIN version --client --short | grep -o -E '[0-9]+\.[0-9]+\.[0-9]+')" == "$KUBECTL_VERSION" ]]
    return
  fi
  return 1
}

function hack::ensure_kubectl() {
  if hack::verify_kubectl; then
    return 0
  fi
  echo "Installing kubectl v$KUBECTL_VERSION..."
  tmpfile=$(mktemp)
  trap "test -f $tmpfile && rm $tmpfile" RETURN
  curl --retry 10 -L -o $tmpfile https://storage.googleapis.com/kubernetes-release/release/v${KUBECTL_VERSION}/bin/${OS}/${ARCH}/kubectl
  mv $tmpfile $KUBECTL_BIN
  chmod +x $KUBECTL_BIN
}


function hack::verify_kind() {
  if test -x "$KIND_BIN"; then
    [[ "$($KIND_BIN --version 2>&1 | cut -d ' ' -f 3)" == "$KIND_VERSION" ]]
    return
  fi
  return 1
}

function hack::ensure_kind() {
  if hack::verify_kind; then
    return 0
  fi
  echo "Installing kind v$KIND_VERSION..."
  tmpfile=$(mktemp)
  trap "test -f $tmpfile && rm $tmpfile" RETURN
  curl --retry 10 -L -o $tmpfile https://github.com/kubernetes-sigs/kind/releases/download/v${KIND_VERSION}/kind-$(uname)-amd64
  mv $tmpfile $KIND_BIN
  chmod +x $KIND_BIN
}

function hack::verify_kustomize() {
  if test -x "$KUSTOMIZE_BIN"; then
    return 0
  fi
  return 1
}

function hack::ensure_kustomize() {
  if hack::verify_kustomize; then
    return 0
  fi
  curl -s "https://raw.githubusercontent.com/kubernetes-sigs/kustomize/master/hack/install_kustomize.sh" | bash > /dev/null
  mv ./kustomize $KUSTOMIZE_BIN
  chmod +x $KUSTOMIZE_BIN
}

# hack::version_ge "$v1" "$v2" checks whether "v1" is greater or equal to "v2"
function hack::version_ge() {
  [ "$(printf '%s\n' "$1" "$2" | sort -V | head -n1)" = "$2" ]
}

# Usage:
#
#	hack::wait_for_success 120 5 "cmd arg1 arg2 ... argn"
#
# Returns 0 if the shell command get output, 1 otherwise.
# From https://github.com/kubernetes/kubernetes/blob/v1.17.0/hack/lib/util.sh#L70
function hack::wait_for_success() {
  local wait_time="$1"
  local sleep_time="$2"
  local cmd="$3"
  while [ "$wait_time" -gt 0 ]; do
    if eval "$cmd"; then
      return 0
    else
      sleep "$sleep_time"
      wait_time=$((wait_time - sleep_time))
    fi
  done
  return 1
}
