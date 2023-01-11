#!/bin/sh
set -ex

go test -count=10 "$@" ./...
go test -race -count=10 "$@" ./...

# Configure linters in .golangci.yml
GOBIN="$PWD/bin" go install github.com/golangci/golangci-lint/cmd/golangci-lint@v1.50.1
./bin/golangci-lint run

