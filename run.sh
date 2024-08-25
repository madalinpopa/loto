#!/bin/bash

DOCKER_RUN_ARGS="-d --rm -p 8090:8090 -v pocketbase:/app/pb_data --name loto coderustle/loto:latest"

# Check if docker volume exists
if docker volume ls | grep -q pocketbase; then
  echo "Volume exists"
else
  echo "Creating volume"
  docker volume create --name=pocketbase
fi

# Check if container exists
if docker ps -a | grep -q loto; then
  echo "Container exists"
  echo "Stopping container"
  docker stop loto
  echo "Building image"
  docker build -t coderustle/loto:latest .
  docker run $DOCKER_RUN_ARGS
else
  echo "Building image"
  docker build -t coderustle/loto:latest .
  echo "Creating container"
  docker run $DOCKER_RUN_ARGS
fi

