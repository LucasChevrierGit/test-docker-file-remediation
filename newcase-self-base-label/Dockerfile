# Reproduces a base image that names itself as its own base, like a GBI.
#
# A GBI sets LABEL baseimage.name to its own name so that images built FROM it
# inherit the label. On the GBI itself, the label names the image, not a base.
# Here the image is tagged registry.example.com/images/base/selfbase-label-test
# and labels itself images/base/selfbase-label-test (no registry), the same
# shape as images/base/gbi-ubuntu_2404 on registry.ddbuild.io.
#
# The final stage is FROM scratch + COPY --from=build / / (one flattened layer,
# as in the GBI), so ISR's layer matcher finds no base and falls back to the
# label. The build stage is the vulnerable node base of
# newcase-remediated_copy_no_commit, so the image has critical findings.

FROM node@sha256:f9ab18e354e6855ae56ef2b290dd225c1e51a564f87584b9bd21dd651838830e AS build

FROM scratch
COPY --from=build / /

LABEL org.opencontainers.image.title="self-base-label-test" \
      org.opencontainers.image.version="1.0.0" \
      org.opencontainers.image.description="Base image that labels itself as its own base image" \
      org.opencontainers.image.source="https://github.com/LucasChevrierGit/test-docker-file-remediation" \
      com.datadoghq.image.source_path="newcase-self-base-label/Dockerfile" \
      container.baseimage.name="images/base/selfbase-label-test"

CMD ["/bin/bash"]
