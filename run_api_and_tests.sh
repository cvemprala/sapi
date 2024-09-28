#!/bin/bash

# Build the Go application
go build -o sapi app/main.go

# Start the API in the background
./sapi &

# Wait for the API to start by checking the health endpoint
until $(curl --output /dev/null --silent --head --fail http://localhost:8080/api/todos); do
    printf '.'
    sleep 1
done

# Run the integration tests
python -m unittest discover -s integration

# Kill the API process after the tests are completed
kill %1
