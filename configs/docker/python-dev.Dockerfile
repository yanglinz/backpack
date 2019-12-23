FROM python:3.7

# Install watchman
# This step takes a LONG time, so cache this layer first
ARG DISABLE_WATCHMAN
COPY ./.backback/configs/watchman/install-watchman.sh /tmp/
RUN /tmp/install-watchman.sh

ENV PYTHONUNBUFFERED 1

# Install dependencies
RUN apt-get update && apt-get install -y \
    git \
    postgresql-client

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
COPY .backpack/configs/watchman/watchman.json /etc/

# Copy application code
COPY . /app
