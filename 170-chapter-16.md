### Chapter 16 - Dependency Management

## Introduction

Dependent libraries or 'dependencies' are a way to quickly add features to our app. For example, Ruby on Rails is a collection of libraries (aka 'gems') that can be added either individually (for example, if we only need the features provided by ActiveRecord, we can include just that library and use its features) or as a collection if you want to use all of the framework's features.

## Examples

### Ruby

Ruby's libraries are called 'gems' which is a clever term based on the name of the Ruby language. Gems, or gemstones, are defined as a precious piece of mineral crystal, organic matter or rock that has been cut and polished. There might be gems embedded in the rings on your hands or other jewelry. In Ruby terminology a 'gem' is a valuable and polished add-on that enhances our core app.

Ruby includes a tool named Bundler. Bundler uses two files to manage an app's dependencies or gems. These files are `Gemfile` and `Gemfile.lock`. The `Gemfile` defines which top level libraries the app depends on in order to run correctly. The `Gemfile.lock` file is maintained by the Bundler app and defines the specific versions and dependency tree. Running `bundle update` will download new versions of gems and keep old versions on your system. Prefixing `rails server` with `bundle exec` (e.g. `bundle exec rails server`) when starting a Ruby on Rails app will read the Gemfile and Gemfile.lock files and run the app with only the versions defined in those files.

### Go

## References

* https://bundler.io

## Wrap Up

