#!/usr/bin/env bash
set -euo pipefail
IFS=$'\n\t'

. "$(dirname "$0")/env-loader.sh"

pipenv run python manage.py runserver 0.0.0.0:8000
