#!/usr/bin/env bash
set -euo pipefail
golangci-lint version || (echo "golangci-lint ausente"; exit 1)
golangci-lint run
