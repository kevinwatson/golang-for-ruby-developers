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

