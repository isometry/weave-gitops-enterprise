#!/bin/sh

# This is a deliberately opinionated script for developing Weave GitOps Enterprise.
# Adapted from the original script in the Weave GitOps repository.
#
# WARN: This script is designed to be "turn it off and on again". It will delete
# the given kind cluster (if it exists) and recreate, installing everything from
# scratch.

export KIND_CLUSTER_NAME="${KIND_CLUSTER_NAME:-wge-dev}"

do_kind() {
    kind delete cluster --name "$KIND_CLUSTER_NAME"
    kind create cluster --name "$KIND_CLUSTER_NAME" --config "$(dirname "$0")/kind-cluster-with-extramounts.yaml"
}

do_capi(){
    EXP_CLUSTER_RESOURCE_SET=true clusterctl init --infrastructure docker
}

do_flux(){
    flux bootstrap github --owner="$GITHUB_USER" --repository=fleet-infra --branch=main --path=./clusters/management --personal
}

create_local_values_file(){
    envsubst < "$(dirname "$0")/dev-values-local.yaml.tpl" > "$(dirname "$0")/dev-values-local.yaml"
}

main() {
    do_kind
    do_capi
    do_flux
    create_local_values_file
}

main