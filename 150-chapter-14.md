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

Hashes are key/value structures where the keys can be anything. In the example below we'll use a hash to track the number of each fruit that we have on hand.

Example

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

## Concurrency

## Loops

## Control Expressions

## Dependency Management

## References

* https://go.dev/talks/2012/splash.article
* https://www.rubyguides.com/2019/04/ruby-data-structures
* https://objectcomputing.com/resources/publications/sett/november-2018-way-to-go-part-1

## Wrap Up

