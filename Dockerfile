# =====================================
# GeoVision3D Dockerfile
# Multi-stage Production Build
# =====================================

# ---------- BUILD STAGE ----------
FROM golang:1.22-alpine AS builder

# Install git
RUN apk add --no-cache git

# Working directory
WORKDIR /app

# Copy go modules
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy source code
COPY . .

# Build binary
RUN CGO_ENABLED=0 GOOS=linux go build -o geovision3d .

# ---------- FINAL STAGE ----------
FROM alpine:latest

# Install certificates
RUN apk --no-cache add ca-certificates

# Working directory
WORKDIR /root/

# Copy binary
COPY --from=builder /app/geovision3d .

# Copy static assets
COPY --from=builder /app/static ./static

# Copy environment example
COPY --from=builder /app/.env ./.env

# Expose port
EXPOSE 8080

# Environment
ENV PORT=8080

# Start application
CMD ["./geovision3d"]
