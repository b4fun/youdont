#!/bin/bash

set -e
set -o errexit
set -o pipefail


CURRENT_DIR="$(cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd)"
# shellcheck source=./helper.sh
source "${CURRENT_DIR}/helper.sh"

PROJECT_ROOT="$(dirname "$CURRENT_DIR")"
DIST_ROOT="$PROJECT_ROOT/dist"


# deploy a function
#
#   deploy {FUNCTION_NAME}
function deploy() {
    local -r func_name="$1"
    local -r zip_name="$func_name.zip"

    aws lambda update-function-code \
        --function-name "$func_name" \
        --zip-file "fileb://$DIST_ROOT/$zip_name" \
        --publish
    echo "deployed $func_name"
}

func_name=${1:-''}
if [[ -z "$func_name" ]]
then
    echo -n 'func name: '
    read -r func_name
fi

deploy "$func_name"
