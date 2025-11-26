# Multi-stage build for Cloudflare Manager
FROM node:18-alpine AS frontend-builder

WORKDIR /app/frontend
COPY frontend/package*.json ./
RUN npm install
COPY frontend/ ./
RUN npm run build

# Go builder
FROM golang:1.21-alpine AS go-builder

WORKDIR /app
COPY cloudflare-manager/ ./cloudflare-manager/
WORKDIR /app/cloudflare-manager
RUN go mod download
RUN CGO_ENABLED=0 GOOS=linux go build -o cfm

# Final stage
FROM python:3.11-slim

WORKDIR /app

# Install system dependencies
RUN apt-get update && apt-get install -y \
    curl \
    && rm -rf /var/lib/apt/lists/*

# Copy Go binary
COPY --from=go-builder /app/cloudflare-manager/cfm /app/cloudflare-manager/cfm

# Copy and install Python backend
COPY backend/ ./backend/
WORKDIR /app/backend
RUN pip install --no-cache-dir -r requirements.txt

# Copy frontend build
COPY --from=frontend-builder /app/frontend/build /app/frontend/build

# Create config directory
RUN mkdir -p /root/.cloudflare-manager

# Initialize with provided credentials
ENV CLOUDFLARE_EMAIL="tqa88tawlq@downnaturer.me"
ENV CLOUDFLARE_TOKEN="4e2dd4818267ebd2ab8d1aa2e7f9bf4151b70"

# Expose port
EXPOSE 7860

# Health check
HEALTHCHECK --interval=30s --timeout=10s --start-period=5s --retries=3 \
    CMD curl -f http://localhost:7860/health || exit 1

# Start script
COPY docker-entrypoint.sh /app/
RUN chmod +x /app/docker-entrypoint.sh

CMD ["/app/docker-entrypoint.sh"]
