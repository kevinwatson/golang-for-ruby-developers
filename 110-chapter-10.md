### Chapter 10 - Logging

> 

## Introduction

Logging is one way to observe the inner workings of an application as it runs. It allows you to monitor whether your app is behaving as expected, or display valuable information when something goes wrong. In the last chapter we covered debugging which allows you to pause a program and step through code line by line. Logging is similar, except that you can let a program or algorithm run to completion at its normal speed and then observe the various states of the program in real time or afterwards.

## Examples

### Ruby

Ruby on Rails supports multiple log levels. This allows you to write logging details in different modes, depending on how the app is configured. For example, you can add write `debug` lines which only print when the app is configured in `debug` mode, or when you're running the application locally. These debug lines can be verbose because they add no overhead to the app when it's running in your production environment, as long as you use the `logger.debug {"my string"}` block syntax instead of the normal `logger.debug "my string"` syntax.

Let's look at a few examples. By default, Ruby on Rails generates and writes its logs to a file in the `log` folder. For each environment (the defaults are `development`, `test` and `production`), a separate file is written, for example `log/development.log` and `log/production.log`. If you're running your apps in a container, you'll want it to write logs to standard output (`stdout`) so that it can be captured and aggregated.

Configure logs to write to stdout

```ruby
# config/environment.rb
config.logger = Logger.new(STDOUT)
```

Rails has a few built-in logging classes, such as the default [ActiveSupport::Logger](https://api.rubyonrails.org/classes/ActiveSupport/Logger.html) class, the [ActiveSupport::TaggedLogging](https://api.rubyonrails.org/classes/ActiveSupport/TaggedLogging.html) class, as well as the [ActiveSupport::BroadcastLogger](https://api.rubyonrails.org/classes/ActiveSupport/BroadcastLogger.html) class. You can also use a third-party class or create your own.

#### Log Levels

As mentioned previously, Ruby on Rails runs in a mode such as development which can be configured differently than the test or production environments. Over time as our app has grown, we may not want to bog down our production environment with unnecessary debug messages. In our `config/environments/production.rb` file we can set the `config.log_level` to `:warn` or a higher level so that it generates only the information that would be useful if there was an error. The log levels are as follows:

|Level|Description|
|---|---|
|debug|The most verbose level, normally used for writing low level debugging lines|
|info|The second level, informing us of normal app behavior|
|warn|The third level, informing us of unexpected behavior|
|error|The fourth level, informing us of a failed operation|
|fatal|The fifth level, informing us of an process exit|

In our app we can target which logger level and above where the messages should be written. For example, if we use the following line

```ruby
Rails.logger.info "The employee was saved successfully"
```

When `config.log_level = :info` is in our settings, the message will be written. If the app's environment were configured at a higher level like `config.log_level = :warn`, the line above wouldn't show up in our logs.

### Go

## References

* https://www.dash0.com/knowledge/rails-log-levels
* https://guides.rubyonrails.org/debugging_rails_applications.html#what-is-the-logger-questionmark
* https://guides.rubyonrails.org/debugging_rails_applications.html#impact-of-logs-on-performance

## Wrap Up


[Next >>](120-chapter-11.md)

