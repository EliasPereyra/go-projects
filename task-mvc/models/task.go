package models

import "gorm.io/gorm"

type Task struct {
	gorm.Model
	Name      string `json:"name"`
	Completed bool   `json:"completed"`
	UserId    int    `json:"user_id"`
	ProjectId int    `json:"project_id"`
}
