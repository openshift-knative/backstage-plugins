# DO NOT EDIT! Generated Dockerfile for vendor/knative.dev/pkg/apiextensions/storageversion/cmd/migrate.
ARG GO_BUILDER=registry.ci.openshift.org/openshift/release:rhel-8-release-golang-1.23-openshift-4.19
ARG GO_RUNTIME=registry.access.redhat.com/ubi8/ubi-minimal

FROM $GO_BUILDER as builder

WORKDIR /workspace
COPY . .

ENV CGO_ENABLED=1
ENV GOEXPERIMENT=strictfipsruntime

RUN go build -tags strictfipsruntime -o /usr/bin/main ./vendor/knative.dev/pkg/apiextensions/storageversion/cmd/migrate

FROM $GO_RUNTIME

ARG VERSION=knative-v1.16

COPY --from=builder /usr/bin/main /usr/bin/migrate
COPY LICENSE /licenses/

USER 65532

LABEL \
      com.redhat.component="openshift-serverless-1-backstage-plugins-migrate-rhel8-container" \
      name="openshift-serverless-1/backstage-plugins-migrate-rhel8" \
      version=$VERSION \
      summary="Red Hat OpenShift Serverless 1 Backstage Plugins Migrate" \
      maintainer="serverless-support@redhat.com" \
      description="Red Hat OpenShift Serverless 1 Backstage Plugins Migrate" \
      io.k8s.display-name="Red Hat OpenShift Serverless 1 Backstage Plugins Migrate" \
      io.k8s.description="Red Hat OpenShift Serverless Backstage Plugins Migrate" \
      io.openshift.tags="migrate"

ENTRYPOINT ["/usr/bin/migrate"]
