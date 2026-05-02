### Chapter 12 - Code Generators

![code generators](images/code-generators.png)

## Introduction

Code generators are command-line tools that generate files and folders to get you up and running quickly. Because Rails is built on the MVC framework, it includes a number of built-in generators that can generate views, controllers, models or `scaffold` which will generate all of the above. You can also create custom generators for your own needs.

Because Go is a typed language, it requires that we explicity define things like database models with structs (Ruby on Rails determines this at runtime by querying its database). Code generators like the built-in `go generate` command can be used to auto-generate some of this repetitive code.

Nowadays with generative AI you may be asking 'do we need generators? AI can create the code for us!' This may be true. AI can generate some or all of the code you need to add a new feature for common use cases (like adding a model file). But, one day we may have custom generators that the AI is not aware of that writes custom code needed to run our app. We can add the step of running our generator for specific cases to our AI agent's context. To guarantee consistent results, the agent can then be instructed to run the generator command at the appropriate time when making code changes.

## Examples

### Ruby

Rails includes a `scaffold` option that will create a full set of files for a resource, including views, a controller, a model and a database migration file. For example, to add a `benefits` table to our employee database with a model, controller and views we can run this command:

```ruby
bin/rails generate scaffold benefit name:string benefit_type:integer activated_at:datetime

      invoke  active_record
      create    db/migrate/20260408080055_create_benefits.rb
      create    app/models/benefit.rb
      invoke    test_unit
      create      test/models/benefit_test.rb
      create      test/fixtures/benefits.yml
      invoke  resource_route
       route    resources :benefits
      invoke  scaffold_controller
      create    app/controllers/benefits_controller.rb
      invoke    erb
      create      app/views/benefits
      create      app/views/benefits/index.html.erb
      create      app/views/benefits/edit.html.erb
      create      app/views/benefits/show.html.erb
      create      app/views/benefits/new.html.erb
      create      app/views/benefits/_form.html.erb
      create      app/views/benefits/_benefit.html.erb
      invoke    resource_route
      invoke    test_unit
      create      test/controllers/benefits_controller_test.rb
      invoke    helper
      create      app/helpers/benefits_helper.rb
      invoke      test_unit
      invoke    jbuilder
      create      app/views/benefits/index.json.jbuilder
      create      app/views/benefits/show.json.jbuilder
      create      app/views/benefits/_benefit.json.jbuilder
```

We can then customize the files to complete our new feature.

### Go

