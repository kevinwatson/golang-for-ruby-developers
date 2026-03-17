### Chapter 7 - Console

> 

## Introduction

## Comparing Languages

Rails console is a commonly used REPL (read eval print loop) feature. It builds on IRB which is a built in Ruby tool for running code in an interactive shell environment. With the dynamic nature of Ruby, you can create new classes or access existing classes and call methods on those classes and objects. In the shell, you also have access to the rest of the environment that's available at runtime, so you can access resources to help you test out ideas while in development or access databases in a test or production environment.

Go is a compiled language, which makes running a REPL more difficult. While not a built-in feature, that hasn't stopped the Go community from building their own. With a few limitations, the third party gomacro and yaegi projects work similar to Rails console.

## Examples

### Ruby on Rails

Let's build a Docker image to run our environment.

```dockerfile
# golang-for-ruby-developers/console/ruby/Dockerfile

FROM ruby:4.0.1

RUN apt-get update && apt-get install -qq -y --no-install-recommends \
    build-essential \
    nodejs

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
# golang-for-ruby-developers/console/ruby/docker-compose.yml

services:
  ruby:
    build:
      context: .
    stdin_open: true
    tty: true
```

Now, let's start the app from our terminal.

```bash
docker-compose run ruby bash
```

Now that we have a shell in a Ruby environment, let's run a few ActiveRecord query interface commands to interact with data in our database.

```ruby
# rails console
Loading development environment (Rails 8.1.2)
console-app(dev):001> Employee.first
  Employee Load (0.3ms)  SELECT "employees".* FROM "employees" ORDER BY "employees"."id" ASC LIMIT 1 /*application='ConsoleApp'*/
=> nil
console-app(dev):002> Employee.create(guid: SecureRandom.uuid, first_name: "George", last_name: "Jetson")
  TRANSACTION (0.2ms)  BEGIN immediate TRANSACTION /*application='ConsoleApp'*/
  Employee Create (1.1ms)  INSERT INTO "employees" ("guid", "first_name", "last_name", "created_at", "updated_at") VALUES ('072dd52c-3906-452a-a9f8-ce7eae34566c', 'George', 'Jetson', '2026-03-17 08:33:05.320604', '2026-03-17 08:33:05.320604') RETURNING "id" /*application='ConsoleApp'*/
  TRANSACTION (4.3ms)  COMMIT TRANSACTION /*application='ConsoleApp'*/
=>
#<Employee:0x0000ffff94fad7c0
 id: 1,
 guid: "072dd52c-3906-452a-a9f8-ce7eae34566c",
 first_name: "George",
 last_name: "Jetson",
 created_at: "2026-03-17 08:33:05.320604000 +0000",
 updated_at: "2026-03-17 08:33:05.320604000 +0000">
console-app(dev):003> Employee.first
  Employee Load (0.6ms)  SELECT "employees".* FROM "employees" ORDER BY "employees"."id" ASC LIMIT 1 /*application='ConsoleApp'*/
=>
#<Employee:0x0000ffff9500b050
 id: 1,
 guid: "072dd52c-3906-452a-a9f8-ce7eae34566c",
 first_name: "George",
 last_name: "Jetson",
 created_at: "2026-03-17 08:33:05.320604000 +0000",
 updated_at: "2026-03-17 08:33:05.320604000 +0000">
```

### Go

#### Gomacro

```bash
gomacro
```

```go
```

## References

* https://github.com/cosmos72/gomacro
* https://github.com/traefik/yaegi
* https://guides.rubyonrails.org/command_line.html#bin-rails-console

## Wrap Up


[Next >>](090-chapter-08.md)
