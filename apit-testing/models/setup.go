package models

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() {
	db, err := gorm.Open(mysql.Open("test:password@tcp(127.0.0.1:3306)/blog?charset=utf8&parseTime=true"), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to the Database")
	}

	DB = db
}

func DBMigrate() {
	DB.AutoMigrate(&Blog{})
}
