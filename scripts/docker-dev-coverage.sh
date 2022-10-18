#!/bin/bash

set -e
set -o pipefail

# It will be available in the "dev" target of the Dockerfile

SCRIPT_DIR=$(cd -- "$(dirname -- "${BASH_SOURCE[0]}")" &>/dev/null && pwd)
ROOT_DIR=$(cd -- "${SCRIPT_DIR}/.." &>/dev/null && pwd)
OUTPUT_DIR="${ROOT_DIR}/reports"

case "$1" in
"--output|-o")
  echo "Output directory: $2"
  OUTPUT_DIR="$2"
  shift 2
  ;;
esac

mkdir -p "$OUTPUT_DIR"

echo "Running tests..."
mkdir -p "${OUTPUT_DIR}"
go test ./...
go test ./... -coverpkg=./... -v -coverprofile="${OUTPUT_DIR}/coverage.txt" -covermode count 2>&1 | go-junit-report >"${OUTPUT_DIR}/report.xml"
rc=${PIPESTATUS[0]}

echo "Generating coverage report..."
gocov convert "${OUTPUT_DIR}/coverage.txt" >"${OUTPUT_DIR}/coverage.json"
gocov-xml <"${OUTPUT_DIR}/coverage.json" >"${OUTPUT_DIR}/coverage.xml"
go tool cover -html="${OUTPUT_DIR}/coverage.txt" -o "${OUTPUT_DIR}/coverage.html"

if [ "$rc" -ne 0 ]; then
  echo "Tests failed"
  exit "$rc"
else
  echo "Tests passed"
fi