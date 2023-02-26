#!/bin/sh

set -e
echo "run db migration"
migrate -path /app/migration -database "$DB_URL" -verbose up
echo "start the app"
exec "$@"