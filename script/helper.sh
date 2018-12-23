#!/bin/bash

set -o errexit
set -o nounset
set -o pipefail

function helper::date() {
    date "+%Y%m%d"
}

function helper::git_revision() {
    local -r head_commit=$(git rev-parse --short HEAD)
    local -r suffix=$(test -n "$(git status --porcelain)" && echo "-dev" || echo '')
    echo "${head_commit}${suffix}"
}
