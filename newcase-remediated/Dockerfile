FROM moby/buildkit:v0.28.1 AS buildkit
FROM alpine:3.17

LABEL org.opencontainers.image.title="new Case" \
      org.opencontainers.image.version="1.0.0" \
      org.opencontainers.image.description="Test image with critical vuln" \
      org.opencontainers.image.base.name="alpine:3.17" \
      org.opencontainers.image.source="https://github.com/LucasChevrierGit/test-docker-file-remediation" \
      com.datadoghq.image.source_path="newcase-remediated/Dockerfile" \
      com.datadoghq.remediation.cve="CVE-2024-23652" \
      com.datadoghq.remediation.package="github.com/moby/buildkit" \
      com.datadoghq.remediation.from="v0.11.4" \
      com.datadoghq.remediation.to="v0.29.1-0.20260402225136-c90e21f2f1e6"

RUN apk add --no-cache curl
COPY --from=buildkit /usr/bin/buildctl /usr/local/bin/buildctl
COPY . /app
WORKDIR /app
CMD ["./start.sh"]
