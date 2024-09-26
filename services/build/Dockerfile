# Start from the official Go image
FROM --platform=$TARGETPLATFORM golang:1.20-alpine AS builder

# Set the working directory inside the container
WORKDIR /app

# Copy the source code
COPY main.go .

# Initialize Go module
RUN go mod init build-service

# Download the dependencies and create go.sum
RUN go mod tidy

# Build the application
ARG TARGETPLATFORM
RUN case ${TARGETPLATFORM} in \
  "linux/amd64")  GOARCH=amd64  ;; \
  "linux/arm64")  GOARCH=arm64  ;; \
  "linux/arm64/v8")  GOARCH=arm64  ;; \
esac && \
  GOOS=linux go build -o build-service

# Start a new stage for a smaller final image
FROM --platform=$TARGETPLATFORM alpine:latest

# Copy the binary from the builder stage
COPY --from=builder /app/build-service /app/build-service

# Expose the port the app runs on
EXPOSE 8080

# Run the application
CMD ["./build-service"]
