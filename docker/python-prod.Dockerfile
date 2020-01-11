FROM python:3.7-alpine3.9

ENV PYTHONUNBUFFERED 1

# Install dependencies
RUN apk add --no-cache \
  bash \
  ca-certificates \
  cyrus-sasl-dev \
  g++ \
  gcc \
  libffi-dev \
  libxslt-dev \
  linux-headers \
  openssl \
  openssl-dev \
  python3-dev \
  postgresql-dev \
  postgresql-client

# Install berglas
COPY --from=gcr.io/berglas/berglas:0.5.0 /bin/berglas /bin/berglas

# Install application dependencies
RUN mkdir -p /app
WORKDIR /app

RUN pip install --no-cache-dir --trusted-host pypi.python.org pipenv
COPY Pipfile /app/
COPY Pipfile.lock /app/
RUN pipenv install --system --deploy
RUN pip install uwsgi==2.0.18

# Create users and directories
RUN mkdir -p /app \
  && mkdir -p /home/app \
  && addgroup backpack \
  && adduser -D -u 1000 -G backpack backpack \
  && chown --recursive backpack:backpack /app \
  && chown --recursive backpack:backpack /home/app

# Switch to non-root user
USER backpack
ENV HOME /home/app

# Copy application code
COPY --chown=backpack:backpack . /app

# Entrypoint
STOPSIGNAL SIGTERM
ENTRYPOINT [".backpack/runtime/django-prod.sh"]
