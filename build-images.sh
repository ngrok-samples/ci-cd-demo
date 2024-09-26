#!/bin/bash

# Create a new builder instance
docker buildx create --name mybuilder --use

# Build and push the Build Service
docker buildx build --platform linux/amd64,linux/arm64,linux/arm64/v8 \
  -t your-registry/build-service:latest \
  -f services/build/Dockerfile \
  --push \
  services/build

# Build and push the Test Service
docker buildx build --platform linux/amd64,linux/arm64,linux/arm64/v8 \
  -t your-registry/test-service:latest \
  -f services/test/Dockerfile \
  --push \
  services/test

# Build and push the Deployment Service
docker buildx build --platform linux/amd64,linux/arm64,linux/arm64/v8 \
  -t your-registry/deployment-service:latest \
  -f services/deployment/Dockerfile \
  --push \
  services/deployment

# Remove the builder instance
docker buildx rm mybuilder
