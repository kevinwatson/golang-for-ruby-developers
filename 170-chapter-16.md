### Chapter 16 - Dependency Management

## Introduction

Dependent libraries or 'dependencies' are a way to quickly add features to our app. For example, Ruby on Rails is a collection of libraries (aka 'gems') that can be added either individually (for example, if we only need the features provided by ActiveRecord, we can include just that library and use its features) or as a collection if you want to use all of the framework's features.

As libraries are written and modified, they can be versioned. Over time, as bugs are fixed and features are added, incrementing the version allows consumers of the library to identify and target a specific version for their app and upgrade when they're ready. Without version management, our app would behave unpredictably each time it was built and deployed. One common way to manage version an app is with Semantic versioning (https://semver.org). This versioning scheme uses the format X.Y.Z, where each is an incrementing number (e.g. 1.4.6) and X is the major version, Y is the minor version and Z is the patch version. Bug fixes that don't affect the public API of the library increment the patch number, backward compatible changes to the API increment the minor version number, and backward incompatible changes increment the major version number. With this versioning scheme, using the version number alone, the code maintainer can understand the risk of upgrading a dependency without reviewing all of the changes. For example, if we're currently using library A version 1.5.12 and 1.5.13 is available, we can be fairly confident that consuming the next version won't break our app. As the saying goes, we should always 'trust but verify' by running our unit and other tests after upgrading any dependencies.

## Examples

### Ruby

Ruby's libraries are called 'gems' which is a clever term based on the name of the Ruby language. Gems, or gemstones, are defined as a precious piece of mineral crystal, organic matter or rock that has been cut and polished. There might be gems embedded in the rings on your hands or other jewelry. In Ruby terminology a 'gem' is a valuable and polished add-on that enhances our core app.

Ruby includes a tool named Bundler. Bundler uses two files to manage an app's dependencies or gems. These files are `Gemfile` and `Gemfile.lock`. The `Gemfile` defines which top level libraries the app depends on in order to run correctly. The `Gemfile.lock` file is maintained by the Bundler app and defines the specific versions and dependency tree. Running `bundle install` will scan the `Gemfile` file and download and install any missing libraries on your system. Prefixing `rails server` with `bundle exec` (e.g. `bundle exec rails server`) when starting a Ruby on Rails app will read the Gemfile and Gemfile.lock files and run the app with only the versions defined in those files.

Let's spin up a new container to create a fresh environment where we can experiment with the bundler gem. We'll create a new Gemfile and add a couple of gems via the command line so we can run some code in the console.

```bash
docker run --rm -it ruby:latest bash
```

Install the bundler gem

```bash
# gem install bundler
Fetching bundler-4.0.15.gem
Successfully installed bundler-4.0.15
1 gem installed
```

Create a new folder and add a Gemfile with the `bundle init` command

```bash
# mkdir opt/bundle_test

# cd opt/bundle_test/
# bundle init
Writing new Gemfile to /opt/bundle_test/Gemfile
```

Inspect the file

```bash
# cat Gemfile
# frozen_string_literal: true

source "https://rubygems.org"
```

Add two gems, one to pretty print objects and one to manage JSON objects and show where they're installed

```bash
# bundle add amazing_print
Fetching gem metadata from https://rubygems.org/.
Resolving dependencies...
Fetching gem metadata from https://rubygems.org/.
Fetching amazing_print 2.0.0
Installing amazing_print 2.0.0

# bundle show amazing_print
/usr/local/bundle/gems/amazing_print-2.0.0

# bundle add json
Fetching gem metadata from https://rubygems.org/.
Resolving dependencies...
Fetching gem metadata from https://rubygems.org/.
Fetching json 2.20.0
Installing json 2.20.0 with native extensions

# bundle show json
/usr/local/bundle/gems/json-2.20.0
```

Inspect the modified Gemfile

```ruby
# cat Gemfile
# frozen_string_literal: true

source "https://rubygems.org"

# gem "rails"

gem "amazing_print", "~> 2.0"

gem "json", "~> 2.20"
```

Inspect the Gemfile.lock

```ruby
# cat Gemfile.lock
GEM
  remote: https://rubygems.org/
  specs:
    amazing_print (2.0.0)
    json (2.20.0)

PLATFORMS
  aarch64-linux
  ruby

DEPENDENCIES
  amazing_print (~> 2.0)
  json (~> 2.20)

CHECKSUMS
  amazing_print (2.0.0) sha256=2e36aba46ac78d37ed27ca0e2056afe3583183bb5c64f157c246b267355e5d6a
  bundler (4.0.15) sha256=a4ceb882fe94a0e0ac63cd0813932bbfd631a14e5ac0b7975189b19a4d28d9e7
  json (2.20.0) sha256=9362bc6e55a952b056abf9167cf053358181c904cb70cd6eee0808ea830fc32b

BUNDLED WITH
  4.0.15
```

