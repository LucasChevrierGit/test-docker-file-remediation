FROM alpine:3.17

LABEL org.opencontainers.image.title="dockertest-caseb" \
      org.opencontainers.image.version="1.0.0" \
      org.opencontainers.image.description="Test image for Case B vulnerability remediation" \
      org.opencontainers.image.base.name="alpine:3.17" \
      org.opencontainers.image.source="https://github.com/LucasChevrierGit/test-docker-file-remediation" \
      com.datadoghq.image.source_path="caseB/Dockerfile"

RUN apk add --no-cache curl
COPY . /app
WORKDIR /app
CMD ["./start.sh"]
