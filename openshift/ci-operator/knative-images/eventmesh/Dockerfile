# DO NOT EDIT! Generated Dockerfile for backends/cmd/eventmesh.
ARG GO_BUILDER=registry.ci.openshift.org/openshift/release:rhel-8-release-golang-1.21-openshift-4.16
ARG GO_RUNTIME=registry.access.redhat.com/ubi8/ubi-minimal

FROM $GO_BUILDER as builder

COPY . .

ENV CGO_ENABLED=1
ENV GOEXPERIMENT=strictfipsruntime

RUN go build -tags strictfipsruntime -o /usr/bin/main ./backends/cmd/eventmesh

FROM $GO_RUNTIME

ARG VERSION=knative-nightly

COPY --from=builder /usr/bin/main /usr/bin/eventmesh

USER 65532

LABEL \
      com.redhat.component="openshift-serverless-1-backstage-plugins-eventmesh-rhel8-container" \
      name="openshift-serverless-1/backstage-plugins-eventmesh-rhel8" \
      version=$VERSION \
      summary="Red Hat OpenShift Serverless 1 Backstage Plugins Eventmesh" \
      maintainer="serverless-support@redhat.com" \
      description="Red Hat OpenShift Serverless 1 Backstage Plugins Eventmesh" \
      io.k8s.display-name="Red Hat OpenShift Serverless 1 Backstage Plugins Eventmesh"

ENTRYPOINT ["/usr/bin/eventmesh"]
