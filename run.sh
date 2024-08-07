#!/bin/bash

case "$1" in
    exercises)
        echo "Running exercises..."
        go run -race ./exercises/main.go
        ;;
    concurrency-exercises)
        echo "Running concurrency exercises..."
        go run -race ./concurrency_exercises/main.go
        ;;
    cli-app)
        echo "Running command line app..."
        go run ./cli_app/main.go
        ;;
    webapp-backend)
        echo "Running WebApp Backend..."
        cd ./webapp/backend
        go run ./main.go
        cd -  # Return to the original directory
        ;;
    webapp-frontend)
        echo "Running WebApp Frontend..."
        cd ./webapp/frontend
        go run ./main.go
        cd -  # Return to the original directory
        ;;
    webapp)
        echo "Running WebApp Backend and Frontend..."
        # Run backend in background
        (cd ./webapp/backend && go run ./main.go) &
        # Run frontend in background
        (cd ./webapp/frontend && go run ./main.go) &
        # Wait for all background processes to complete
        wait
        ;;
    *)
        echo "Usage: $0 {exercises|concurrency-exercises|cli-app|webapp|webapp-backend|webapp-frontend}"
        exit 1
        ;;
esac