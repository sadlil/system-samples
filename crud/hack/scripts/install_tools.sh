#!/usr/bin/env bash

set -o errexit
set -o nounset
set -o pipefail

#!/usr/bin/env bash

set -o errexit
set -o nounset
set -o pipefail

function install_linux_tools() {
  sudo apt-get update -y

  PROTOC_ZIP=protoc-3.7.1-linux-x86_64.zip
  curl -OL https://github.com/protocolbuffers/protobuf/releases/download/v3.7.1/$PROTOC_ZIP
  sudo unzip -o $PROTOC_ZIP -d /usr/local bin/protoc
  sudo unzip -o $PROTOC_ZIP -d /usr/local 'include/*'
  rm -f $PROTOC_ZIP

  sudo apt-get install build-essential wget -y
  sudo apt-get install jq -y

  # binary will be $(go env GOPATH)/bin/golangci-lint
  curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | \
    sh -s -- -b $(go env GOPATH)/bin v1.31.0
}

function install_mac_tools() {
  if ! command -v brew &> /dev/null
  then
    echo "Couldn't find brew, you have to manually install brew. See https://brew.sh/."
    exit 1
  fi

  brew update

  brew install jq
  brew install cmake
  brew install protobuf
  brew install golangci/tap/golangci-lint
  brew upgrade golangci/tap/golangci-lint
}

function install_common_tools() {
  if ! command -v go version &> /dev/null
  then
    echo "Couldn't find go, you have to manually install golang. See https://golang.org/doc/install."
    exit 1
  fi

  make dep
  go install \
      github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway \
      github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2 \
      google.golang.org/protobuf/cmd/protoc-gen-go \
      google.golang.org/grpc/cmd/protoc-gen-go-grpc \
      github.com/envoyproxy/protoc-gen-validate \
      golang.org/x/tools/cmd/goimports \
      github.com/vektra/mockery/v2/... \
      github.com/go-bindata/go-bindata/... \
      github.com/stormcat24/protodep/...
}

os=$(uname)
echo 'Detected OS' ${os}'. Installing local tools.'

if [ "$os" = 'Darwin' ]; then
    install_mac_tools
elif [ "$os" = 'Linux' ]; then
    install_linux_tools
fi
install_common_tools
