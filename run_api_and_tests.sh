#!/bin/bash

# Build the Go application
cd app
go build -o sapi

# Start the API in the background
./sapi &

# Wait for the API to start by checking the health endpoint
response=0
until [ $response -eq 200 ]; do
    response=$(curl --write-out '%{http_code}' --silent --output /dev/null http://localhost:8080/health)
    printf '.'
    sleep 1
done

echo "API is up and running"

cd ..
# Run the integration tests
# print the output of the tests
pytest integration -v

# Kill the API process after the tests are completed
kill $(pgrep sapi)
