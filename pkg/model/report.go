package model

import "gorm.io/gorm"

type Report struct {
	gorm.Model
	Name string `json:"name"`
	Description string `json:"description"`
}
