package handler

import (
	"github.com/hibiken/asynq"
	"gorm.io/gorm"
)

type Handler struct {
	DB     *gorm.DB
	Client *asynq.Client
}
