FROM ruby:3.1.0-slim

# install rails dependencies
RUN apt-get update -qq && apt-get install -y build-essential libpq-dev \
  git unixodbc unixodbc-dev tdsodbc \
  postgresql-client libxrender1 libfontconfig wget ghostscript

RUN gem install bundler:1.17.3

RUN curl -fsSL https://deb.nodesource.com/setup_16.x | bash -
RUN apt-get install -y nodejs npm

RUN npm install -g yarn

# create a folder /myapp in the docker container and go into that folder
RUN mkdir /myapp
WORKDIR /myapp

# Copy the whole app
COPY . /myapp
