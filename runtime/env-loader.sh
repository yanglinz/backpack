#!/usr/bin/env bash
set -euo pipefail
IFS=$'\n\t'

eval $(python3 "$(dirname "$0")/load_env.py" /app/etc/development.json)
