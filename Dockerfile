# Stage 1: Build the application
FROM golang:1.17 AS builder

# Set the working directory inside the container
WORKDIR /app

# Copy the Go application files
COPY . .

# Build the Go application
RUN go build -o main .

# Stage 2: Create a minimal image to run the application
FROM alpine:latest

# Set the working directory inside the container
WORKDIR /app

# Copy only the necessary files from the builder stage
COPY --from=builder /app/main .

# Expose the port the application runs on
EXPOSE 8080

# Command to run the application
CMD ["./main"]
