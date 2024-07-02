# Dockerfile

# Use the official Golang image to create a build artifact.
FROM golang:1.20 AS builder

# Set the Current Working Directory inside the container
WORKDIR /app

# Copy go mod and sum files
COPY go.mod go.sum ./

# Download all dependencies. Dependencies will be cached if the go.mod and go.sum files are not changed
RUN go mod download

# Copy the source from the current directory to the Working Directory inside the container
COPY . .

# Build the Go app
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main .

# Start a new stage from scratch
FROM alpine:latest

WORKDIR /root/

# Copy the Pre-built binary file from the previous stage
COPY --from=builder /app/main .
COPY --from=builder /app/messages /root/messages

# Create the messages directory if it doesn't exist
RUN mkdir -p /root/messages

# Ensure the main file is executable
RUN chmod +x ./main

EXPOSE 8080

# Command to run the executable
CMD ["./main"]
