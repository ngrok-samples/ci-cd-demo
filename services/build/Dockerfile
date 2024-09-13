# Start from the official Go image
FROM golang:1.20-alpine

# Set the working directory inside the container
WORKDIR /app

# Copy the source code
COPY main.go .

# Initialize Go module
RUN go mod init build-service

# Download the dependencies and create go.sum
RUN go mod tidy

# Build the application
RUN go build -o build-service

# Expose the port the app runs on
EXPOSE 8080

# Run the application
CMD ["./build-service"]
