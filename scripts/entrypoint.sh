#!/bin/sh

# Wait for PostgreSQL to be up
/usr/local/bin/wait-for-postgres.sh db "$@"

# Apply migration
echo "Apply all pending migrations..."
/usr/local/bin/atlas migrate apply --url "postgresql://$POSTGRES_USER:$POSTGRES_PASSWORD@db:5432/$POSTGRES_DB?sslmode=disable"

# Run seeds
echo "Run all pending seeds..."
go run ./seeds

# Run the Go application
echo "Starting Go application..."
air
