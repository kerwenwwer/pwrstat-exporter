# Define variables for repeated values
DOCKER_TAG := pwrstat-exporter:latest

# Phony targets for workflows
.PHONY: all build docker-build docker-push

# Default target to compile the application and build the Docker image
all: build docker-build

# Rule to build the main application
build: 
	go build -o ./bin/pwrstat-exporter ./main.go

# Rule to build the Docker image
docker-build:
	docker build -t $(DOCKER_TAG) .
