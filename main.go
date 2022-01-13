package main

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

const (
  host     = "localhost"
  port     = 5432
  user     = "foreignfood"
  password = "jello"
  dbname   = "calhounio_demo"
)

func main() {
  psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
    "password=%s dbname=%s sslmode=disable",
    host, port, user, password, dbname)
  psqlDB, err := sql.Open("postgres", psqlInfo)
  if err != nil {
    panic(err)
  }

  gormDB, err := gorm.Open(postgres.New(postgres.Config{
    Conn: psqlDB,
  }), &gorm.Config{})
  if err != nil {
    panic(err)
  }

type User struct {
  Age int
  Email string
  FirstName string
  LastName string
}

gormDB.AutoMigrate(&User{})

  id := 0
  gormDB.Create(&User{50, "thingOne@me.com", "Thing", "One"}).Scan(&id)


  fmt.Println("New record ID is:", id)
}
