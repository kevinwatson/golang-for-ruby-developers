### Chapter 5 - Views

> 

## Introduction

Views in the MVC pattern are the presentation layer of an application that the user or client interacts with directly. It converts the data received from the controller and model into a format that can be consumed. It can be in various formats, for example HTML, XML or JSON. HTML is common for building a website that will be rendered in a browser, while XML and JSON are commonly used for views that will be read by code to either render HTML to a user in single page application (SPA) or machine to machine interactions.

For simplicty's sake, we'll compare JSON rendering between the two languages.

## Comparing Languages

### Ruby on Rails

Ruby on Rails provides an `api` mode which runs a leaner version of the framework. Here are a few of these API features:

* Generators will skip helpers, views and assets
* Middleware that is needed for rending browser assets (html, cookies) is not included
* Familiar Rails features
  * Routing
  * Security
  * Request parameter parsing
  * Logging
  * Caching
  * Controllers
  * Models
  * Database schema version management

As we can see, Rails provides a mode that is leaner on the presentation layer but still includes full controller and model functionality. It's a quick way to spin up a JSON API as either a stand-alone application or as the backend of a website.

### Go


## Examples

### Go

## References

* https://guides.rubyonrails.org/api_app.html

## Wrap Up

[Next >>](070-chapter-06.md)

