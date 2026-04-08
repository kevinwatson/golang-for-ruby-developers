FROM ruby:4.0.1

RUN apt-get update && apt-get install -qq -y --no-install-recommends \
    build-essential nodejs vim

WORKDIR /opt

RUN gem install rails -v 8.1
RUN rails new console-app -T # don't include the default Test::Unit files
WORKDIR /opt/console-app
# blah blah blah
RUN sed -i '/^group :development, :test do/,/^end/ s/^end/  gem "rspec-rails"\nend/' Gemfile
RUN bundle install
RUN rails generate rspec:install
RUN rails generate scaffold Employee guid:string first_name:string last_name:string
RUN rails db:migrate

