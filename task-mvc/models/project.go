package models

import "gorm.io/gorm"

type Project struct {
	gorm.Model
	Title       string `json:"title"`
	Description string `json:"description"`
	UserId      int    `json:"user_id"`
	TaskId      int    `json:"task_id"`
}
