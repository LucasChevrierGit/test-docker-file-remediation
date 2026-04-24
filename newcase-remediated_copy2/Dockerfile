FROM alpine:3.17

LABEL org.opencontainers.image.title="new Case copy 2" \
      org.opencontainers.image.version="1.0.2" \
      org.opencontainers.image.base.name="alpine:3.17" \
      org.opencontainers.image.source="https://github.com/LucasChevrierGit/test-docker-file-remediation" \
      com.datadoghq.image.source_path="newcase-remediated_copy2/Dockerfile" \
      com.datadoghq.remediation.package="github.com/moby/buildkit"

RUN apk add --no-cache curl go
RUN go install github.com/moby/buildkit/cmd/buildctl@v0.11.4
COPY . /app
WORKDIR /app
CMD ["./start.sh"]
