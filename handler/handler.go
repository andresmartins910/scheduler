package handler

import "gorm.io/gorm"

type Handler struct {
	DB *gorm.DB
}

// TODO: separar handler de testes