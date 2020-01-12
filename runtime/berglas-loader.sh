#!/usr/bin/env bash
set -euo pipefail
IFS=$'\n\t'

berglas access "$BERGLAS_SECRET_PATH" > /tmp/berglas-app.json
eval $(python3 "$(dirname "$0")/load_env.py" /tmp/berglas-app.json)
