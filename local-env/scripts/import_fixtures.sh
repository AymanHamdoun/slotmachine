#!/bin/bash

DB_NAME="default"
DB_USER="root"
DB_PASS="dbuserpassword"
DB_HOST="127.0.0.1"
DB_PORT=4306

FIXTURE_DIR="$(dirname "$0")/../../internal/app/database/fixtures"

if [ ! -d "$FIXTURE_DIR" ]; then
  echo "❌ Fixture directory not found: $FIXTURE_DIR"
  exit 1
fi

for file in "$FIXTURE_DIR"/*.sql; do
  [ -e "$file" ] || continue
  echo "Loading $(basename "$file")..."
  mysql -h "$DB_HOST" -P $DB_PORT -u "$DB_USER" -p"$DB_PASS" "$DB_NAME" < "$file"
done

echo "✅ All fixtures loaded into $DB_NAME"