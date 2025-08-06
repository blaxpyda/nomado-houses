# Multi-stage Dockerfile for Nomado Houses Application
# Stage 1: Build the Go backend
FROM golang:1.21-alpine AS backend-builder

# Install build dependencies
RUN apk add --no-cache git

# Set working directory for backend
WORKDIR /app/backend

# Copy go mod files
COPY backend/go.mod backend/go.sum ./

# Download dependencies
RUN go mod download

# Copy backend source code
COPY backend/ ./

# Build the Go application
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main .

# Stage 2: Create the final runtime image
FROM alpine:latest

# Install ca-certificates for HTTPS requests and timezone data
RUN apk --no-cache add ca-certificates tzdata

# Create app directory
WORKDIR /app

# Create non-root user for security
RUN addgroup -g 1001 -S appgroup && \
    adduser -S appuser -u 1001 -G appgroup

# Copy the binary from builder stage
COPY --from=backend-builder /app/backend/main ./

# Copy static files (frontend)
COPY public/ ./public/

# Create logs directory
RUN mkdir -p logs && chown -R appuser:appgroup logs

# Copy environment example (you should mount your actual .env file)
COPY backend/.env.example ./.env.example

# Set proper permissions
RUN chown -R appuser:appgroup /app

# Switch to non-root user
USER appuser

# Expose port (default 8080, can be overridden with environment variable)
EXPOSE 8080

# Run the application
CMD ["./main"]
