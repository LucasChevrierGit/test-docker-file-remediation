# Image built on the self-labelled base image in newcase-self-base-label.
#
# It inherits container.baseimage.name="images/base/selfbase-label-test" from
# its base, like a service image built FROM a GBI. Its own name is different
# (registry.example.com/team/selfbase-child), so ISR must keep
# images/base/selfbase-label-test as its base image, while the base itself
# reports none. Its critical findings come from the base's packages.

FROM registry.example.com/images/base/selfbase-label-test:v2

LABEL org.opencontainers.image.title="self-base-child-test" \
      org.opencontainers.image.version="1.0.0" \
      org.opencontainers.image.description="Image built on a base image that labels itself" \
      org.opencontainers.image.source="https://github.com/LucasChevrierGit/test-docker-file-remediation" \
      com.datadoghq.image.source_path="newcase-self-base-child/Dockerfile"
