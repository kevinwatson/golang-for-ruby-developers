package main

import (
  "gorm.io/driver/sqlite"
  "gorm.io/gorm"
  "log"
  "net/http"
  "net/http/httptest"
  "testing"
)

func TestEmployeeIndexHandler(t *testing.T) {
  // initialize the database
  db, err := gorm.Open(sqlite.Open("employees.db"), &gorm.Config{})
  if err != nil {
    log.Fatal("Could not connect to database employees.db: ", err)
  }

  db.AutoMigrate(&Employee{})

  // create a 'get' request
  req, err := http.NewRequest("GET", "/employees", nil)
  if err != nil {
    t.Fatal(err)
  }
  // create a ResponseRecorder for the http.ResponseWriter
  rr := httptest.NewRecorder()
  handler := http.HandlerFunc(EmployeeIndexHandler)
  handler.ServeHTTP(rr, req)
  if status := rr.Code; status != http.StatusOK {
    t.Errorf("handler returned wrong status code: got %v, want %v", status, http.StatusOK)
  }
  expected := `[]`
  if rr.Body.String() != expected {
    t.Errorf("handler returned unexpected body: got %v, want %v", rr.Body.String(), expected)
  }
}

