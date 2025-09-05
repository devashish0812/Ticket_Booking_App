# ----------------------
# Stage 1: Builder
# ----------------------
FROM golang:1.24-bookworm AS builder

WORKDIR /app

# Copy dependency files first
COPY go.mod go.sum ./
RUN go mod download

# Copy source
COPY . .

# Build static binary
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o api-gateway .

# ----------------------
# Stage 2: Final Image
# ----------------------
FROM alpine:3.20

# Install certificates (needed for HTTPS calls to other services)
RUN apk add --no-cache ca-certificates

WORKDIR /app

# Copy binary from builder
COPY --from=builder /app/api-gateway .

# Expose the gateway port
EXPOSE 8081

# Run the binary
CMD ["./api-gateway"]