Various Go frameworks provide code generation tooling (e.g. [the Buffalo framework](https://gobuffalo.io/documentation/getting_started/new-project/)) where you can run commands like `buffalo new console-app` which will generate all of the necessary files to get you up and running quickly.

Ruby is a dynamic language and an ActiveModel library in the Rails framework automatically adds methods for the columns it finds in its database (e.g `employee.id` for the `id` field). Go is a typed, compiled language. In order to provide the same functionality which translates to developer productivity, we need to write the code in Go that Rails generates automatically.

For our purposes, let's use the Gorm framework to generate helper methods to access our database.

Our Dockerfile.

```dockerfile
# Dockerfile
FROM golang:1.25

WORKDIR /opt/app
RUN go mod init example.com/my-app
COPY . .
#RUN go get ./...
RUN go get -u gorm.io/gen@latest
RUN go get -u golang.org/x/tools@latest
RUN go mod tidy
```

Our docker-compose file.

```yaml
services:
  go:
    build:
      context: .
    stdin_open: true
    tty: true
```

Our employee model.

```golang
# models/employee.go
package models

import (
  "gorm.io/gorm"
)

type Employee struct {
  gorm.Model
  Id        int
  Guid      string
  FirstName string
  LastName  string
}
```

Our Gorm data access layer (dal) code generator.

```golang
# cmd/gen/main.go
package main

import (
    "example.com/my-app/models"
    "log"
    "gorm.io/driver/sqlite"
    "gorm.io/gen"
    "gorm.io/gorm"
)

func main() {
    db, err := gorm.Open(sqlite.Open("employees.db"), &gorm.Config{})
    if err != nil {
      log.Fatal("Could not connect to database employees.db: ", err)
    }

    // create the employee table
    db.AutoMigrate(&models.Employee{})

    g := gen.NewGenerator(gen.Config{
        OutPath: "internal/dal/query",
	Mode:    gen.WithDefaultQuery | gen.WithQueryInterface, // enable query.SetDefault(db)
    })

    g.UseDB(db)
    g.ApplyBasic(g.GenerateAllTable()...)
    g.Execute()
}
```

Now let's build the image and log in.

```bash
docker-compose build
docker-compose run go bash
```

Once we have logged in, we will run our cmd/gen main package to create a few files.

```bash
go run ./cmd/gen
```

This will generate the following output. Note the files it created. 

```bash
find 2 table from db: [employees sqlite_sequence]
got 7 columns from table <employees>
got 2 columns from table <sqlite_sequence>
Start generating code.
generate model file(table <sqlite_sequence> -> {model.SqliteSequence}): /opt/app/internal/dal/model/sqlite_sequence.gen.go
generate model file(table <employees> -> {model.Employee}): /opt/app/internal/dal/model/employees.gen.go
generate query file: /opt/app/internal/dal/query/sqlite_sequence.gen.go
generate query file: /opt/app/internal/dal/query/employees.gen.go
generate query file: /opt/app/internal/dal/query/gen.go
Generate code done.
```

Inspecting one of our new files

```golang
# grep func /opt/app/internal/dal/query/employees.gen.go

func newEmployee(db *gorm.DB, opts ...gen.DOOption) employee {
func (e employee) Table(newTableName string) *employee {
func (e employee) As(alias string) *employee {
func (e *employee) updateTableName(table string) *employee {
func (e *employee) WithContext(ctx context.Context) IEmployeeDo { return e.employeeDo.WithContext(ctx) }
func (e employee) TableName() string { return e.employeeDo.TableName() }
func (e employee) Alias() string { return e.employeeDo.Alias() }
func (e employee) Columns(cols ...field.Expr) gen.Columns { return e.employeeDo.Columns(cols...) }
func (e *employee) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
func (e *employee) fillFieldMap() {
func (e employee) clone(db *gorm.DB) employee {
func (e employee) replaceDB(db *gorm.DB) employee {
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	Scopes(funcs ...func(gen.Dao) gen.Dao) IEmployeeDo
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.Employee, err error)
	FindInBatches(result *[]*model.Employee, batchSize int, fc func(tx gen.Dao, batch int) error) error
func (e employeeDo) Debug() IEmployeeDo {
func (e employeeDo) WithContext(ctx context.Context) IEmployeeDo {
func (e employeeDo) ReadDB() IEmployeeDo {
func (e employeeDo) WriteDB() IEmployeeDo {
func (e employeeDo) Session(config *gorm.Session) IEmployeeDo {
func (e employeeDo) Clauses(conds ...clause.Expression) IEmployeeDo {
func (e employeeDo) Returning(value interface{}, columns ...string) IEmployeeDo {
func (e employeeDo) Not(conds ...gen.Condition) IEmployeeDo {
func (e employeeDo) Or(conds ...gen.Condition) IEmployeeDo {
func (e employeeDo) Select(conds ...field.Expr) IEmployeeDo {
func (e employeeDo) Where(conds ...gen.Condition) IEmployeeDo {
func (e employeeDo) Order(conds ...field.Expr) IEmployeeDo {
func (e employeeDo) Distinct(cols ...field.Expr) IEmployeeDo {
func (e employeeDo) Omit(cols ...field.Expr) IEmployeeDo {
func (e employeeDo) Join(table schema.Tabler, on ...field.Expr) IEmployeeDo {
func (e employeeDo) LeftJoin(table schema.Tabler, on ...field.Expr) IEmployeeDo {
func (e employeeDo) RightJoin(table schema.Tabler, on ...field.Expr) IEmployeeDo {
func (e employeeDo) Group(cols ...field.Expr) IEmployeeDo {
func (e employeeDo) Having(conds ...gen.Condition) IEmployeeDo {
func (e employeeDo) Limit(limit int) IEmployeeDo {
func (e employeeDo) Offset(offset int) IEmployeeDo {
func (e employeeDo) Scopes(funcs ...func(gen.Dao) gen.Dao) IEmployeeDo {
	return e.withDO(e.DO.Scopes(funcs...))
func (e employeeDo) Unscoped() IEmployeeDo {
func (e employeeDo) Create(values ...*model.Employee) error {
func (e employeeDo) CreateInBatches(values []*model.Employee, batchSize int) error {
func (e employeeDo) Save(values ...*model.Employee) error {
func (e employeeDo) First() (*model.Employee, error) {
func (e employeeDo) Take() (*model.Employee, error) {
func (e employeeDo) Last() (*model.Employee, error) {
func (e employeeDo) Find() ([]*model.Employee, error) {
func (e employeeDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.Employee, err error) {
	err = e.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
func (e employeeDo) FindInBatches(result *[]*model.Employee, batchSize int, fc func(tx gen.Dao, batch int) error) error {
func (e employeeDo) Attrs(attrs ...field.AssignExpr) IEmployeeDo {
func (e employeeDo) Assign(attrs ...field.AssignExpr) IEmployeeDo {
func (e employeeDo) Joins(fields ...field.RelationField) IEmployeeDo {
func (e employeeDo) Preload(fields ...field.RelationField) IEmployeeDo {
func (e employeeDo) FirstOrInit() (*model.Employee, error) {
func (e employeeDo) FirstOrCreate() (*model.Employee, error) {
func (e employeeDo) FindByPage(offset int, limit int) (result []*model.Employee, count int64, err error) {
func (e employeeDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
func (e employeeDo) Scan(result interface{}) (err error) {
func (e employeeDo) Delete(models ...*model.Employee) (result gen.ResultInfo, err error) {
func (e *employeeDo) withDO(do gen.Dao) *employeeDo {
```

In our project, we can now query for employees with code like the following:

```golang
package main

import (
	"context"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"

	"example.com/my-app/internal/dal/query"
)

func main() {
	db, err := gorm.Open(sqlite.Open("employees.db"), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	query.SetDefault(db)

	e := query.Employee
	_, _ = e.WithContext(context.Background()).
		Where(e.Guid.Eq('abcd')).
		Find()
}
```

## References

* https://github.com/go-gorm/gen?tab=readme-ov-file#quick-start
* https://go.dev/blog/generate
* https://go.googlesource.com/proposal/+/refs/heads/master/design/go-generate.md
* https://guides.rubyonrails.org/command_line.html#generating-code

## Wrap Up

As we can see, code generators are an essential part of your toolkit. They can be used at the beginning of a project to add the files and folders to get started on that new feature, or run at the end to generate the boilerplate code after the logic or models have been added.

[Next >>](140-chapter-13.md)

