package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name      string `json:"name"`
	Email     string `json:"email"`
	Password  string `json:"password"`
	Role      string `json:"role"`
	ProjectId int    `json:"project_id"`
	TaskId    int    `json:"task_id"`
}
