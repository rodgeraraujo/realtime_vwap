#!/usr/bin/env bash

set -o errexit
set -o nounset
set -o pipefail

echo "> Test Cover"

GO111MODULE=on ginkgo -cover -timeout=2m -race  ../...

REPO_ROOT="$(git rev-parse --show-toplevel)"
COVERPROFILE="$REPO_ROOT/test.coverprofile"
COVERPROFILE_TMP="$REPO_ROOT/coverprofile.tmp"
COVERPROFILE_HTML="$REPO_ROOT/coverage.html"

echo "mode: set" > "$COVERPROFILE_TMP"
find . -name "*.coverprofile" -type f | xargs cat | grep -v mode: | sort -r | awk '{if($1 != last) {print $0;last=$1}}' >> "$COVERPROFILE_TMP"
cat "$COVERPROFILE_TMP" | grep -vE "\.pb\.go|zz_generated" > "$COVERPROFILE"
rm -rf "$COVERPROFILE_TMP"
go tool cover -html="$COVERPROFILE" -o="$COVERPROFILE_HTML"

go tool cover -func="$COVERPROFILE"
