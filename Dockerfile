# Use a Golang base image with a specific version tag for stability
FROM golang:1.22.1-alpine AS builder

# Set the working directory inside the container
WORKDIR /app

# Copy the source code into the container
COPY src/ .

# Build the Go app
RUN go build -o /app/main .

# Use a lightweight base image for the final stage
FROM alpine:latest

# Set the working directory inside the container
WORKDIR /app

# Copy the built binary from the previous stage
COPY --from=builder /app/main .

# Expose port 8080 to the outside world
EXPOSE 8080

# Use exec form for ENTRYPOINT to avoid shell processing
ENTRYPOINT ["./main"]
