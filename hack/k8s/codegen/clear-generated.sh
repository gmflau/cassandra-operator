#!/usr/bin/env bash

set -o errexit
set -o nounset
set -o pipefail

rm -rf /Users/gilbertlau/go/src/github.com/gmflau/cassandra-operator/pkg/apis/cassandra/v1beta2/zz_generated.deepcopy.go
rm -rf /Users/gilbertlau/go/src/github.com/gmflau/cassandra-operator/pkg/generated/*
