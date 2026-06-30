FROM ubuntu:18.04

LABEL org.opencontainers.image.title="apt-vuln-test" \
      org.opencontainers.image.version="1.0.0" \
      org.opencontainers.image.description="Test image for apt and Go remediation template validation" \
      org.opencontainers.image.source="https://github.com/LucasChevrierGit/test-docker-file-remediation" \
      com.datadoghq.image.source_path="newcase-remediated_copy_no_commit/Dockerfile"

ENV DEBIAN_FRONTEND=noninteractive

RUN apt-get update \
    && apt-get install -y --no-install-recommends \
        git \
        curl \
        openssl \
        wget \
        python3 \
        python3-pip \
        ca-certificates \
    && apt-get clean \
    && rm -rf /var/lib/apt/lists/*

# Install Go
RUN curl -sSL https://go.dev/dl/go1.21.0.linux-amd64.tar.gz | tar -C /usr/local -xz
ENV PATH="/usr/local/go/bin:${PATH}"

# Install a Go binary at a known-vulnerable version
RUN go install github.com/moby/buildkit/cmd/buildctl@v0.11.4

WORKDIR /app

# Copy in a Go module with vulnerable dependencies
COPY go.mod go.sum ./
RUN go mod download

CMD ["/bin/bash"]
