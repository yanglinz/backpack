#!/usr/bin/env bash
set -euo pipefail
IFS=$'\n\t'

function set_buildpack() {
  dokku config:set "$APP_NAME" BUILDPACK_URL=https://github.com/heroku/heroku-buildpack-python.git#v184
}

function set_vars() {
  local ENV_SOURCE="/opt/backpack-app/var/env/production.json"

  echo "Setting ${APP_NAME} application configs..."

  for name in $(jq --raw-output 'keys | .[]' "$ENV_SOURCE"); do
    value=$(jq --raw-output ".${name}" "$ENV_SOURCE")
    dokku config:set --no-restart "$APP_NAME" "$name"="$value"
  done
}

function deploy() {
  cat /tmp/app-artifact.tar.gz | dokku tar:in "$APP_NAME"
}

set_buildpack
set_vars
deploy

dokku config:set "$APP_NAME" TIMESTAMP=$(date +%s)
dokku domains:enable backpack-example
