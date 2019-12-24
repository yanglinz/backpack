#!/usr/bin/env bash
set -euo pipefail
IFS=$'\n\t'

RELEASE_BRANCH="master"
SOURCE_IMAGE="source-image"
DOCKER_REPO="gcr.io/${GCP_PROJECT_ID}/${APP_NAME}"

function build_release() {
  docker build \
    -f .backpack/configs/docker/python-prod.Dockerfile \
    --tag "$SOURCE_IMAGE" \
    .
}

function publish_release() {
  gcloud auth configure-docker
  docker tag "$SOURCE_IMAGE" "$DOCKER_REPO"
  docker push "$DOCKER_REPO"
}

build_release

if [ "$GITHUB_REF" == "refs/heads/${RELEASE_BRANCH}" ]; then
  publish_release
else
  echo "Not on ${RELEASE_BRANCH} branch. Nothing to build."
fi
