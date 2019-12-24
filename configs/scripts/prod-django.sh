#!/usr/bin/env bash
set -euo pipefail
IFS=$'\n\t'

python manage.py runserver "0.0.0.0:$PORT"
