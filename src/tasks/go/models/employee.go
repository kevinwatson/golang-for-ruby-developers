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

