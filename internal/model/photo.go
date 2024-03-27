package model

import (
	"time"

	"gorm.io/gorm"
)

type Photo struct {
	ID        uint64         `json:"id"`
	Title     string         `json:"title"`
	Caption   string         `json:"caption"`
	PhotoURL  string         `json:"photo_url"`
	UserID    uint64         `json:"user_id"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"-" gorm:"column:deleted_at"`
}

type PhotoPost struct {
	Title    string `json:"title" binding:"required"`
	Caption  string `json:"caption" binding:"required"`
	PhotoURL string `json:"photo_url" binding:"required"`
}
