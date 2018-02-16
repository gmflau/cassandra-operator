#!/usr/bin/env bash

set -o errexit
set -o nounset
set -o pipefail

echo GOPATH $GOPATH

DOCKER_REPO_ROOT="/Users/gilbertlau/go/src/github.com/gmflau/cassandra-operator"
IMAGE=${IMAGE:-"gcr.io/coreos-k8s-scale-testing/codegen"}

docker run --rm \
  -v "$PWD":"$DOCKER_REPO_ROOT" \
  -w "$DOCKER_REPO_ROOT" \
  "$IMAGE" \
  "./hack/k8s/codegen/codegen.sh" \
  "all" \
  "github.com/gmflau/cassandra-operator/pkg/generated" \
  "github.com/gmflau/cassandra-operator/pkg/apis" \
  "cassandra:v1beta2" \
  --go-header-file "./hack/k8s/codegen/boilerplate.go.txt" \
  $@


docker run --rm \
  -v "$PWD":"$DOCKER_REPO_ROOT" \
  -w "$DOCKER_REPO_ROOT" \
  -t \
  "$IMAGE" \
  "bash"
