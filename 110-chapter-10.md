### Chapter 10 - Logging

> 

## Introduction

Logging is one way to observe the inner workings of an application as it runs. It allows you to monitor whether your app is behaving as expected, or display valuable information when something goes wrong. In the last chapter we covered debugging which allows you to pause a program and step through code line by line. Logging is similar, except that you can let a program or algorithm run to completion at its normal speed and then observe the various states of the program in real time or afterwards.

## Examples

### Ruby

Ruby on Rails supports multiple log levels. This allows you to write logging details in different modes, depending on how the app is configured. For example, you can add write `debug` lines which only print when the app is configured in `debug` mode, or when you're running the application locally. These debug lines can be verbose because they add no overhead to the app when it's running in your production environment, as long as you use the `logger.debug {"my string"}` block syntax instead of the normal `logger.debug "my string"` syntax.

### Go

## References

* https://guides.rubyonrails.org/debugging_rails_applications.html
* https://guides.rubyonrails.org/debugging_rails_applications.html#impact-of-logs-on-performance

## Wrap Up


[Next >>](120-chapter-11.md)

