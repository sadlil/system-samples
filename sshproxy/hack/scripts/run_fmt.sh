#!/usr/bin/env bash

set -o errexit
set -o nounset
set -o pipefail

echo 'Running gofmt and goimports from' $(pwd)
if [[ $(basename "$PWD") != sshproxy ]]; then
  echo 'This scripts needs to be run from repo root'
  exit 1
fi

# Run the formatter for root directory too.
goimports -w .
gofmt -s -w .

prefix=github.com/sadlil/system-samples/sshproxy
DIRs=$(go list ./... | grep -v vendor/)
for dir in ${DIRs}
    do
    	root=${dir#"$prefix/"}
    	if [[ "$root" != "$prefix" ]]; then
        goimports -w $root
        gofmt -s -w $root
      fi
done
