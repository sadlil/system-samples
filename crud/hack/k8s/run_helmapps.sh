#!/usr/bin/env bash

set -o errexit
set -o nounset
set -o pipefail

helm install -f hack/k8s/helmapps/mysql.bitnami.yaml --namespace=crud todo-mysql bitnami/mysql
helm install -f hack/k8s/helmapps/redis.bitnami.yaml --namespace=crud todo-redis bitnami/redis
