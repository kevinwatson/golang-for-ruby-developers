### Chapter 11 - Tasks

> 

## Introduction

Ruby on Rails ships with a tool named 'Rake' (short for Ruby Make), which Ruby's version of the `make` command in Unix. While the `make` command is used to automate the process of building and maintaining software, the Ruby version can be used for similar tasks including managing the data in your Ruby on Rails app's database.

## Examples

### Ruby

Ruby on Rails has a number of rake tasks that are included out of the box. Running `rake -T` will show the full list of built-in and custom tasks. The built-in rake tasks can be used to migrate the schema (`rake db:migrate`) and show the database version (`rake db:version`). Converting custom scripts that you run in Rails console will save you time because you can run the same code with a single command from the command line. For example, if I need to generate a report of employee records in my database, I might have a script that looks like this:

```ruby
rails console

> puts "id,first_name,last_name"
> Employee.all.each do |employee|
>   puts "#{employee.id},#{employee.first_name},#{employee.last_name}"
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
    puts "id,guid,first_name,last_name"
    Employee.all.each do |employee|
      puts "#{employee.id},#{employee.guid},#{employee.first_name},#{employee.last_name}"
    end
  end
end
```

After saving the `employee_report.rake` file, we can now run the code from the command line.

```bash
rake employee_report:print
```

We should see the following output.

```bash
id,guid,first_name,last_name
1,abcd,George,Jetson
```

Each time we need to run the report, it will save us time because we no longer need to run `rails console`, wait for it to load, and then copy and paste into the console. One line either copy and pasted or typed into the terminal will run the code in the task.

### Go


## References

* https://en.wikipedia.org/wiki/Make_(software)
* https://guides.rubyonrails.org/command_line.html#custom-rake-tasks

## Wrap Up

[Next >>](130-chapter-12.md)

