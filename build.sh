#!/bin/bash

echo "Generating database models..."
# Cleanup all files other than "queries.go", "utils.go" and "queries_input.go" in all the sub-dirs of db/dbmodels/
for dir in ./internal/app/database/dbmodels/*/;
  do find $dir -type f ! -name 'queries.go' ! -name 'utils.go' ! -name 'queries_input.go' -delete;
done
sqlc generate

echo "Building go binary..."
go build -o api cmd/api/main.go

echo "Done";