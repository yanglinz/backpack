#!/usr/bin/env bash
set -euo pipefail
IFS=$'\n\t'

pipenv run python manage.py runserver 0.0.0.0:8000
