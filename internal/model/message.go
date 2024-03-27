package model

import (
	"time"

	"gorm.io/gorm"
)

type Message struct {
	ID        uint64         `json:"id"`
	UserID    uint64         `json:"user_id"`
	PhotoID   uint64         `json:"photo_id"`
	Message   string         `json:"message"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"-" gorm:"column:deleted_at"`
}

type MessagePost struct {
	PhotoID uint64 `json:"photo_id"`
	Message string `json:"message"`
}
