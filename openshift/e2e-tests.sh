#!/usr/bin/env bash

# shellcheck disable=SC1090
set -Eeuox pipefail

env

failed=0

echo "TODO: going to run e2e tests"

(( failed )) && dump_cluster_state

(( failed )) && exit 1

success
