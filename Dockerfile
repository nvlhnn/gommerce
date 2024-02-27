# Stage 1: Build the Go application
FROM golang:1.21.1-alpine AS builder

WORKDIR /app

# Copy only the necessary Go files and dependencies
COPY . .
RUN go mod download

# Copy the entire cmd/ directory into the /app directory

# Build the Go application
RUN CGO_ENABLED=0 GOOS=linux go build -o gommerce ./cmd/main.go

# Stage 2: Create a minimal image
FROM alpine:latest

WORKDIR /app

# Copy the binary from the builder stage
COPY --from=builder /app/gommerce .
COPY --from=builder /app/.env .

# Expose the port your Go application listens on (change to match your app)
EXPOSE 3000

# Command to run the executable
CMD ["./gommerce"]
