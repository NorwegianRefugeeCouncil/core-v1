#!/bin/bash

set -e
set -o pipefail

echo "Running tests..."
go test ./...
