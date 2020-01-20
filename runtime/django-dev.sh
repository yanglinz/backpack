#!/usr/bin/env bash
set -euo pipefail
IFS=$'\n\t'

. "$(dirname "$0")/env-loader.sh"

# Run django migrations
pipenv run python manage.py migrate

# Configure nginx and start the development server
cp .backpack/docker/nginx/nginx-dev.conf /etc/nginx/nginx.conf
/usr/bin/supervisord
