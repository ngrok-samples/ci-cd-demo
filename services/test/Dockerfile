# Start from the official Go image
FROM golang:1.20-alpine

# Set the working directory inside the container
WORKDIR /app

# Copy the source code
COPY main.go .

# Initialize Go module
RUN go mod init test-service

# Download the dependencies and create go.sum
RUN go mod tidy

# Build the application
RUN go build -o test-service

# Expose the port the app runs on
EXPOSE 8081

# Run the application
CMD ["./test-service"]
