### Chapter 13 - Database Migrations

> 

## Introduction

Your company's data is in most cases the reason that you build and maintain your apps. It's what distinguishes you from your competitors. Long ago I remember a coworker and I discussing the development of a new database structure and the web app that would be built to provide access to the data. I remember them saying something along the lines of "Let's get the database schema right before we start on the app, because the apps and the languages they're written in will come and go, but the data will persist longer than any app we create that has access to it." In my career, this statement has been true on almost every project that I've worked on. It's at the core of this guide, which is replacing one language and framework with another. Most apps, whether enterprise, social media or almost any other category, are just a means to generate and retrieve some data in a database.

Keeping your database schema in sync across all of your environments is key to software development. Ruby on Rails has a unique feature called database migrations. It versions database changes and allows you to write Ruby code that's database agnostic but converts at runtime into SQL code that runs a DDL (data definition language) command to modify the schema. Each time the migration task is run, it compares the versions of the files in the migration folder with those found in the `schema_migrations` table and only runs the code that's missing from that table, thus ensuring that it doesn't try to make the same changes twice.

## Examples

### Ruby

Rails provides an end-to-end solution with command-line options to manage your database schema. The `db` folder is where migration files are stored in a database-agnostic format.

Creating a new table.

```ruby
bin/rails generate migration CreateBenefits name:string benefit_type:integer activated_at:datetime

      invoke  active_record
      create    db/migrate/20260417073020_create_benefits.rb
```

The auto-generated migration file.

```ruby
cat db/migrate/20260417073020_create_benefits.rb
class CreateBenefits < ActiveRecord::Migration[8.1]
  def change
    create_table :benefits do |t|
      t.string :name
      t.integer :benefit_type
      t.datetime :activated_at

      t.timestamps
    end
  end
end
```

Running the migration.

```ruby
bin/rails db:migrate

== 20260417073020 CreateBenefits: migrating ===================================
-- create_table(:benefits)
   -> 0.0014s
== 20260417073020 CreateBenefits: migrated (0.0014s) ==========================
```

The contents of the schema.rb file. This file can be used to create our test database and empty databases in other environments (with the `bin/rails db:setup RAILS_ENV=test` command) vs running the migration files individually.

```ruby
cat db/schema.rb

# This file is auto-generated from the current state of the database. Instead
# of editing this file, please use the migrations feature of Active Record to
# incrementally modify your database, and then regenerate this schema definition.
#
# This file is the source Rails uses to define your schema when running `bin/rails
# db:schema:load`. When creating a new database, `bin/rails db:schema:load` tends to
# be faster and is potentially less error prone than running all of your
# migrations from scratch. Old migrations may fail to apply correctly if those
# migrations use external dependencies or application code.
#
# It's strongly recommended that you check this file into your version control system.

ActiveRecord::Schema[8.1].define(version: 2026_04_17_073020) do
  create_table "benefits", force: :cascade do |t|
    t.datetime "activated_at"
    t.integer "benefit_type"
    t.datetime "created_at", null: false
    t.string "name"
    t.datetime "updated_at", null: false
  end
end
```

### Go

## References

* https://guides.rubyonrails.org/command_line.html#migrations

## Wrap Up


[Next >>](150-chapter-14.md)

