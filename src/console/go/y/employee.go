package employee

import (
  "context"
  "gorm.io/driver/sqlite"
  "gorm.io/gorm"
)

type Employee struct {
  gorm.Model
  FirstName string
  LastName  string 
}

func main() {
  db, err := gorm.Open(sqlite.Open("employees.db"), &gorm.Config{})
  if err != nil {
    panic("failed to connect database")
  }

  ctx := context.Background()

  // Migrate the schema
  db.AutoMigrate(&Employee{})

/* show examples in the console
  // Create
  err = gorm.G[Employee](db).Create(ctx, &Employee{Guid: , FirstName: "", LastName: ""})

  // Read
  product, err := gorm.G[Employee](db).Where("guid = ?", 1).First(ctx)

  // Update - update product's price to 200
  err = gorm.G[Product](db).Where("id = ?", product.ID).Update(ctx, "Price", 200)
  // Update - update multiple fields
  err = gorm.G[Product](db).Where("id = ?", product.ID).Updates(ctx, Product{Code: "D42", Price: 100})

  // Delete - delete product
  err = gorm.G[Product](db).Where("id = ?", product.ID).Delete(ctx)
*/
}
