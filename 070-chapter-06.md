### Chapter 6 - Controllers

> 

## Introduction

Controllers in the MVC pattern manage the passing of data between the models and the views. They receive user input or other client data that can be passed to the model. Controllers manage other app requirements such as security and routing.

## Comparing Languages

### Ruby on Rails

In Ruby on Rails, controllers are created to match the resource whenever a generator is used. By default, generators will create a number of routes, such as `index`, `show`, `create`, `update` and `destroy` and add the appropriate model related code to list or retrieve and modify an instance of a model. This out of the box functionality can be customized as needed by modifying the controller class and the `config/routes.rb` file.

### Go

The Go frameworks mentioned in the previous chapter provide routing and controller matching which can be defined in the `main` function, or separated into other packages as needed. By design, a Go app can be as simple (a single package or code file) or as complex as needed. Defining routes and handler functions (aka controllers) are defined as needed, making a Go JSON API easy to implement and maintain.

## Examples

## References

* https://github.com/gorilla/mux
* https://guides.rubyonrails.org/action_controller_overview.html

## Wrap Up

[Next >>](080-chapter-07.md)

