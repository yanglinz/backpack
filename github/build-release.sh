#!/usr/bin/env bash
set -euo pipefail
IFS=$'\n\t'

RELEASE_BRANCH="master"
SOURCE_IMAGE="source-image"
DOCKER_REGISTRY="gcr.io/${GCP_PROJECT_ID}/${APP_NAME}"
RELEASE_TAG="$(. "$(dirname "$0")/hash-files.sh")"

function build_release() {
  docker build \
    -f .backpack/docker/python-prod.Dockerfile \
    --tag "$SOURCE_IMAGE" \
    .
}

function publish_release() {
  gcloud auth configure-docker
  docker tag "$SOURCE_IMAGE" "$DOCKER_REGISTRY"
  docker tag "$SOURCE_IMAGE" "$DOCKER_REGISTRY:${RELEASE_TAG}"
  docker push "$DOCKER_REGISTRY"
  docker push "$DOCKER_REGISTRY:${RELEASE_TAG}"
}

build_release

if [ "$GITHUB_REF" == "refs/heads/${RELEASE_BRANCH}" ]; then
  publish_release
else
  echo "Not on ${RELEASE_BRANCH} branch. Nothing to build."
fi
