#!/usr/bin/env bash

set -euo pipefail

source $(dirname $0)/resolve.sh

root_dir=$(dirname $0)/../..

release=$(yq '.project.tag' openshift/project.yaml)
release=${release/knative-/}

echo "Release: $release"

./openshift/generate.sh

artifacts_dir="openshift/release/artifacts"
rm -rf $artifacts_dir
mkdir -p $artifacts_dir

image_prefix="registry.ci.openshift.org/openshift/knative-${release}:knative-backstage-plugins-"
tag=""

plugins="${artifacts_dir}/backstage-plugins-eventmesh-backend.yaml"

# EventMesh Backend resources
resolve_resources backends/config/100-eventmesh "${plugins}" "$image_prefix" "$tag"
