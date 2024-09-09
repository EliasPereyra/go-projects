package models

import "gorm.io/gorm"

type Blog struct {
	gorm.Model
	Title string `gorm:"size:255"`
	Body  string `gorm:"type:text"`
}

func GetAll() *[]Blog {
	var posts []Blog
	DB.Where("deleted_at is NULL").Order("updated_at desc").Find(&posts)
	return &posts
}

func GetOnePost(id uint64) *Blog {
	var post Blog
	DB.Where("id = ?", id).First(&post)
	return &post
}
