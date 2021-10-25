#!/usr/bin/env bash

echo "> Install dependencies"

GOOS=$(go env GOOS) GOARCH=$(go env GOARCH) GO111MODULE=on \
  go mod download \
  $@

echo "> Done"