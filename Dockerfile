# MCP Ultra Multi-stage Dockerfile
# Build stage
FROM golang:1.24-rc-alpine AS builder

# Install build dependencies
RUN apk update && \
    apk add --no-cache git make ca-certificates tzdata && \
    rm -rf /var/cache/apk/*

WORKDIR /build

# Copy dependency files first for better caching
COPY go.mod go.sum ./
RUN go mod download && go mod verify

# Copy source code
COPY . .

# Build the application with optimizations
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build \
    -trimpath \
    -ldflags="-w -s -X 'github.com/vertikon/mcp-ultra/pkg/version.Version=v1.0.0' -X 'github.com/vertikon/mcp-ultra/pkg/version.GitCommit=$(git rev-parse HEAD 2>/dev/null || echo 'unknown')' -X 'github.com/vertikon/mcp-ultra/pkg/version.BuildDate=$(date -u +%Y-%m-%dT%H:%M:%SZ)'" \
    -o mcp-ultra cmd/mcp-model-ultra/main.go

# Runtime stage with minimal Alpine
FROM alpine:latest

# Install only essential packages
RUN apk update && \
    apk add --no-cache ca-certificates tzdata wget && \
    rm -rf /var/cache/apk/* && \
    adduser -D -g '' -u 1000 appuser

WORKDIR /app

# Copy binary and configuration with proper ownership
COPY --from=builder --chown=appuser:appuser /build/mcp-ultra ./
COPY --from=builder --chown=appuser:appuser /build/config ./config
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/

# Switch to non-root user
USER appuser

# Expose ports
EXPOSE 9655 9656

# Health check
HEALTHCHECK --interval=30s --timeout=3s --start-period=10s --retries=3 \
    CMD wget --quiet --tries=1 --spider http://localhost:9655/healthz || exit 1

# Set entrypoint
ENTRYPOINT ["./mcp-ultra"]