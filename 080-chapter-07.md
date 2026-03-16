### Chapter 7 - Console

> 

## Introduction

Rails console is a commonly used REPL (read eval print loop) feature. It builds on IRB which is a built in Ruby tool for running code in an interactive shell environment. With the dynamic nature of Ruby, you can create new classes or access existing classes and call methods on those classes and objects. In the shell, you also have access to the rest of the environment that's available at runtime, so you can access resources to help you test out ideas while in development or access databases in a test or production environment.

Go is a compiled language, which makes running a REPL more difficult. While not a built-in feature, that hasn't stopped the Go community from building their own. With a few limitations, the third party gomacro and yaegi projects work similar to Rails console.

## Comparing Languages

## Examples

### Ruby on Rails

```bash
rails console
```

```ruby
```

### Go

#### Gomacro

```bash
gomacro
```

```go
```

## References

* https://github.com/cosmos72/gomacro
* https://github.com/traefik/yaegi
* https://guides.rubyonrails.org/command_line.html#bin-rails-console

## Wrap Up


[Next >>](090-chapter-08.md)
