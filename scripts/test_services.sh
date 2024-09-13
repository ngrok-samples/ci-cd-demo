#!/bin/bash

# Set the base URL to your ngrok URL
BASE_URL="http://your-ngrok-url"

# Test Build Service
echo "Testing Build Service..."
BUILD_RESPONSE=$(curl -s -X POST "${BASE_URL}/builds/trigger")
BUILD_ID=$(echo $BUILD_RESPONSE | jq -r .id)
echo "Triggered build: $BUILD_ID"
sleep 5
curl -s "${BASE_URL}/builds/${BUILD_ID}" | jq .

# Test Test Service
echo "Testing Test Service..."
TEST_RESPONSE=$(curl -s -X POST "${BASE_URL}/tests/run" -H "Content-Type: application/json" -d "{\"build_id\": \"${BUILD_ID}\"}")
TEST_ID=$(echo $TEST_RESPONSE | jq -r .id)
echo "Triggered test run: $TEST_ID"
sleep 5
curl -s "${BASE_URL}/tests/${TEST_ID}" | jq .

# Test Deployment Service
echo "Testing Deployment Service..."
DEPLOY_RESPONSE=$(curl -s -X POST "${BASE_URL}/deployments/create" -H "Content-Type: application/json" -d "{\"build_id\": \"${BUILD_ID}\", \"environment\": \"staging\"}")
DEPLOY_ID=$(echo $DEPLOY_RESPONSE | jq -r .id)
echo "Triggered deployment: $DEPLOY_ID"
sleep 5
curl -s "${BASE_URL}/deployments/${DEPLOY_ID}" | jq .

echo "Test complete!"
