# DO NOT EDIT! Generated Dockerfile for vendor/knative.dev/eventing/test/test_images/event-sender.
ARG GO_BUILDER=registry.ci.openshift.org/openshift/release:rhel-8-release-golang-1.21-openshift-4.16
ARG GO_RUNTIME=registry.access.redhat.com/ubi8/ubi-minimal

FROM $GO_BUILDER as builder

COPY . .

ENV CGO_ENABLED=1
ENV GOEXPERIMENT=strictfipsruntime

RUN go build -tags strictfipsruntime -o /usr/bin/main ./vendor/knative.dev/eventing/test/test_images/event-sender

FROM $GO_RUNTIME

ARG VERSION=knative-nightly

COPY --from=builder /usr/bin/main /usr/bin/event-sender

USER 65532

LABEL \
      com.redhat.component="openshift-serverless-1-backstage-plugins-vendor-knative.dev-eventing-test-test-images-event-sender-rhel8-container" \
      name="openshift-serverless-1/backstage-plugins-vendor-knative.dev-eventing-test-test-images-event-sender-rhel8" \
      version=$VERSION \
      summary="Red Hat OpenShift Serverless 1 Backstage Plugins Vendor Knative.Dev Eventing Test Test Images Event Sender" \
      maintainer="serverless-support@redhat.com" \
      description="Red Hat OpenShift Serverless 1 Backstage Plugins Vendor Knative.Dev Eventing Test Test Images Event Sender" \
      io.k8s.display-name="Red Hat OpenShift Serverless 1 Backstage Plugins Vendor Knative.Dev Eventing Test Test Images Event Sender" \
      io.k8s.description="Red Hat OpenShift Serverless Backstage Plugins Vendor Knative.Dev Eventing Test Test Images Event Sender" \
      io.openshift.tags="vendor-knative.dev-eventing-test-test-images-event-sender"

ENTRYPOINT ["/usr/bin/event-sender"]
