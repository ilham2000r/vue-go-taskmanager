package models

import (
	"time"

	"gorm.io/gorm"
)

type Task struct {
	gorm.Model
	TaskName    string    `json:"taskName"`
	Description string    `json:"description"`
	DueDate     time.Time `json:"dueDate" gorm:"type:date"`
	TaskStatus  bool      `json:"taskStatus" gorm:"default:false"`
	Priority    string    `json:"priority"`
	UserID      uint      `json:"userId"`
}
