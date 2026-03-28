### Chapter 8 - Debugging

> 

## Introduction

Debugging is the process of finding and fixing errors in your software. There are a number of ways to debug an application, including printing log lines and pausing and stepping through the code in a running process. Each technique has their own advantages, and Ruby and Go support these techniques with either built-in or third-party libraries.

In a later chapter, we'll focus on logging but in this chapter, we'll focus on starting and navigating debugging sessions in Ruby and Go.

## Examples

Many interactive development environments (IDEs) can be used to start a debugging session. Here we'll set up environments in the terminal which should work on any operating system.

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

First, let's add a `debugger` statement to pause program execution and start a debugging session. Vim is included in this Docker image but you can add another editor to the image if you prefer. We'll need to use the `-p` or `--port` flag to open up port 3000 in and out of the container.

```bash
docker-compose run -p 3000:3000 ruby bash
```

```ruby
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

We should see output like the following in our terminal.

```bash
Puma starting in single mode...
* Puma version: 7.2.0 ("On The Corner")
* Ruby version: ruby 4.0.1 (2026-01-13 revision e04267a14b) +PRISM [aarch64-linux]
*  Min threads: 3
*  Max threads: 3
*  Environment: development
*          PID: 20
* Listening on http://0.0.0.0:3000
Use Ctrl-C to stop
```

Open a browser window and navigate to the employees index route in our running app. This will then pause on our `debugger` command and start a debugging session.

http://localhost:3000/employees

Switching back to your terminal, you should see that it paused on the `debugger` line with a prompt. You can now check variables, print headers, set breakpoints, show a backtrace, etc.

```ruby
[2, 11] in /opt/console-app/app/controllers/employees_controller.rb
     2|   before_action :set_employee, only: %i[ show edit update destroy ]
     3|
     4|   # GET /employees or /employees.json
     5|   def index
     6|     @employees = Employee.all
=>   7|     debugger
     8|   end
     9|
    10|   # GET /employees/1 or /employees/1.json
    11|   def show
=>#0	EmployeesController#index at /opt/console-app/app/controllers/employees_controller.rb:7
  #1	ActionController::BasicImplicitRender#send_action(method="index", args=[]) at /usr/local/bundle/gems/actionpack-8.1.2/lib/action_controller/metal/basic_implicit_render.rb:8
  # and 78 frames (use `bt' command for all frames)
(ruby) @employees
[]
(rdbg) p headers    # print the headers
=> {"x-frame-options" => "SAMEORIGIN", "x-xss-protection" => "0", "x-content-type-options" => "nosniff", "x-permitted-cross-domain-policies" => "none", "referrer-policy" => "strict-origin-when-cross-origin"}
```

When you're finished, type `quit` and then `y` to exit the debugging session.

### Go

Let's set up our environment and use `GDB` to debug our code.

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
docker-compose build
docker-compose docker-compose run go bash
```

Now that we're logged into the container, let's run the app in debug mode. We'll start the app with `dlv debug` which will start the app in debug mode. We can then set a breakpoint, step through the code and inspect our variables. The `example.com/employees` is the package name we used in the Dockerfile above.

```golang
dlv debug example.com/employees
Type 'help' for list of commands.

(dlv) break main.main
Breakpoint 1 set at 0x938c20 for main.main() ./employees.go:45
(dlv) continue
> [Breakpoint 1] main.main() ./employees.go:45 (hits goroutine(1):1 total:1) (PC: 0x938c20)
    40:
    41:	  w.Header().Set("Content-Type", "application/json")
    42:	  w.Write(js)
    43:	}
    44:
=>  45:	func main() {
    46:	  db, err := gorm.Open(sqlite.Open("employees.db"), &gorm.Config{})
    47:	  if err != nil {
    48:	    log.Fatal("Could not connect to database employees.db: ", err)
    49:	  }
    50:

(dlv) next
> main.main() ./employees.go:46 (PC: 0x938c30)
    41:	  w.Header().Set("Content-Type", "application/json")
    42:	  w.Write(js)
    43:	}
    44:
    45:	func main() {
=>  46:	  db, err := gorm.Open(sqlite.Open("employees.db"), &gorm.Config{})
    47:	  if err != nil {
    48:	    log.Fatal("Could not connect to database employees.db: ", err)
    49:	  }
    50:
    51:	  db.AutoMigrate(&Employee{})

(dlv) next
> main.main() ./employees.go:47 (PC: 0x938d78)
    42:	  w.Write(js)
    43:	}
    44:
    45:	func main() {
    46:	  db, err := gorm.Open(sqlite.Open("employees.db"), &gorm.Config{})
=>  47:	  if err != nil {
    48:	    log.Fatal("Could not connect to database employees.db: ", err)
    49:	  }
    50:
    51:	  db.AutoMigrate(&Employee{})
    52:

(dlv) print err
error nil

(dlv) print db
Sending output to pager...
("*gorm.io/gorm.DB")(0x400022ade0)
*gorm.io/gorm.DB {
	Config: *gorm.io/gorm.Config {
		SkipDefaultTransaction: false,
		DefaultTransactionTimeout: 0,
		DefaultContextTimeout: 0,
...

(dlv) exit
```

## References

* https://github.com/go-delve/delve/tree/master/Documentation/cli
* https://guides.rubyonrails.org/debugging_rails_applications.html#debugging-with-the-debug-gem

## Wrap Up

In this chapter we've demonstrated how to start an app with a debugger enabled. While stepping through the code, we can inspect variables and monitor the state of our app to squash any bugs. In the next chapter, we'll dive into comparing unit testing between languages.

[Next >>](100-chapter-09.md)

