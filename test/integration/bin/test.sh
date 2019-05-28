#!/bin/bash

set -e

# Install go
export GOROOT=$HOME/go-$GOVERSION
(cd ~ && curl -O https://dl.google.com/go/go${GOVERSION}.linux-amd64.tar.gz)
mkdir ~/go-${GOVERSION} && tar xf ~/go${GOVERSION}.linux-amd64.tar.gz -C $GOROOT --strip-components 1

# Initialise Kerberos
KERBEROS_IP=$(jq -r '.public_ips.value[0]' /tmp/terraform_output.json)
$(dirname $0)/../kerberos/install_kerberos.sh "$KERBEROS_IP"

# Install Kubectl
sudo cp /tmp/workspace/kubectl /usr/bin/kubectl

# Run integration tests
IMGTAG=$(./tools/image-tag)

docker login -u="$DOCKER_USER" -p="$DOCKER_PASSWORD" quay.io
export PATH=$GOROOT/bin:$PATH
# Work around for broken docker package in RHEL
# See https://github.com/weaveworks/wks/issues/235
export DOCKER_VERSION='1.13.1-75*'
go test -failfast -v -timeout 1h ./test/integration/test -args -run.interactive -cmd /tmp/workspace/cmd/wksctl/wksctl -tags.wks-k8s-krb5-server=$IMGTAG -tags.wks-mock-authz-server=$IMGTAG