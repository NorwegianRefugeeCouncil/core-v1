#!/bin/bash

# This script is used as the entrypoint for the core "dev" container.
# For now, only the "coverage" target is supported, but more targets
# related to development can be added in the future.

set -e
set -o pipefail

SCRIPT_DIR=$(cd -- "$(dirname -- "${BASH_SOURCE[0]}")" &>/dev/null && pwd)

case "$1" in
"coverage")
  shift
  "${SCRIPT_DIR}"/docker-dev-coverage.sh "$@"
  ;;
"test")
  shift
  "${SCRIPT_DIR}"/docker-dev-tests.sh "$@"
  ;;
esac
