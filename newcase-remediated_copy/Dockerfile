FROM moby/buildkit:v0.28.1 AS buildkit
FROM alpine:3.17

LABEL org.opencontainers.image.title="new Case copy" \
      org.opencontainers.image.version="1.0.1" \
      org.opencontainers.image.base.name="alpine:3.17" \
      org.opencontainers.image.source="https://github.com/LucasChevrierGit/test-docker-file-remediation" \
      com.datadoghq.image.source_path="newcase-remediated_copy/Dockerfile" \
      com.datadoghq.remediation.package="github.com/moby/buildkit"

RUN apk add --no-cache curl
COPY --from=buildkit /usr/bin/buildctl /usr/local/bin/buildctl
COPY . /app
WORKDIR /app
CMD ["./start.sh"]
