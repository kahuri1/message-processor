# Start from the official Go image
FROM golang:1.22-alpine AS builder

# Set the Current Working Directory inside the container
WORKDIR /app

# Copy go mod and sum files
COPY go.mod go.sum ./

# Download all dependencies
RUN go mod download

# Copy the source code and migration files from the current directory to the Working Directory inside the container
COPY . .

# Build the Go app
RUN go build -o /app/message-processor ./cmd/message-processor

# Start a new stage from scratch
FROM alpine:latest

# Copy the Pre-built binary file from the previous stage
COPY --from=builder /app/message-processor /app/message-processor

# Copy the configuration file
COPY configs/config.yml /configs/config.yml

# Copy the migrations to the app directory
COPY migrations /app/migrations

# Expose port 8080 to the outside world
EXPOSE 8080

# Command to run the executable
ENTRYPOINT ["/app/message-processor"]
