package config

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

var (
	db *gorm.DB
)

func Connect() {
	d, err := gorm.Open("postgres", "user=postgres password=collins2005 dbname=bookstore sslmode=disable")
	if err != nil {
		panic(err)
	}
	db = d

	if db.DB().Ping() != nil {
		panic("Failed to ping the database")
	}

	fmt.Println("Connected to the database successfully!")

}

func GetDB() *gorm.DB {
	return db
}
