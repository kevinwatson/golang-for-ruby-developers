### Chapter 8 - Debugging

> 

## Introduction

Debugging is the process of finding and fixing errors in your software. There are a number of ways to debug an application, including printing log lines and pausing and stepping through the code in a running process. Each technique has their own advantages, and Ruby and Go support these techniques with either built-in or third-party libraries.

In a later chapter, we'll focus on logging but in this chapter, we'll focus on starting and navigating debugging sessions in Ruby and Go.

## Examples

### Ruby

Let's build a Docker image to run our environment. We'll use the `debug` gem which is included in this version of Rails.

```dockerfile
# golang-for-ruby-developers/debugger/ruby/Dockerfile

FROM ruby:4.0.1

RUN apt-get update && apt-get install -qq -y --no-install-recommends \
    build-essential nodejs vim

WORKDIR /opt
RUN gem install rails -v 8.1
RUN rails new console-app
WORKDIR /opt/console-app
RUN bundle install
RUN rails generate scaffold Employee guid:string first_name:string last_name:string
RUN rails db:migrate
```

Let's use a Docker Compose file to configure our environment.

```yaml
# golang-for-ruby-developers/debugger/ruby/docker-compose.yml

services:
  ruby:
    build:
      context: .
    stdin_open: true
    tty: true
```

First, let's add a `debugger` statement to pause program execution and start a debugging session. Vim is included in this Docker image but you can add another editor to the image if you prefer.

```bash
docker-compose run -p 3000:3000 ruby bash

# vim app/controllers/employees_controller.rb
def index
  @employees = Employee.all
  debugger
end
```

After saving and closing the `employees_controller.rb` file, let's run a command to start the Rails app.

```bash
# bundle exec puma -C config/puma.rb
```

Open a browser window and navigate to the employees route in our running app, which will then hit our `debugger` command and start the debugging session.

http://localhost:3000/employees

### Go

## References

* https://guides.rubyonrails.org/debugging_rails_applications.html

## Wrap Up


[Next >>](100-chapter-09.md)

