#!/usr/bin/env bash
set -euo pipefail
IFS=$'\n\t'

. "$(dirname "$0")/berglas-loader.sh"

python manage.py runserver "0.0.0.0:$PORT"
