# 1. Build stage: compile the Go app
FROM golang:1.24 AS builder

WORKDIR /app

# Copy go.mod first
COPY go.mod ./

# Download modules
RUN go mod download || true

# Copy the rest of the source code
COPY . .

# Build the Go binary
RUN go build -o yenup

# 2. Run stage: minimal image
FROM debian:bullseye-slim

WORKDIR /app

# Copy the compiled binary from the builder
COPY --from=builder /app/yenup .

# Copy rate.json
COPY rate.json .

# Command to run the app
CMD ["./yenup"]



