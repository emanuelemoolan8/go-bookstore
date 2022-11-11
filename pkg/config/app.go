package config

import (
	_ "github.com/jinzhu/dialects/mysql"
	"gorm.io/gorm"
)

var db *gorm.DB

func Connect() {
	d, err := gorm.Open("mysql", "hjfgkweqwlkjfehwljkqfgqkwegqwk")

	if err != nil {
		panic(err)
	}

	db = d
}

func GetDB() *gorm.DB {
	return db
}
