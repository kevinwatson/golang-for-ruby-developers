### Chapter 4 - Object Relational Mapping (ORMs)

> 

## Introduction

Object relational mapping is generally a library that converts code into structured query language (SQL) code that allows you to interact with your database. Ruby on Rails supports multiple databases and can convert the same Ruby code to the SQL that is specific to the database you're using. This allows for developer productivity, since they don't have to write SQL code and mix that in with their Ruby code.

## Comparing Languages

### Ruby on Rails

Ruby on Rails provides the ActiveRecord gem which provides a number of features. These features include:

* Model base classes
* Primary and foreign key support with related one-to-one, one-to-many and many-to-many relationships
* Database change support (aka migrations)
* Generic methods like `save` which abstract away whether to generate a SQL `insert` or `update`
* Validations which can be configured to run before inserting or updating rows
* Callbacks which can run before or after certain events

### Go

## Example

## References

* https://guides.rubyonrails.org/active_record_basics.html

## Wrap Up

