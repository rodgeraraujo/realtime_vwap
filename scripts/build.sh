#!/usr/bin/env bash

echo "> Build app"

GOOS=$(go env GOOS) GOARCH=$(go env GOARCH) \
  go build ./cmd/vwap \
  $@
