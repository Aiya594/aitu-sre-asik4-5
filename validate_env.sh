#!/bin/bash


echo "Validating environment configuration..."

if [ ! -f ".env" ]; then
    echo "ERROR: .env file not found"
    exit 1
fi

source .env

if [ -z "$DB_HOST" ]; then
    echo "ERROR: DB_HOST is not set"
    exit 1
fi

if [ -z "$DB_USER" ]; then
    echo "ERROR: DB_USER is not set"
    exit 1
fi

if [ -z "$DB_PASSWORD" ]; then
    echo "ERROR: DB_PASSWORD is not set"
    exit 1
fi

if [ -z "$DB_NAME" ]; then
    echo "ERROR: DB_NAME is not set"
    exit 1
fi

if [ -z "$DB_PORT" ]; then
    echo "ERROR: DB_PORT is not set"
    exit 1
fi

# Validate service ports
if [ -z "$AUTH_SERVICE_PORT" ]; then
    echo "ERROR: AUTH_SERVICE_PORT is not set"
    exit 1
fi

if [ -z "$PRODUCT_SERVICE_PORT" ]; then
    echo "ERROR: PRODUCT_SERVICE_PORT is not set"
    exit 1
fi

if [ -z "$ORDER_SERVICE_PORT" ]; then
    echo "ERROR: ORDER_SERVICE_PORT is not set"
    exit 1
fi

if [ -z "$FRONTEND_PORT" ]; then
    echo "ERROR: FRONTEND_PORT is not set"
    exit 1
fi

# Validate internal URLs
if [ -z "$PRODUCT_SERVICE_URL" ]; then
    echo "ERROR: PRODUCT_SERVICE_URL is not set"
    exit 1
fi

# Validate monitoring ports
if [ -z "$PROMETHEUS_PORT" ]; then
    echo "ERROR: PROMETHEUS_PORT is not set"
    exit 1
fi

if [ -z "$GRAFANA_PORT" ]; then
    echo "ERROR: GRAFANA_PORT is not set"
    exit 1
fi

echo "All environment variables are properly configured."
echo "Validation successful!"