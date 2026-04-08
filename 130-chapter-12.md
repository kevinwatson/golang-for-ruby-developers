### Chapter 12 - Code Generators

> 

## Introduction

Code generators are command-line tools that generate files and folders to get you up and running quickly. Because Rails is built on the MVC framework, it includes a number of built-in generators that can generate views, controllers, models or `scaffold` which will generate all of the above. You can also create custom generators for your own needs.

Because Go is a typed language, it requires that we explicity define things like database models with structs (Ruby on Rails determines this at runtime by querying its database). Code generators like the built-in `go generate` command can be used to auto-generate some of this repetitive code.

Nowadays with generative AI you may be asking 'do we need generators? AI can create the code for us!' This may be true. AI can generate some or all of the code you need to add a new feature for common use cases (like adding a model file). But, one day we may have custom generators that the AI is not aware of that writes custom code needed to run our app. We can add the step of running our generator for specific cases to our AI agent's context. To guarantee consistent results, the agent can then be instructed to run the generator command at the appropriate time when making code changes.

## Examples

### Ruby

Rails includes a `scaffold` option that will create a full set of files for a resource, including views, a controller, a model and a database migration file. For example, to add a `benefits` table to our employee database with a model, controller and views we can run this command:

```ruby
bin/rails generate scaffold benefit name:string benefit_type:integer activated_at:datetime

      invoke  active_record
      create    db/migrate/20260408080055_create_benefits.rb
      create    app/models/benefit.rb
      invoke    test_unit
      create      test/models/benefit_test.rb
      create      test/fixtures/benefits.yml
      invoke  resource_route
       route    resources :benefits
      invoke  scaffold_controller
      create    app/controllers/benefits_controller.rb
      invoke    erb
      create      app/views/benefits
      create      app/views/benefits/index.html.erb
      create      app/views/benefits/edit.html.erb
      create      app/views/benefits/show.html.erb
      create      app/views/benefits/new.html.erb
      create      app/views/benefits/_form.html.erb
      create      app/views/benefits/_benefit.html.erb
      invoke    resource_route
      invoke    test_unit
      create      test/controllers/benefits_controller_test.rb
      invoke    helper
      create      app/helpers/benefits_helper.rb
      invoke      test_unit
      invoke    jbuilder
      create      app/views/benefits/index.json.jbuilder
      create      app/views/benefits/show.json.jbuilder
      create      app/views/benefits/_benefit.json.jbuilder
```

We can then customize the files to finish our new feature.

### Go


## References

* https://go.dev/blog/generate
* https://go.googlesource.com/proposal/+/refs/heads/master/design/go-generate.md
* https://guides.rubyonrails.org/command_line.html#generating-code

## Wrap Up


[Next >>](140-chapter-13.md)

