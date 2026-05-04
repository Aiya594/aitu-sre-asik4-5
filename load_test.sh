#!/bin/bash

echo "Starting load simulation..."

make_request() {
    local service=$1
    local endpoint=$2
    local count=$3

    for i in $(seq 1 $count); do
        curl -s -o /dev/null -w "%{http_code}\n" http://localhost:$service$endpoint &
    done
}

echo "Simulating authentication requests..."
make_request 8080 "/health" 50

echo "Simulating product requests..."
make_request 8082 "/health" 50

echo "Simulating order requests..."
make_request 8081 "/health" 50

wait

echo "Load simulation completed."
echo "Check Prometheus/Grafana for metrics on CPU usage, response times, and error rates."