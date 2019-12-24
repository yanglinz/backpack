#!/usr/bin/env bash
set -euo pipefail
IFS=$'\n\t'

function build_release() {
  docker build \
    -f .backpack/configs/docker/python-prod.Dockerfile \
    --tag hello-world \
    .
}

build_release
