#This makefile is used by ci-operator

CGO_ENABLED=0
GOOS=linux
BRANCH=
TEST=
IMAGE=
TEST_IMAGE_TAG ?= latest


test-e2e:
	./openshift/e2e-tests.sh
.PHONY: test-e2e

# Generate an aggregated knative release yaml file, as well as a CI file with replaced image references
generate-release:
	./openshift/release/generate-release.sh
.PHONY: generate-release
