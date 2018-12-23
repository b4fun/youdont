#!/bin/bash

set -e

CURRENT_DIR="$(cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd)"
PROJECT_ROOT="$(dirname "$CURRENT_DIR")"
CMD_ROOT="$PROJECT_ROOT/cmd"
DIST_ROOT="$PROJECT_ROOT/dist"

# build a function binary
#
#   build {FUNCTION_NAME}
function build() {
    local -r func_name="$1"
    local -r binary_name="$func_name-linux-amd64"
    local -r zip_name="$func_name.zip"

    echo "buidling $func_name..."

    mkdir -p "$DIST_ROOT"
    GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build \
        -o "$DIST_ROOT/$binary_name" \
        "$CMD_ROOT/$func_name"
    echo "built binary: $binary_name"

    zip -qjD "$DIST_ROOT/$zip_name" "$DIST_ROOT/$binary_name"
    echo "built zip archive: $zip_name"
}

func_name=$1
if [[ -z "$func_name" ]]
then
    echo 'Usage: build-function.sh function-name'
    exit -1
fi

build "$func_name"
