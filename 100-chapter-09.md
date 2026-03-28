### Chapter 9 - Unit Testing

> 

## Introduction


## Examples

### Ruby

Let's build a Docker image to run our environment. We'll use the `debug` gem which is included in this version of Rails.

```dockerfile
# golang-for-ruby-developers/unit-testing/ruby/Dockerfile

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

