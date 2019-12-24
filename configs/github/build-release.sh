#!/usr/bin/env bash
set -euo pipefail
IFS=$'\n\t'

function build_release() {
  docker build \
    -f .backpack/configs/docker/python-prod.Dockerfile \
    --tag application-release \
    .
}

function publish_release() {
  gcloud auth configure-docker
  # docker tag [SOURCE_IMAGE] [HOSTNAME]/[PROJECT-ID]/[IMAGE]
  # docker tag [SOURCE_IMAGE] [HOSTNAME]/[PROJECT-ID]/[IMAGE]:[TAG]
  # SOURCE_IMAGE = application-release
  # HOSTNAME = gcr.io
  # PROJECT-ID = default-123
  # IMAGE = application-name
}

build_release
