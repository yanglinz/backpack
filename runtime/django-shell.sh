#!/usr/bin/env bash
set -euo pipefail
IFS=$'\n\t'

. "$(dirname "$0")/env-loader.sh"

pipenv shell
