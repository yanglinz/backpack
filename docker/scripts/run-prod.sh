#!/usr/bin/env bash
set -euo pipefail
IFS=$'\n\t'

# Manage static files
python manage.py collectstatic

# Start production server
uwsgi --ini /app/.backpack/docker/uwsgi/uwsgi.ini
