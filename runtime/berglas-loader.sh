#!/usr/bin/env bash
set -euo pipefail
IFS=$'\n\t'

echo "Plain env"
env

echo "Berglas env"
berglas exec -- env

BERGLAS_ENV_PY="$(dirname "$0")/berglas_env.py"
berglas access "${BERGLAS_BUCKET}/BERGLAS_APP_JSON" > /tmp/berglas-app.json
cat /tmp/berglas-app.json
eval $(python3 "$BERGLAS_ENV_PY" /tmp/berglas-app.json)

echo "Plain env"
env
