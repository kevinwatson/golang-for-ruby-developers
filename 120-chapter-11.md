### Chapter 11 - Tasks

> 

## Introduction

Ruby on Rails ships with a tool named 'Rake' (short for Ruby Make), which Ruby's version of the `make` command in Unix. While the `make` command is used to automate the process of building and maintaining software, the Ruby version can be used for similar tasks including managing the data in your Ruby on Rails app's database.

## Examples

### Ruby

Ruby on Rails has a number of rake tasks that are included out of the box. Running `rake -T` will show the full list of built-in and custom tasks. The built-in rake tasks can be used to migrate the schema (`rake db:migrate`) and show the database version (`rake db:version`). Converting custom scripts that you run in Rails console will save you time because you can run the same code with a single command from the command line. For example, if I need to generate a report of employee records in my database, I might have a script that looks like this:

```ruby
rails console

> puts "guid,first_name,last_name"
> Employee.all.each do |employee|
>   puts "#{employee.guid},#{employee.first_name},#{employee.last_name}"
> end
```

If this report is something I need on a regular basis but the script hasn't changed for a while, I could write the following rake task once and run it with a single command. Using the following steps we'll write a rake task to hard-code our logic.

```bash
bin/rails generate task employee_report print
```

This will generate an `lib/tasks/employee_report.rake` file. We can then edit the file and add our code.

```ruby
# lib/tasks/employee_report.rake

namespace :employee_report do
  desc "prints a list of employees"
  task print: :environment do

    # print the output in CSV format
    puts "guid,first_name,last_name"
    Employee.all.each do |employee|
      puts "#{employee.guid},#{employee.first_name},#{employee.last_name}"
    end
  end
end
```

After saving the `employee_report.rake` file, we can now run the code from the command line.

```bash
rake employee_report:print
```

We should see the following output.

```csv
guid,first_name,last_name
abcd,George,Jetson
```

Each time we need to run the report, it will save us time because we no longer need to run `rails console`, wait for it to load, and then copy and paste into the console. One line either copy and pasted or typed into the terminal will run the code in the task.

### Go

In chapter 7 (the 'Console' chapter), we discussed the need to create a file for our code because Go is a compiled language. This is fitting for this chapter on Rake tasks, because the script that we want to reuse can be captured as a 'task' for our future use. For example, we can create a `tasks` folder in our project and add script that we'll use there.

Let's recreate the Ruby logic above in Go. We'll import the `models` package and insert a record to generate data for our report.

Our Dockerfile

```dockerfile
# Dockerfile
FROM golang:1.25

WORKDIR /opt/app
RUN go mod init example.com/my-app
COPY . .
RUN go get ./...
```

Our Docker Compose file

```yaml
# docker-compose.yml

docker-compose.yml
services:
  go:
    build:
      context: .
    stdin_open: true
    tty: true
```

Our employee model (in the `models` directory)

```golang
// models/employee.go

package models

import (
  "gorm.io/gorm"
)

type Employee struct {
  gorm.Model
  Id        int
  Guid      string
  FirstName string
  LastName  string
}
```

Our employee print task (in the `tasks/employee_report` directory). Note that because we want to run this file on its own, we named the package `main` and the only function in the package is the `main` function.

```golang
// tasks/employee_report/print.go

package main

import (
    "context"
    "example.com/my-app/models"
    "fmt"
    "log"
    "gorm.io/driver/sqlite"
    "gorm.io/gorm"
)

func main() {
    db, err := gorm.Open(sqlite.Open("employees.db"), &gorm.Config{})
    if err != nil {
      log.Fatal("Could not connect to database employees.db: ", err)
    }

    ctx := context.Background()

    // query our employee table
    records, err := gorm.G[models.Employee](db).Find(ctx)
    if err != nil {
      fmt.Printf("Could not connect to database %s\n", err)
    }

    // print the output in CSV format
    fmt.Println("guid,first_name,last_name")
    for _, e := range records {
        fmt.Printf("%s,%s,%s\n", e.Guid, e.FirstName, e.LastName)
    }
}
```

We can run the task with this command:

```bash
docker-compose run go
go run tasks/employee_report/print.go
```

To produce this output

```csv
guid,first_name,last_name
abcd,George,Jetson
```

## References

* https://en.wikipedia.org/wiki/Make_(software)
* https://gorm.io/docs
* https://guides.rubyonrails.org/command_line.html#custom-rake-tasks

## Wrap Up

No matter which language your app is written in, writing tasks that encapsulate some bit of repetitive logic can save you and your team time down the road. Capturing this logic and being able to run it from the command line also gives you the ability to run it on a timer (for example in a cron job). Writing a task in Go is similar to writing a task in Ruby, you can access the models and other code as needed to 'accomplish' your task.

[Next >>](130-chapter-12.md)

