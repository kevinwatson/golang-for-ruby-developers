### Chapter 10 - Logging

> 

## Introduction

Logging is one way to observe the inner workings of an application as it runs. It allows you to monitor whether your app is behaving as expected, or display valuable information when something goes wrong. In the last chapter we covered debugging which allows you to pause a program and step through code line by line. Logging is similar, except that you can let a program or algorithm run to completion at its normal speed and then observe the various states of the program in real time or afterwards.

## Examples

### Ruby

Ruby on Rails supports multiple log levels. This allows you to write logging details in different modes, depending on how the app is configured. For example, you can add write `debug` lines which only print when the app is configured in `debug` mode, or when you're running the application locally. These debug lines can be verbose because they add no overhead to the app when it's running in your production environment, as long as you use the `logger.debug {"my string"}` block syntax instead of the normal `logger.debug "my string"` syntax.

Let's look at a few examples. By default, Ruby on Rails generates and writes its logs to a file in the `log` folder. For each environment (the defaults are `development`, `test` and `production`), a separate file is written, for example `log/development.log` and `log/production.log`. If you're running your apps in a container, you'll want it to write logs to standard output (`stdout`) so that they can be captured and aggregated.

Configure logs to write to stdout:

```ruby
# config/environment.rb
config.logger = Logger.new(STDOUT)
```

Rails has a few built-in logging classes, such as the default [ActiveSupport::Logger](https://api.rubyonrails.org/classes/ActiveSupport/Logger.html) class, the [ActiveSupport::TaggedLogging](https://api.rubyonrails.org/classes/ActiveSupport/TaggedLogging.html) class, as well as the [ActiveSupport::BroadcastLogger](https://api.rubyonrails.org/classes/ActiveSupport/BroadcastLogger.html) class. You can also use a third-party class or create your own.

#### Log Levels

As mentioned previously, Ruby on Rails always runs in a mode (e.g. `development`) which can be configured differently than the `test` or `production` environments. In a high-volume production environment, we don't want to fill our logs with unnecessary `debug` or even `info` messages. In our `config/environments/production.rb` file we can set the `config.log_level = :warn` or a higher level so that it generates only the information that would be useful if there was an error. The log levels are as follows:

|Level|Description|
|---|---|
|fatal|The fifth level, informing us of a process exit|
|error|The fourth level, informing us of a failed operation|
|warn|The third level, informing us of unexpected behavior|
|info|The second level, informing us of normal app behavior|
|debug|The first and most verbose level, normally used for writing low level debugging lines|

For each level in the above table, if our app is configured at that level, the level above it are automatically included. For example, if we configured our app to write 'info' messages and our app executes the following line of code, we'll see it in our logs.

```ruby
Rails.logger.warn "The method 'pay_all_bonuses' is deprecated and will be removed in the next version"
```

This would result in the following line in our log file:

```
W, [2026-05-07T11:00:00.000000] WARN -- : The method 'pay_all_bonuses' is deprecated and will be removed in the next version
```

### Go

Similar to Ruby, Go apps can be configured to write logging information with various logging levels. The `log/slog` (structured logging) package provides these options. In the default Ruby `Logger` class, you're expected to build your string (e.g. `Rails.logger.debug { "Employee first_name = '#{employee.first_name}' last_name = '#{employee.last_name} could not be saved" }`), whereas with the `log/slog` package's logging function signature allows you to pass any type as the last argument. For example:

```golang
slog.Warn("Employee could not be saved", "first_name", employee.FirstName, "last_name", employee.LastName)
```

Would result in this line:

```
2026/05/07 11:00:00 WARN Employee could not be saved first_name=George last_name=Jetson
```

#### Log Levels

|Level|Description|
|---|---|
|Error|The fourth level, informing us of a process exit|
|Warn|The third level, informing us that something unexpected happened, but the app recovered|
|Info|The second and default level, informing us that things are working as expected|
|Debug|The first and most verbose level, normally used for writing low level debugging lines|

By default, the `Info` level and above will be written to stdout. We can use the following code to set it to `Debug`:

```golang
slog.SetLogLoggerLevel(slog.LevelDebug)
```

We can also customize the output with the `SetFlags` function:

```golang
log.SetFlags(log.Ldate | log.Lmicroseconds | log.Llongfile)
```

Logging after the line above will have this output:

```bash
2026/05/07 11:00:00.000000 /opt/app/employee.go:201: WARN Employee could not be saved first_name=George last_name=Jetson
```

## References

* https://www.dash0.com/knowledge/rails-log-levels
* https://guides.rubyonrails.org/debugging_rails_applications.html#what-is-the-logger-questionmark
* https://guides.rubyonrails.org/debugging_rails_applications.html#impact-of-logs-on-performance
* https://pkg.go.dev/log/slog

## Wrap Up

As we can see, both languages provide you With various options you can use to capture the data you're looking for and later filter the log files by log level or keywords to find what you need to debug or observe the behavior of your application. In the next chapter, we'll discuss writing one-off tasks that can be executed from the command line.

[Next >>](120-chapter-11.md)

