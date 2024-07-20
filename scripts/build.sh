#!/bin/bash

set -e
export GOPATH=$(go env GOPATH)
export PATH=$GOPATH/bin:$PATH

run_tests() {
  echo "Running tests..."
  go test ./...
}

run_verification() {
  echo "Running golint..."

  go install golang.org/x/lint/golint@latest
  golint ./...

  echo "---------"
  echo "Running gosec..."

  go install github.com/securego/gosec/v2/cmd/gosec@latest
  gosec ./...
}

build_project() {
  echo "Building project..."
  go build -v -o bin/aw2 main.go
}

run_tests
run_verification
build_project

echo "Build successful!"