Start a console in the context of the current bundle

```ruby
# bundle console
irb(main):001> ap JSON.parse('{"a": 1, "b": 2, "c": 3}')
{
    "a" => 1,
    "b" => 2,
    "c" => 3
}
```

With a few bundler commands, we're able to add new gems and lock them to specific versions which ensures that our apps behaves the same wherever it's running.

### Go

Go has native support for dependency management with its `go mod` command line tool. A `go.mod` file is similar to Bundler's Gemfile and Gemfile.lock files. It contains dependency names, the repository domain name and path for each module where the code can be downloaded, and the version of the dependency. While the `go.mod` file can be edited by hand, the `go mod` commands can be used to manage the contents of the go.mod file.

The `go.sum` file is used to manage the checksums of the modules we depend on. The file is generated by the `go get` command when the first module is downloaded.

Let's run through a couple of examples.

#### Managing Dependencies Using go get package

First, we'll start a fresh container on the latest version of Go.

```bash
docker run --rm -it golang:latest sh
```

Now, let's look through the existing directories. We can see that they're empty.

```bash
# ls
bin  src

# ls src

# ls bin

```

Next, we'll create a new module with the `go mod init` command. We can see that it created the `go.mod` file and the `pkg` directory.

```bash
# go mod init dependencies/main
go: creating new go.mod: module dependencies/main
go: to add module requirements and sums:
	go mod tidy

# ls
bin  go.mod  pkg  src

# cat go.mod
module dependencies/main

go 1.26.4
```

Next, we'll pull a package and its dependencies. We'll inspect the modified `go.mod` file and the newly created `go.sum` file.

```bash
# go get gorm.io/gorm
go: warning: ignoring go.mod in $GOPATH /go
go: downloading gorm.io/gorm v1.31.2
go: downloading github.com/jinzhu/now v1.1.5
go: downloading github.com/jinzhu/inflection v1.0.0
go: downloading golang.org/x/text v0.20.0
go: added github.com/jinzhu/inflection v1.0.0
go: added github.com/jinzhu/now v1.1.5
go: added golang.org/x/text v0.20.0
go: added gorm.io/gorm v1.31.2

# cat go.mod
module dependencies/main

go 1.26.4

require (
	github.com/jinzhu/inflection v1.0.0 // indirect
	github.com/jinzhu/now v1.1.5 // indirect
	golang.org/x/text v0.20.0 // indirect
	gorm.io/gorm v1.31.2 // indirect
)

# cat go.sum
github.com/jinzhu/inflection v1.0.0 h1:K317FqzuhWc8YvSVlFMCCUb36O/S9MCKRDI7QkRKD/E=
github.com/jinzhu/inflection v1.0.0/go.mod h1:h+uFLlag+Qp1Va5pdKtLDYj+kHp5pxUVkryuEj+Srlc=
github.com/jinzhu/now v1.1.5 h1:/o9tlHleP7gOFmsnYNz3RGnqzefHA47wQpKrrdTIwXQ=
github.com/jinzhu/now v1.1.5/go.mod h1:d3SSVoowX0Lcu0IBviAWJpolVfI5UJVZZ7cO71lE/z8=
golang.org/x/text v0.20.0 h1:gK/Kv2otX8gz+wn7Rmb3vT96ZwuoxnQlY+HlJVj7Qug=
golang.org/x/text v0.20.0/go.mod h1:D4IsuqiFMhST5bX19pQ9ikHC2GsaKyk/oF+pn3ducp4=
gorm.io/gorm v1.31.2 h1:3o8FXNo9v9S858gil+3LlZA1LkCOzgb4g5BL64FgaCo=
gorm.io/gorm v1.31.2/go.mod h1:XyQVbO2k6YkOis7C2437jSit3SsDK72s7n7rsSHd+Gs=
```

Now let's check the folders where the code was downloaded. Notice that the folder names include the verison number.

```bash
# ls pkg/mod/github.com/jinzhu
inflection@v1.0.0  now@v1.1.5

# ls pkg/mod/github.com/jinzhu/inflection@v1.0.0
LICENSE  README.md  go.mod  inflections.go  inflections_test.go  wercker.yml

# ls pkg/mod/github.com/jinzhu/now@v1.1.5
Guardfile  License  README.md  go.mod  main.go	now.go	now_test.go  time.go
```

#### Managing Dependencies Using Imports

Another way to manage dependencies is by adding imports to the project files and running the command `go get ./...` to retrieve the dependencies. Let's try it with a source code file.

Copying a file from a previous chapter:

