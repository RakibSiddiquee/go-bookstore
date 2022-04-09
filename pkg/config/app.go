package config

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db *gorm.DB

// Connect connects the databases
func Connect() {
	dns := "root:@tcp(127.0.0.1:3306)/book_store?charset=utf8&parseTime=True&loc=Local"
	d, err := gorm.Open(mysql.Open(dns), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	db = d
}

// GetDB gets the database
func GetDB() *gorm.DB {
	return db
}
