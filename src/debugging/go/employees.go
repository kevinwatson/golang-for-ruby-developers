package main

import (
  "context"
  "encoding/json"
  "fmt"
  "gorm.io/driver/sqlite"
  "gorm.io/gorm"
  "log"
  "net/http"
)

type Employee struct {
  gorm.Model
  guid      string
  firstName string
  lastName  string
}

func employeeIndexHandler(w http.ResponseWriter, req *http.Request) {
  //list the employees

  db, err := gorm.Open(sqlite.Open("employees.db"), &gorm.Config{})
  if err != nil {
    log.Fatal("Could not connect to database employees.db: ", err)
  }

  ctx := context.Background()

  records, err := gorm.G[Employee](db).Find(ctx)
  if err != nil {
    fmt.Printf("Could not connect to database %s\n", err)
  }

  js, err := json.Marshal(records)
  if err != nil {
    http.Error(w, err.Error(), http.StatusInternalServerError)
    return
  }

  w.Header().Set("Content-Type", "application/json")
  w.Write(js)
}

func main() {
  db, err := gorm.Open(sqlite.Open("employees.db"), &gorm.Config{})
  if err != nil {
    log.Fatal("Could not connect to database employees.db: ", err)
  }

  db.AutoMigrate(&Employee{})

  http.HandleFunc("/employees", employeeIndexHandler)
  fmt.Println("Server starting on port 8080")
  err = http.ListenAndServe(":8080", nil)
  if err != nil {
    log.Fatal("Server exited with error: ", err)
  }
}
