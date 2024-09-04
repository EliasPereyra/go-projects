package models

import "gorm.io/gorm"

type Blog struct {
	gorm.Model
	Title string `gorm:"size:255"`
	Body  string `gorm:"type:text"`
}

func GetAll() {

}
