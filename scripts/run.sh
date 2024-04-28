#!/bin/bash

# Assign the first command line argument to 'command', defaulting to "start" if not provided
command=${1:-"start"}

# Set the ProjectRoot relative to the script location
ProjectRoot="$(dirname "$BASH_SOURCE")/.."

# Export environment variables
export LAB_API_ENVIRONMENT="Development"
export LAB_API_PORT="8080"
export LAB_API_MONGODB_USERNAME="root"
export LAB_API_MONGODB_PASSWORD="neUhaDnes"

# Define the mongo function
mongo() {
    docker compose --file "${ProjectRoot}/deployments/docker-compose/compose.yaml" "$@"
}

# The command is expected to be passed as the first script argument
command=$1

# Switch-case-like behavior in bash
case "$command" in
    openapi)
        docker run --rm -ti -v "${ProjectRoot}:/local" openapitools/openapi-generator-cli generate -c /local/scripts/generator-cfg.yaml
        ;;
    start)
        mongo up --detach
        go run "${ProjectRoot}/cmd/lab-api-service"
        mongo down
        ;;
    up)
        mongo up
        ;;
    test)
        go test -v ./...
        ;;
    docker)
        docker build -t magum/lab-webapi:local-build -f ${ProjectRoot}/build/docker/Dockerfile .
        ;;
    *)
        echo "Unknown command: $command"
        exit 1
        ;;
esac