package models_test

import (
	"api-testing/models"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func SetupTest() error {
	connStr := "test:password@tcp(127.0.0.1:3306)/blog?charset=utf8&parseTime=true"
	db, err := gorm.Open(mysql.Open(connStr), &gorm.Config{})
	if err != nil {
		return err
	}

	models.DB = db
	models.DBMigrate()

	return nil
}

func TeardownTestDB() {
	var tableNames []string
	models.DB.Raw("SHOW TABLES").Scan(&tableNames)

	for _, tableName := range tableNames {
		models.DB.Exec("DROP TABLE IF EXISTS " + tableName)
	}
}
