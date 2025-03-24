package models

import "gorm.io/gorm"

type Task struct {
	gorm.Model
	Title string `json:"title" gorm:"not null"`
	Description string `json:"description" gorm:"not null"`
	Completed bool `json:"completed" gorm:"not null"`
	UserID uint `json:"user_id" gorm:"not null"`
}