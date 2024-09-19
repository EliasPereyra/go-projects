package models

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {
	// https://gorm.io/docs/connecting_to_the_database.html
	// TODO: Set a stronger password

	database, err := gorm.Open(sqlite.Open("tasks.db"), &gorm.Config{})

	if err != nil {
		panic("failed to connect the db")
	}

	DB = database
}

func DBMigrate() {
	DB.AutoMigrate(&User{}, &Project{}, &Task{})
}
