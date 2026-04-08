FROM ruby:4.0.1

RUN apt-get update && apt-get install -qq -y --no-install-recommends \
    build-essential nodejs vim

WORKDIR /opt

RUN gem install rails -v 8.1
RUN rails new generator-app
WORKDIR /opt/generator-app
RUN bundle install
RUN rails generate scaffold Employee guid:string first_name:string last_name:string
RUN rails db:migrate

