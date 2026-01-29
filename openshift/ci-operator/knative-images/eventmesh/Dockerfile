# DO NOT EDIT! Generated Dockerfile for backends/cmd/eventmesh.
ARG GO_BUILDER=registry.ci.openshift.org/openshift/release:rhel-9-release-golang-1.24-openshift-4.20
ARG GO_RUNTIME=registry.access.redhat.com/ubi9/ubi-minimal

FROM $GO_BUILDER as builder

WORKDIR /workspace
COPY . .

ENV CGO_ENABLED=1
ENV GOEXPERIMENT=strictfipsruntime
ENV GOFLAGS='-mod=mod'

RUN go build -tags strictfipsruntime -o /usr/bin/main ./backends/cmd/eventmesh

FROM $GO_RUNTIME

ARG VERSION=knative-v1.21

COPY --from=builder /usr/bin/main /usr/bin/eventmesh
COPY LICENSE /licenses/

USER 65532

LABEL \
      com.redhat.component="openshift-serverless-1-kn-backstage-plugins-eventmesh-rhel9-container" \
      name="openshift-serverless-1/kn-backstage-plugins-eventmesh-rhel9" \
      version=$VERSION \
      summary="Red Hat OpenShift Serverless 1 Backstage Plugins Eventmesh" \
      maintainer="serverless-support@redhat.com" \
      description="Red Hat OpenShift Serverless 1 Backstage Plugins Eventmesh" \
      io.k8s.display-name="Red Hat OpenShift Serverless 1 Backstage Plugins Eventmesh" \
      io.k8s.description="Red Hat OpenShift Serverless Backstage Plugins Eventmesh" \
      io.openshift.tags="eventmesh" \
      vendor="Red Hat, Inc." \
      release=$VERSION \
      cpe="cpe:/a:redhat:openshift_serverless:1.41::el9"

ENTRYPOINT ["/usr/bin/eventmesh"]
