# Use a pinned, official Go image for the build stage.
# Using --platform=$BUILDPLATFORM ensures the base image matches the build architecture.
# Pinning with a SHA256 hash guarantees reproducibility and security.
FROM --platform=$BUILDPLATFORM golang:1.24.4-alpine3.22@sha256:68932fa6d4d4059845c8f40ad7e654e626f3ebd3706eef7846f319293ab5cb7a AS build

WORKDIR /usr/local/src/app

COPY go.mod go.sum ./

# Use a cache mount for Go modules to speed up subsequent builds.
RUN --mount=type=cache,target=/root/go/pkg/mod go mod download

# These ARGs are automatically populated by Docker Buildx from the --platform flag.
ARG TARGETOS
ARG TARGETARCH

COPY ./ ./

# Build the application, mounting both module and build caches.
# This creates a statically linked binary for the target OS/architecture.
RUN --mount=type=cache,target=/root/go/pkg/mod --mount=type=cache,target=/root/.cache/go-build \
    CGO_ENABLED=0 GOOS=${TARGETOS} GOARCH=${TARGETARCH} \
    go build -trimpath -ldflags="-s -w" -o /usr/local/bin/app ./cmd/demo-http-server/cmd

# --- Final Stage ---
# Start from a completely empty image for maximum security and minimal size.
FROM scratch

# Copy the compiled application from the build stage.
COPY --from=build /usr/local/bin/app /usr/local/bin/app

# Expose the server port, specifying the protocol.
EXPOSE 3000/tcp

# Run as a non-root user ('nobody') for enhanced security.
# UID 65534 is standard for the 'nobody' user.
USER 65534:65534

# Set the entrypoint for the container.
ENTRYPOINT ["/usr/local/bin/app"]
