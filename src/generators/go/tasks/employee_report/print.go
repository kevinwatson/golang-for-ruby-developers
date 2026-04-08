package main

import (
    "context"
    "example.com/my-app/models"
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

    db.AutoMigrate(&models.Employee{})

    err = gorm.G[models.Employee](db).Create(ctx, &models.Employee{Guid: "abcd", FirstName: "George", LastName: "Jetson"})

    records, err := gorm.G[models.Employee](db).Find(ctx)
    if err != nil {
      fmt.Printf("Could not connect to database %s\n", err)
    }

    fmt.Println("id,guid,first_name,last_name")
    for _, e := range records {
        fmt.Printf("%d,%s,%s,%s\n", e.Id, e.Guid, e.FirstName, e.LastName)
    }
}

