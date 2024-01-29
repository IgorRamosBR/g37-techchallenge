
FROM golang:1.20-alpine as builder

WORKDIR /app

COPY . ./

# Build the binary.
RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-s -w" -o bin/main ./cmd/g37-lanches/main.go

# Use the official Debian slim image for a lean production container.
# https://hub.docker.com/_/debian
# https://docs.docker.com/develop/develop-images/multistage-build/#use-multi-stage-builds
FROM debian:buster-slim
RUN set -x && apt-get update && DEBIAN_FRONTEND=noninteractive apt-get install -y \
    ca-certificates && \
    rm -rf /var/lib/apt/lists/*

# Copy the binary to the production image from the builder stage.
COPY --from=builder /app/bin/main /app/main
COPY --from=builder /app/configs /configs
COPY --from=builder /app/migrations /migrations

EXPOSE 8080

# Run the web service on container startup.
CMD ["/app/main"]