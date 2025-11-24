# DO NOT EDIT! Generated Dockerfile for backends/cmd/eventmesh.
ARG GO_BUILDER=registry.ci.openshift.org/openshift/release:rhel-8-release-golang-1.22-openshift-4.17
ARG GO_RUNTIME=registry.access.redhat.com/ubi8/ubi-minimal

FROM $GO_BUILDER as builder

WORKDIR /workspace
COPY . .

ENV CGO_ENABLED=1
ENV GOEXPERIMENT=strictfipsruntime

RUN go build -tags strictfipsruntime -o /usr/bin/main ./backends/cmd/eventmesh

FROM $GO_RUNTIME

ARG VERSION=knative-v1.15

COPY --from=builder /usr/bin/main /usr/bin/eventmesh
COPY LICENSE /licenses/

USER 65532

LABEL \
      com.redhat.component="openshift-serverless-1-kn-backstage-plugins-eventmesh-rhel8-container" \
      name="openshift-serverless-1/kn-backstage-plugins-eventmesh-rhel8" \
      version=$VERSION \
      summary="Red Hat OpenShift Serverless 1 Backstage Plugins Eventmesh" \
      maintainer="serverless-support@redhat.com" \
      description="Red Hat OpenShift Serverless 1 Backstage Plugins Eventmesh" \
      io.k8s.display-name="Red Hat OpenShift Serverless 1 Backstage Plugins Eventmesh" \
      io.k8s.description="Red Hat OpenShift Serverless Backstage Plugins Eventmesh" \
      io.openshift.tags="eventmesh" \
      vendor="Red Hat, Inc." \
      release=$VERSION \
      cpe="cpe:/a:redhat:openshift_serverless:1.35::el8"
      
ENTRYPOINT ["/usr/bin/eventmesh"]
