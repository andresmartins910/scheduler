package model

import "gorm.io/gorm"

type Task struct {
	gorm.Model
	Title  string `json:"title"`
	Status string `json:"status"`
}