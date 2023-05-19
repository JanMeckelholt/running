#!/bin/bash

DOCKER_BUILDKIT=1 COMPOSE_DOCKER_CLI_BUILD=1 PLATFORM=linux/arm64 docker-compose build

docker tag running_http_gateway janmeckelholt/running_http_gateway:$(git rev-parse main)_arm64  && \
docker tag running_runner janmeckelholt/running_runner:$(git rev-parse main)_arm64 && \
docker tag running_strava-service janmeckelholt/running_strava-service:$(git rev-parse main)_arm64 && \
docker tag running_database-service janmeckelholt/running_database-service:$(git rev-parse main)_arm64 && \
docker tag running_running_app janmeckelholt/running_running_app:$(git rev-parse main)_arm64 && \
docker tag running_postgres janmeckelholt/running_postgres:$(git rev-parse main)_arm64 && \
echo $1 | docker login -u janmeckelholt --password-stdin
docker push janmeckelholt/running_http_gateway:$(git rev-parse main)_arm64
docker push janmeckelholt/running_runner:$(git rev-parse main)_arm64
docker push janmeckelholt/running_strava-service:$(git rev-parse main)_arm64
docker push janmeckelholt/running_database-service:$(git rev-parse main)_arm64
docker push janmeckelholt/running_running_app:$(git rev-parse main)_arm64
docker push janmeckelholt/running_postgres:$(git rev-parse main)_arm64
