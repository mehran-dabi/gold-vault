# Use the official Golang image as the builder
FROM golang:1.23 AS builder

# Set the working directory inside the container
WORKDIR /app

# Copy the Go module files and download dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy the rest of the application source code
COPY . .

# Change directory to cmd/app to build the application
WORKDIR /app/cmd/app

# Build the Go application with static linking for Alpine
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o /app/user-service .

# Use a lightweight image to run the application
FROM alpine:latest

# Install glibc compatibility for Alpine and CA certificates
RUN apk --no-cache add ca-certificates

# Set the working directory
WORKDIR /app/

# Copy the built binary from the builder stage
COPY --from=builder /app /app

# Expose necessary ports (example: HTTP and gRPC)
EXPOSE 9000 9001 9002

# Run the application
CMD ["./user-service"]
