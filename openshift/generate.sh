#!/usr/bin/env bash

set -euo pipefail

repo_root_dir=$(dirname "$(realpath "${BASH_SOURCE[0]}")")/..

GOFLAGS='' go install github.com/openshift-knative/hack/cmd/generate@latest


$(go env GOPATH)/bin/generate \
  --root-dir "${repo_root_dir}" \
  --generators dockerfile \
  --dockerfile-image-builder-fmt "registry.ci.openshift.org/openshift/release:rhel-8-release-golang-%s-openshift-4.16"
