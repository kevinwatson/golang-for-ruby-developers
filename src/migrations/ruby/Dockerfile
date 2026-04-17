FROM ruby:4.0.1

RUN apt-get update && apt-get install -qq -y --no-install-recommends \
    build-essential nodejs vim

WORKDIR /opt

RUN gem install rails -v 8.1
RUN rails new migrator-app
WORKDIR /opt/migrator-app
RUN bundle install