```golang
# cat src/main.go
package main

import (
    "context"
    "fmt"
    "log"
    "gorm.io/driver/sqlite"
    "gorm.io/gorm"
)

func main() {
    db, err := gorm.Open(sqlite.Open("employees.db"), &gorm.Config{})
    if err != nil {
      log.Fatal("Could not connect to database employees.db: ", err)
    }

    ctx := context.Background()

    // query our employee table
    records, err := gorm.G[models.Employee](db).Find(ctx)
    if err != nil {
      fmt.Printf("Could not connect to database %s\n", err)
    }

    // print the output in CSV format
    fmt.Println("guid,first_name,last_name")
    for _, e := range records {
        fmt.Printf("%s,%s,%s\n", e.Guid, e.FirstName, e.LastName)
    }
}
```

Run a Docker container to load the environment, mounting the main.go file in the `src` directory

```bash
docker run --rm -it -v $(pwd)/main.go:/go/src/main.go golang:latest sh
```

List the existing files

```bash
# ls
bin  src
```

Initialize the module

```bash
# go mod init example/main
go: creating new go.mod: module example/main
go: to add module requirements and sums:
	go mod tidy
```

List the files again

```bash
# ls
bin  go.mod  pkg  src
```

Show the contents of the new go.mod file

```bash
# cat go.mod
module example/main
```

Use `go get .` to retrieve modules from the packages

```bash
go 1.26.4
# go get ./...
go: warning: ignoring go.mod in $GOPATH /go
go: downloading gorm.io/driver/sqlite v1.6.0
go: downloading gorm.io/gorm v1.31.2
go: downloading github.com/jinzhu/now v1.1.5
go: downloading github.com/mattn/go-sqlite3 v1.14.22
go: downloading github.com/jinzhu/inflection v1.0.0
go: downloading golang.org/x/text v0.20.0
go: added github.com/jinzhu/inflection v1.0.0
go: added github.com/jinzhu/now v1.1.5
go: added github.com/mattn/go-sqlite3 v1.14.22
go: added golang.org/x/text v0.20.0
go: added gorm.io/driver/sqlite v1.6.0
go: added gorm.io/gorm v1.31.2
```

List the files again. Now there's a go.sum file.

```bash
# ls
bin  go.mod  go.sum  pkg  src
```

Show the contents of the go.sum file

```bash
# cat go.sum
github.com/jinzhu/inflection v1.0.0 h1:K317FqzuhWc8YvSVlFMCCUb36O/S9MCKRDI7QkRKD/E=
github.com/jinzhu/inflection v1.0.0/go.mod h1:h+uFLlag+Qp1Va5pdKtLDYj+kHp5pxUVkryuEj+Srlc=
github.com/jinzhu/now v1.1.5 h1:/o9tlHleP7gOFmsnYNz3RGnqzefHA47wQpKrrdTIwXQ=
github.com/jinzhu/now v1.1.5/go.mod h1:d3SSVoowX0Lcu0IBviAWJpolVfI5UJVZZ7cO71lE/z8=
github.com/mattn/go-sqlite3 v1.14.22 h1:2gZY6PC6kBnID23Tichd1K+Z0oS6nE/XwU+Vz/5o4kU=
github.com/mattn/go-sqlite3 v1.14.22/go.mod h1:Uh1q+B4BYcTPb+yiD3kU8Ct7aC0hY9fxUwlHK0RXw+Y=
golang.org/x/text v0.20.0 h1:gK/Kv2otX8gz+wn7Rmb3vT96ZwuoxnQlY+HlJVj7Qug=
golang.org/x/text v0.20.0/go.mod h1:D4IsuqiFMhST5bX19pQ9ikHC2GsaKyk/oF+pn3ducp4=
gorm.io/driver/sqlite v1.6.0 h1:WHRRrIiulaPiPFmDcod6prc4l2VGVWHz80KspNsxSfQ=
gorm.io/driver/sqlite v1.6.0/go.mod h1:AO9V1qIQddBESngQUKWL9yoH93HIeA1X6V633rBwyT8=
gorm.io/gorm v1.31.2 h1:3o8FXNo9v9S858gil+3LlZA1LkCOzgb4g5BL64FgaCo=
gorm.io/gorm v1.31.2/go.mod h1:XyQVbO2k6YkOis7C2437jSit3SsDK72s7n7rsSHd+Gs=
```

Inspect the go.mod file

```bash
# cat go.mod
module example/main

go 1.26.4

require (
	gorm.io/driver/sqlite v1.6.0
	gorm.io/gorm v1.31.2
)

require (
	github.com/jinzhu/inflection v1.0.0 // indirect
	github.com/jinzhu/now v1.1.5 // indirect
	github.com/mattn/go-sqlite3 v1.14.22 // indirect
	golang.org/x/text v0.20.0 // indirect
)
```


## References

* https://bundler.io
* https://go.dev/doc/modules/managing-dependencies
* https://semver.org

## Wrap Up

