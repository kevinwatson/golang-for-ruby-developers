### Chapter 3 - Models, Views, Controllers (MVC)

> 

## Introduction

The model-view-controller software pattern is a well established pattern for building web applications. It provides a separation of concerns, allowing a model to represent an object, views to render the object to the user, and controllers to manage access to the object. The model doesn't know anything about how the object will be rendered in the views. Views can display a single or multiple objects.Controllers decide what view to render and whether the user has access to read, modify or delete the object.

Ruby on Rails encourages the use of fat models and lean controllers. Most of the business logic, including validation, associations and custom methods should be added to a model. Controllers should be lean, containing only necessary logic for handling user requests.

## Models

Most Ruby on Rails applications are designed to interact with a database. The model is responsible for persisting and loading the data. It may contain validations that ensure data consistency, associations between entities which provide access methods based on parent-child relationships (for example, loading a user's orders when the user is retrieved from the database) and callbacks which run code when an object's state changes (for example, deleting a user's profile when the user is deleted).

## Views

Views are the presentation layer. They convert the data into something that the user can consume, such as html in a browser. Along with the html that the user downloads and interacts with, the view layer can contain other assets like images and JavaScript.

## Controllers

Controllers act as the intermediary between models and views. They're responsible for responding to a user's request to view, update or delete a resource. They can also check whether the user is authorized to access or modify a particular resource.

## Comparing Languages

The Ruby language, by itself, was not web friendly. Along came the Ruby on Rails framework. Convention over configuration. Software patterns like MVC. A developer, using the Ruby on Rails framework, is productive and can create a full-featured web app in a minimal amount of time.

The Go language, by itself, provides tooling that will allow you to start a simple http server. If you want to write a full-featured MVC app in Go, you can. Import a few packages to provide the functionality you need such as an ORM and view rendering, and you'll be on your way to building a web app that's performant and easy to maintain.

## References

* https://guides.rubyonrails.org
* https://www.calhoun.io/using-mvc-to-structure-go-web-applications

## Wrap Up

MVC is a software pattern that organizes the code in a project to separate concerns. The Ruby on Rails framework is designed around this pattern. Go provides a number of built-in packages that, with a little bit of work, can provide some of the same functionality.

If your goal is developer productivity, Ruby on Rails provides features that make it easy to get up and running. If your platform grows and later you need performance, Go can provide those optimizations.

[Next >>](050-chapter-04.md)
