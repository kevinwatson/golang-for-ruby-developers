### Chapter 14 - Language Comparison

![language comparison](images/language-comparison.png)

## Introduction

The Ruby language is an interpreted, object-oriented language that was designed with software developers in mind. Its interpreted nature allows you to run code in an interactive console as well as make changes to a file while in development and see those changes instantly. It has a rich collection of libraries, or 'gems' available which provide additional functionality beyond the standard library. Rails and other frameworks provide additional functionality.

The Go language is a compiled language designed around simplicity and performance. It's compiled to platform-specific machine code that is optmized for high performance and low latency for network-enabled applications. It also has a rich collection of libraries, or 'packages' that provide additional functionality.

In this chapter, we'll focus on the similarities and differences in the core or standard libraries for each language as well as dependency management to add features.

## Data Structures

Data structures are a way to organize and manage data. Each language provides built in structures that can be implemented depending on the requirements of the app.

### Arrays

#### Ruby

Arrays in Ruby are managed by the language, with features such as automatically resizing the array when items are added or deleted. In the example below we'll collect various fruits, access a couple of them and delete one of them from the array.

```ruby
fruits = []
fruits << "apple"
fruits << "banana"
fruits << "blueberry"
fruits
=> ["apple", "banana", "blueberry"]

fruits.size
=> 3

fruits[0]
=> "apple"

fruits[1]
=> "banana"

fruits.delete("banana") # returns the deleted element
=> "banana"

fruits
=> ["apple", "blueberry"]

fruits.size
=> 2
```

#### Go

Go offers two structures, arrays and slices. Arrays have a fixed length, while slices offer flexibility such as automatic resizing but under the hood point to an underlying array. When a slice grows beyond it's underlying array's capacity, the runtime creates a larger array and copies the data to the new array.

In this example we'll declare an array for our top 3 favorite fruits.

```golang
var fruits [3]string
fruits[0] = "apple"
fruits[1] = "banana"
fruits[2] = "blueberry"
fmt.Println(fruits)
// [apple banana blueberry]

fmt.Println(len(fruits))
// 3

// because array elements can't be removed or deleted,
// we can only set an element back to its default or 'zero' value
fruits[1] = ""

fmt.Println(fruits)
// [apple  blueberry]

fmt.Println(len(fruits))
// 3
```

In this example we'll initialize a slice which will behave more like an array in Ruby.

```golang
var fruits []string
fruits = append(fruits, "apple", "banana", "blueberry")
fmt.Println(fruits)
// [apple banana blueberry]

fmt.Println(len(fruits))
// 3

// remove an element with the newer Delete function
// requires importing the "slices" package
fruits = slices.Delete(fruits, 1, 2)

fmt.Println(fruits)
// [apple blueberry]

fmt.Println(len(fruits))
// 2
```

### Hash

Hashes in Ruby are key/value structures where the keys can be anything. The equivalent in Go is called a map.

#### Ruby

In the example below we'll use a hash to track the number of each fruit that we have on hand.

```ruby
fruits = {}
fruits[:apple] = 2
fruits[:banana] = 1
fruits[:blueberry] = 5
fruits
=> {:apple => 2, :banana => 1, :blueberry => 5}

fruits.size
=> 3

fruits[:apple]
=> 2

fruits[:banana]
=> 1

fruits.delete(:banana) # returns the value of the key/value pair
=> 1

fruits
=> {:apple=>2, :blueberry=>2}

fruits.size
=> 2
```

#### Go

In this example we'll create a map to track the same counts of fruit.

```golang
fruits := make(map[string]int)
fruits["apple"] = 2
fruits["banana"] = 1
fruits["blueberry"] = 5
fmt.Println(fruits)
// map[apple:2 banana:1 blueberry:5]

fmt.Println(len(fruits))
// 3

fmt.Println(fruits["apple"])
// 2

fmt.Println(fruits["banana"])
// 1

delete(fruits, "banana")

fmt.Println(fruits)
// map[apple:2 blueberry:5]

fmt.Println(len(fruits))
// 2
```

## Loops

Loop constructs allow you to iterate through a collection of items, such as an array or hash/map.

### Ruby

Ruby has a number of ways to loop through collections, including `while`, `until`, `for`, `loop do...end` and iterators like `each`, `times` and `upto/downto`. We'll look at the commonly-used `each` iterator on an array.

```ruby
fruits = []
fruits << "apple"
fruits << "banana"
fruits << "blueberry"

fruits.each_with_index do |fruit, index|
  puts index, fruit
end

# 0
# apple
# 1
# banana
# 2
# blueberry
```

### Go

Go, on the other hand, only has a single loop construct: `for`. Depending on the statements used that follow the `for` keyword (initialization, condition, post-statement), this single construct can behave like several different looping constructs from other languages.

```golang
var fruits [3]string
fruits[0] = "apple"
fruits[1] = "banana"
fruits[2] = "blueberry"

for index, fruit := range fruits {
  fmt.Println(index, fruit)
}
// 0 apple
// 1 banana
// 2 blueberry
```

## Conditional Statements

Conditional statements check whether logic evaluates to true or false.

### Ruby

In Ruby, there are `falsy` and `truthy` statements. The number of `falsy` statements is very small, including only `nil` and `false`. Every other value is considered `true`, including `0` and empty strings `""`.

if, elseif, else

```ruby
require 'date'

todays_date = Date.today

if todays_date.saturday? || todays_date.sunday?
  puts "It's the weekend"
elsif todays_date.wday < 3
  puts "It's the start of the work week"
elsif todays_date.wednesday?
  puts "It's midweek"
else
  puts "It's the end of the work week"
end

# It's the start of the work week
```

### Go

In Go, there are only `if`, `if else` and `if else if` conditional statements.

```golang
import (
  "fmt"
  "time"
)

func main() {
  now := time.Now()
  weekday := int(now.Weekday())

  if weekday == 0 || weekday == 7 {
    fmt.Println("It's the weekend")
  } else if weekday < 3 {
    fmt.Println("It's the start of the work week")
  } else if weekday == 3 {
    fmt.Println("It's midweek")
  } else {
    fmt.Println("It's the end of the week")
  }
}

// It's the start of the work week
```
 
## Dependency Management

Dependent libraries or 'dependencies' are a way to quickly add features to our app. For example, Ruby on Rails is a collection of libraries (aka 'gems') that can be added either individually (for example, if you only need the features provided by ActiveRecord, you can include just that library and use its features) or as a collection if you want to use all of the framework's features.

### Ruby

Ruby's libraries are called 'gems' which is a clever term based on the name of the Ruby language. Gems, or gemstones, are defined as a precious piece of mineral crystal, organic matter or rock that has been cut and polished. There might be gems embedded in the rings on your hands or other jewelry. In Ruby terminology a 'gem' is a valuable and polished add-on that enhances your core app.

Ruby includes a tool named Bundler. Bundler uses two files to manage an app's dependencies or gems. These files are `Gemfile` and `Gemfile.lock`. The `Gemfile` defines which top level libraries the app depends on in order to run correctly. The `Gemfile.lock` file is maintained by the Bundler app and defines the specific versions and dependency tree and is read when using `bundle exec` to start the app.



## References

* https://bundler.io
* https://go.dev/talks/2012/splash.article
* https://www.rubyguides.com/2019/04/ruby-data-structures
* https://objectcomputing.com/resources/publications/sett/november-2018-way-to-go-part-1

## Wrap Up

