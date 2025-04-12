#!/bin/bash
# Usage: ./dump-fixtures.sh users products

DB_NAME="default"
DB_USER="root"
DB_PASS="dbuserpassword"
DB_HOST="127.0.0.1"
DB_PORT=4306
FIXTURE_DIR="$(dirname "$0")/../../internal/app/database/fixtures"


if [ "$#" -eq 0 ]; then
  echo "❌ Please specify tables to dump."
  exit 1
fi

mkdir -p fixtures

for table in "$@"; do
  echo "Dumping table: $table"
  mysqldump -h "$DB_HOST" -P $DB_PORT -u "$DB_USER" -p"$DB_PASS" "$DB_NAME" "$table" --skip-extended-insert > "$FIXTURE_DIR/${table}.sql"
done

echo "✅ Fixtures dumped to $FIXTURE_DIR"