# Use the official Golang image to build the application
FROM golang:1.23 as builder

# Set the Current Working Directory inside the container
WORKDIR /app

# Copy go mod and sum files
COPY go.mod go.sum ./

# Download all dependencies
RUN go mod download

# Copy the source code into the container
COPY . .

# Build the Go application
RUN bash build.sh

# Use a minimal image to run the application
FROM alpine:latest

# Install necessary certificates
RUN apk --no-cache add ca-certificates

# Set the Current Working Directory inside the container
WORKDIR /root/

# Copy the pre-built binary from the builder stage
COPY --from=builder /app/api .

# Expose the application port
EXPOSE 8080

# Command to run the executable
CMD ["./api"]