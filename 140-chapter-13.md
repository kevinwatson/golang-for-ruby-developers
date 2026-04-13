### Chapter 13 - Database Migrations

> 

## Introduction

Your company's data is in most cases the reason that you build and maintain your apps. It's what distinguishes you from your competitors. Long ago I remember a coworker and I discussing the development of a new database structure and the web app that would be built to provide access to the data. I remember them saying something along the lines of "Let's get the database schema right before we start on the app, because the apps and the languages they're written in will come and go, but the data will persist longer than any app we create that has access to it." In my career, this statement has been true on almost every project that I've worked on. It's at the core of this guide, which is replacing one language and framework with another. Most apps, whether enterprise, social media or almost any other category, are just a means to generate and retrieve some data in a database.

Keeping your database schema in sync across all of your environments is key to software development. Ruby on Rails has a unique feature called database migrations. It versions database changes and allows you to write Ruby code that's database agnostic but converts at runtime into SQL code that runs a DDL (data definition language) command to modify the schema. Each time the migration task is run, it compares the versions of the files in the migration folder with those found in the `schema_migrations` table and only runs the code that's missing from that table, thus ensuring that it doesn't try to make the same changes twice.

## Examples

### Ruby

### Go

## References

## Wrap Up


[Next >>](150-chapter-14.md)

