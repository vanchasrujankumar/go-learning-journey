# Multi-stage build for Go Learning Journey projects
# Stage 1: Builder
FROM golang:1.26-alpine AS builder

# Install build dependencies
RUN apk add --no-cache git gcc musl-dev

WORKDIR /build

# Copy go mod files
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy source code
COPY . .

# Build the application
ARG SERVICE_NAME=main
RUN CGO_ENABLED=1 GOOS=linux go build -a -installsuffix cgo -o /app/${SERVICE_NAME} ./cmd/${SERVICE_NAME}

# Stage 2: Runtime
FROM alpine:latest

# Install runtime dependencies
RUN apk --no-cache add ca-certificates

WORKDIR /root/

# Copy binary from builder
COPY --from=builder /app/* ./

# Health check
HEALTHCHECK --interval=30s --timeout=10s --start-period=5s --retries=3 \
  CMD wget --quiet --tries=1 --spider http://localhost:8080/health || exit 1

# Run application
CMD ["./main"]

# Expose port (adjust as needed)
EXPOSE 8080
