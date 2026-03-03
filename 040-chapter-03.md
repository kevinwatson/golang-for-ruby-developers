### Chapter 3 - Models, Views, Controllers (MVC)

> 

## Introduction

The model-view-controller software pattern is a well established pattern for building web applications. It provides a separation of concerns, allowing a model to represent an object, views to render the object to the user, and controllers to manage access to the object. The model doesn't know anything about how the object will be rendered in the views. Views can display a single or multiple objects.Controllers decide what view to render and whether the user has access to read, modify or delete the object.

Ruby on Rails encourages the use of fat models and lean controllers. Most of the business logic, including validation, associations and custom methods should be added to a model. Controllers should be lean, containing only necessary logic for handling user requests.

## Models

Most Ruby on Rails applications are designed to interact with a database. The model is responsible for persisting and loading the data. It may contain validations that ensure data consistency, associations between entities which provide access methods based on parent-child relationships (for example, loading a user's orders when the user is retrieved from the database) and callbacks which run code when an object's state changes (for example, deleting a user's profile when the user is deleted).

## Views

## Controllers

## Comparing Languages

## Example

## References


## Wrap Up


[Next >>](050-chapter-04.md)
