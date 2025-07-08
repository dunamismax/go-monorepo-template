# Stage 1: Build the application
FROM golang:1.24-alpine AS builder

# Set the working directory
WORKDIR /app

# Copy go.mod and go.sum files to download dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy the rest of the source code
COPY . .

# Build the application
# -o /app/server: output file name
# -ldflags="-s -w": strip debugging information
# CGO_ENABLED=0: disable CGO for static linking
# GOOS=linux GOARCH=amd64: build for linux
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="-s -w" -o /app/server ./cmd/demo-http-server

# Stage 2: Create the final image
FROM alpine:latest

# Set the working directory
WORKDIR /app

# Copy the built binary from the builder stage
COPY --from=builder /app/server /app/server

# Expose the port the application runs on
EXPOSE 3000

# Set the entrypoint
ENTRYPOINT ["/app/server"]
