### Chapter 4 - Models - Object Relational Mapping (ORM)

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
* Callbacks which run before committing a record

## Examples

I know most of my readers are software engineers and would prefer to look at some code. Let's look at some examples and compare ActiveRecord and GORM to perform the same operation of finding and updating a record from a database. Let's assume two apps are accessing the same table in the same database.

#### Table Definition

```sql
CREATE TABLE users (
  id serial,
  name text
)
```

#### Expected DML

Here's the expected SQL data manipulation language (DML) that each language and framework should generate.

|Command|Statement|
|---|---|
|Insert|`INSERT INTO users (name) VALUES ('John')`|
|Select|`SELECT * FROM users WHERE id = $1`|
|Update|`UPDATE users SET name = 'Ronald' WHERE id = $1`|
|Delete|`DELETE FROM users WHERE id = $1`|

### Ruby

#### Model Definition

```ruby
class User < ApplicationRecord
end
```

#### ActiveRecord Query Interface

```ruby
# create the user
user = User.create(name: "John")

# find the user
user = User.find_by(id: user.id)

# update the user
user.name = "Ronald"
user.save

# delete the user
user.destroy
```

### Go

GORM supports a couple of language options. This example will use Go generics.

```golang
package main

import (
  "context"
  [additional imports]
  "gorm.io/gorm"
)

// define the User
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

  // delete the user
  err = gorm.G[User](db).Where("id = ?", id).Delete(ctx)
}
```

## References

* https://go.dev/doc/tutorial/generics
* https://gorm.io/docs
* https://guides.rubyonrails.org/active_record_basics.html

## Wrap Up

ORMs provide an easy way to write native code to manipulate data in your database. The Ruby ActiveRecord query interface provides helper methods such as `save` and `update` on ApplicationRecord objects which makes the code you write very simple. Go's GORM library requires function chaining to provide the same effect, but the result is the same: readable code that retrieves and manipulates data stored outside of your program.
