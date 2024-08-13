#!/bin/bash

# Define variables
IMAGE_NAME="ascii-art-web"
CONTAINER_NAME="ascii-art-web-container"
PORT=8080

# Build the Docker image
echo "Building Docker image..."
docker build -t $IMAGE_NAME .

# Stop and remove any existing container with the same name
echo "Stopping and removing existing container..."
docker container stop $CONTAINER_NAME 2>/dev/null
docker container rm $CONTAINER_NAME 2>/dev/null

# Run the Docker container
echo "Running Docker container..."
docker run -d -p $PORT:$PORT --name $CONTAINER_NAME $IMAGE_NAME

echo "Docker image built and container started."