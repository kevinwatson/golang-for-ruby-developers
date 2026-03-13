### Chapter 4 - Object Relational Mapping (ORMs)

> 

## Introduction

Object relational mapping is generally a library that converts code into structured query language (SQL) code that allows you to interact with your database. Ruby on Rails supports multiple databases and can convert the same Ruby code to the SQL that is specific to the database you're using. This allows for developer productivity, since they don't have to write SQL code and mix that in with their Ruby code.

## Comparing Languages

### Ruby on Rails

Ruby on Rails provides the ActiveRecord gem which provides a number of features. Some of these features include:

* Association support for related one-to-one, one-to-many and many-to-many relationships
* Database change support (aka migrations)
* CRUD methods such as `create`, `save`, `delete` and `find`
* Validations which can be configured to run before inserting or updating rows
* Callbacks which can run before or after certain events
* Reserved field names like `created_at`, `updated_at` which are automatically populated

### Go

Go has a few ORMs available. One of the most popular ORMs is GORM. GORM provides similar functionality to ActiveRecord. Let's take a look at some of its features:

* Association support for related one-to-one, one-to-many and many-to-many relationships
* Database change support (aka migrations)
* Functions such as `Create`, `Update`, `Delete` and `Find`
* Callbacks which can run before committing a record

## Examples

I know most of my readers are software engineers and would prefer to look at some code. Let's look at some examples and compare ActiveRecord and GORM to perform the same operation of finding and updating a record from a database. Let's assume two apps are accessing the same table in the same database.

#### Table Defintion

```sql
CREATE TABLE users (
  id serial,
  name text
)
```

### Ruby

#### Model Definition

```ruby
class User < ApplicationRecord
end
```

#### ActiveRecord Query Interface

Find the record

```ruby
# create the user
user = User.create(name: "John")

# find the user
user = User.find_by(id: user.id)

# update the user
user.name = "Ronald"
user.save
```

### Go

#### Model Definition

Using Generics

```golang
package main

import (
  "context"
  [additional imports]
  "gorm.io/gorm"
)

type User struct {
  gorm.Model
  Id int
  Name string
}

func main() {
  db, err := [connect to the db]
  ctx := context.Background()

  // define the user
  user := User{Name: "John"})

  // create the user
  result := gorm.WithResult()
  err = gorm.G[User](db, result).Create(ctx, &user)
  id := user.Id
  err = result.Error

  // find the user
  user, err := gorm.G[User](db).Where("id = ?", id).First(ctx)

  // update the user
  err = gorm.G[User](db).Where("id = ?", id).Update(ctx, "name", "Ronald")
}
```

## References

* https://gorm.io/docs
* https://guides.rubyonrails.org/active_record_basics.html

## Wrap Up

