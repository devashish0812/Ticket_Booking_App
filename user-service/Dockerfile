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

# Build static binary (no glibc dependency)
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o server .

# ----------------------
# Stage 2: Final Image
# ----------------------
FROM alpine:3.20

# Install certificates (needed for HTTPS / MongoDB Atlas, etc.)
RUN apk add --no-cache ca-certificates

WORKDIR /app

# Copy binary from builder
COPY --from=builder /app/server .

EXPOSE 8080

CMD ["./server"]
