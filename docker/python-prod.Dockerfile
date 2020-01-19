FROM python:3.7

ENV PYTHONUNBUFFERED 1

# Install dependencies
RUN apt-get update && apt-get install -y \
  gettext-base \
  supervisor \
  nginx \
  git \
  postgresql-client

# Stop nginx
RUN service nginx stop

# Install berglas
COPY --from=gcr.io/berglas/berglas:0.5.0 /bin/berglas /bin/berglas

# Setup the working directory
RUN mkdir /app && mkdir /home/app
WORKDIR /app
ENV HOME /home/app

# Install custom dependencies
COPY scripts/docker /app/scripts/docker
COPY .backpack/docker/scripts/install-extra-deps.sh /tmp/
RUN /tmp/install-extra-deps.sh

# Install application dependencies
RUN pip install --no-cache-dir --trusted-host pypi.python.org pipenv
COPY Pipfile /app/
COPY Pipfile.lock /app/
RUN pipenv install --system --deploy
RUN pip install uwsgi==2.0.18

# Copy configuration
COPY .backpack/docker/nginx/nginx-prod.tmpl.conf /etc/nginx/nginx.conf
COPY .backpack/docker/supervisord/supervisord-prod.conf /etc/supervisord.conf

# Copy application code
COPY . /app

# Entrypoint
STOPSIGNAL SIGTERM
ENTRYPOINT [".backpack/runtime/django-prod.sh"]
