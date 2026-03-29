### Chapter 9 - Unit Testing

> 

## Introduction

Unit testing is a method of debugging or exercising our code to ensure its correctness. It's a way to write scaffolding around our app that helps ensure the future changes won't break current expectations by regression testing. Ruby has several testing frameworks (we'll use RSpec) and Go provides built-in support with the `testing` package.

## Examples

### Ruby

Let's build a Docker image to run our environment. We'll use the `rspec` gem which is included in this version of Rails.

```dockerfile
# golang-for-ruby-developers/unit-testing/ruby/Dockerfile

FROM ruby:4.0.1

RUN apt-get update && apt-get install -qq -y --no-install-recommends \
    build-essential nodejs vim

RUN gem install rails -v 8.1
RUN rails new console-app -T # don't include the default Test::Unit files
WORKDIR /opt/console-app
# add 'gem "rspec-rails"' to the test group in the Gemfile
RUN sed -i '/^group :development, :test do/,/^end/ s/^end/  gem "rspec-rails"\nend/' Gemfile
RUN bundle install
RUN rails generate rspec:install
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

Now let's start the container and log in.

```bash
docker-compose run ruby bash
```

Let's add a validation to require the three fields we added to the employee model. We'll then write a unit test for our model. We've added vim to the docker image but feel free to include your editor of choice.

```ruby
# app/models/employee.rb

class Employee < ApplicationRecord
  validates :guid, presence: true
end
```

After editing and saving the employee.rb file, now let's add a unit test.

```ruby
# spec/models/employee_spec.rb

require 'rails_helper'

RSpec.describe Employee, type: :model do
  subject(:employee) { described_class.new(guid: nil) }
  before { employee.valid? }

  describe 'validations' do
    it 'requires a guid' do
      expect(employee.errors[:guid]).to be_present
    end
  end
end
```

After editing and saving the `employee_spec.rb` file, let's run the test.

```bash
rspec spec/models/
.

Finished in 0.01174 seconds (files took 0.72054 seconds to load)
1 example, 0 failures
```

We can see from the results that one test passed and zero tests failed. If we were to remove the `validates :guid, presence: true` line from the `employee.rb` class at a later time, our tests would fail, indicating that our past assumptions are now incorrect.

### Go

First, we'll use a Docker image.

```dockerfile
# golang-for-ruby-developers/debugger/go/Dockerfile

FROM golang:1.25

WORKDIR /opt/app
RUN go mod init example.com/employees
COPY employees.go ./
RUN go get ./...
```

Next, we'll use a Docker Compose file to configure our environment.

```yaml
# golang-for-ruby-developers/debugger/go/docker-compose.yml

services:
  go:
    build:
      context: .
    stdin_open: true
    tty: true
```

And we'll need some code.

```golang
# golang-for-ruby-developers/debugger/go/employees.go

package main

import (
  "context"
  "encoding/json"
  "fmt"
  "gorm.io/driver/sqlite"
  "gorm.io/gorm"
  "log"
  "net/http"
)

type Employee struct {
  gorm.Model
  guid      string
  firstName string
  lastName  string
}

func employeeIndexHandler(w http.ResponseWriter, req *http.Request) {
  //list the employees

  db, err := gorm.Open(sqlite.Open("employees.db"), &gorm.Config{})
  if err != nil {
    log.Fatal("Could not connect to database employees.db: ", err)
  }

  ctx := context.Background()

  records, err := gorm.G[Employee](db).Find(ctx)
  if err != nil {
    fmt.Printf("Could not connect to database %s\n", err)
  }

  js, err := json.Marshal(records)
  if err != nil {
    http.Error(w, err.Error(), http.StatusInternalServerError)
    return
  }

  w.Header().Set("Content-Type", "application/json")
  w.Write(js)
}

func main() {
  db, err := gorm.Open(sqlite.Open("employees.db"), &gorm.Config{})
  if err != nil {
    log.Fatal("Could not connect to database employees.db: ", err)
  }

  db.AutoMigrate(&Employee{})

  http.HandleFunc("/employees", employeeIndexHandler)
  fmt.Println("Server starting on port 8080")
  err = http.ListenAndServe(":8080", nil)
  if err != nil {
    log.Fatal("Server exited with error: ", err)
  }
}
```

Now let's run the code.

```bash
docker-compose run go bash
```


```golang
```

## References


## Wrap Up


[Next >>](100-chapter-09.md)

