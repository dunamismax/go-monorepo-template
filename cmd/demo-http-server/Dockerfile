# Stage 1: Build the application binary
# Use a specific Go version for reproducibility.
FROM golang:1.24 as build

WORKDIR /build/src

# Copy dependency files and download them.
# This creates a separate cache layer for dependencies.
COPY go.mod go.sum ./
RUN go mod download

# Copy the rest of the source code.
COPY . .

# Build a statically linked and stripped binary.
# -trimpath removes file system paths from the executable.
# -ldflags="-s -w" strips debug information.
# The output is a single binary named 'app'.
RUN CGO_ENABLED=0 go build -trimpath -ldflags="-s -w" -o app ./cmd/demo-http-server/cmd

# Stage 2: Create the final, minimal image
# Start from a completely empty image.
FROM scratch

# Copy the compiled application binary from the build stage.
COPY --from=build /build/src/app /usr/bin/app

# Expose the port the server listens on.
# This is metadata and good practice for consumers of the image.
EXPOSE 3000

# Set the entrypoint for the container.
# This is the command that will be executed when the container starts.
ENTRYPOINT ["/usr/bin/app"]