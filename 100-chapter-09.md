### Chapter 9 - Unit Testing

> 

## Introduction

Unit testing is a method of debugging or exercising our code to ensure its correctness. It's a way to write scaffolding around our app that helps ensure the future changes won't break current expectations by regression testing. Ruby has several testing frameworks (we'll use RSpec) and Go provides built-in support with the `testing` package.

## Examples

For both languages, we'll set up and run a simple test for our Employees model.

### Ruby

Ruby on Rails convention suggests adding tests to a `test` or `spec` directory. In the examples below, we'll use the [rspec](https://rspec.info) gem.

Let's build a Docker image to run our environment.

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

Go's [testing](https://pkg.go.dev/testing) package convention suggests adding your test files to the same folders as your code but with the `_test.go` file suffix. For example, if we have an `employees.go` file, we would also have an `employees_test.go` file in the same folder.

First, we'll use a Docker image.

```dockerfile
# golang-for-ruby-developers/unit-testing/go/Dockerfile

FROM golang:1.25

WORKDIR /opt/app
RUN go mod init example.com/employees
COPY employees* ./
RUN go get ./...
```

Next, we'll use a Docker Compose file to configure our environment.

```yaml
# golang-for-ruby-developers/unit-testing/go/docker-compose.yml

services:
  go:
    build:
      context: .
    stdin_open: true
    tty: true
```

And we'll need some code.

```golang
# golang-for-ruby-developers/unit-testing/go/employees.go

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

func EmployeeIndexHandler(w http.ResponseWriter, req *http.Request) {
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

  http.HandleFunc("/employees", EmployeeIndexHandler)
  fmt.Println("Server starting on port 8080")
  err = http.ListenAndServe(":8080", nil)
  if err != nil {
    log.Fatal("Server exited with error: ", err)
  }
}
```

Next, let's write a test to test our `employeeIndexHandler` function.

```golang
# golang-for-ruby-developers/unit-testing/go/employees_test.go

package main

import (
  "gorm.io/driver/sqlite"
  "gorm.io/gorm"
  "log"
  "net/http"
  "net/http/httptest"
  "testing"
)

func TestEmployeeIndexHandler(t *testing.T) {
  // initialize the database
  db, err := gorm.Open(sqlite.Open("employees.db"), &gorm.Config{})
  if err != nil {
    log.Fatal("Could not connect to database employees.db: ", err)
  }

  db.AutoMigrate(&Employee{})

  // create a 'get' request
  req, err := http.NewRequest("GET", "/employees", nil)
  if err != nil {
    t.Fatal(err)
  }
  // create a ResponseRecorder for the http.ResponseWriter
  rr := httptest.NewRecorder()
  handler := http.HandlerFunc(EmployeeIndexHandler)
  handler.ServeHTTP(rr, req)
  if status := rr.Code; status != http.StatusOK {
    t.Errorf("handler returned wrong status code: got %v, want %v", status, http.StatusOK)
  }
  expected := `[]`
  if rr.Body.String() != expected {
    t.Errorf("handler returned unexpected body: got %v, want %v", rr.Body.String(), expected)
  }
}
```

Now let's run the container and the tests.

```bash
docker-compose run go bash

root@e656e3284925:/opt/app# go test
PASS
ok  	example.com/employees	0.014s
```

## References

* https://go.dev/doc/tutorial/add-a-test
* https://guides.rubyonrails.org/testing.html

## Wrap Up

As we have discussed, unit testing is important for the maintenance of our application. It prevents us from changing the behavior of our app unexpectedly as we add features or fix bugs. In the next chapter, we'll add logs to our app.

[Next >>](110-chapter-10.md)

