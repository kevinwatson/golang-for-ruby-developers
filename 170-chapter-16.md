### Chapter 16 - Dependency Management

## Introduction

Dependent libraries or 'dependencies' are a way to quickly add features to our app. For example, Ruby on Rails is a collection of libraries (aka 'gems') that can be added either individually (for example, if we only need the features provided by ActiveRecord, we can include just that library and use its features) or as a collection if you want to use all of the framework's features.

As libraries are written and modified, they can be versioned. Over time, as bugs are fixed and features are added, incrementing the version allows consumers of the library to identify and target a specific version for their app and upgrade when they're ready. Without version management, our app would behave unpredictably each time it was built and deployed. One common way to manage version an app is with Semantic versioning (https://semver.org). This versioning scheme uses the format X.Y.Z, where each is an incrementing number (e.g. 1.4.6) and X is the major version, Y is the minor version and Z is the patch version. Bug fixes that don't affect the public API of the library increment the patch number, backward compatible changes to the API increment the minor version number, and backward incompatible changes increment the major version number. With this versioning scheme, using the version number alone, the code maintainer can understand the risk of upgrading a dependency without reviewing all of the changes. For example, if we're currently using library A version 1.5.12 and 1.5.13 is available, we can be fairly confident that consuming the next version won't break our app. As the saying goes, we should always 'trust but verify' by running our unit and other tests after upgrading any dependencies.

## Examples

### Ruby

Ruby's libraries are called 'gems' which is a clever term based on the name of the Ruby language. Gems, or gemstones, are defined as a precious piece of mineral crystal, organic matter or rock that has been cut and polished. There might be gems embedded in the rings on your hands or other jewelry. In Ruby terminology a 'gem' is a valuable and polished add-on that enhances our core app.

Ruby includes a tool named Bundler. Bundler uses two files to manage an app's dependencies or gems. These files are `Gemfile` and `Gemfile.lock`. The `Gemfile` defines which top level libraries the app depends on in order to run correctly. The `Gemfile.lock` file is maintained by the Bundler app and defines the specific versions and dependency tree. Running `bundle update` will download new versions of gems and keep old versions on your system. Prefixing `rails server` with `bundle exec` (e.g. `bundle exec rails server`) when starting a Ruby on Rails app will read the Gemfile and Gemfile.lock files and run the app with only the versions defined in those files.

### Go

## References

* https://bundler.io
* https://semver.org

## Wrap Up

