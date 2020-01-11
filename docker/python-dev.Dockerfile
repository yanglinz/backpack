FROM python:3.7

# Install watchman
# This step takes a LONG time, so cache this layer first
ARG DISABLE_WATCHMAN
COPY .backpack/docker/watchman/install-watchman.sh /tmp/
RUN /tmp/install-watchman.sh

ENV PYTHONUNBUFFERED 1

# Install gcloud
COPY .backpack/docker/scripts/install-gcloud.sh /tmp/
RUN /tmp/install-gcloud.sh

# Install dependencies
RUN apt-get update && apt-get install -y \
  git \
  postgresql-client

# Install berglas
COPY --from=gcr.io/berglas/berglas:0.5.0 /bin/berglas /bin/berglas

# Setup the working directory
RUN mkdir /app && mkdir /home/app
WORKDIR /app
ENV HOME /home/app

# Install application dependencies
RUN pip install --no-cache-dir --trusted-host pypi.python.org pipenv
COPY Pipfile /app/
COPY Pipfile.lock /app/
RUN pipenv install --dev

# Copy configurations
COPY .backpack/docker/watchman/watchman.json /etc/

# Copy application code
COPY . /app