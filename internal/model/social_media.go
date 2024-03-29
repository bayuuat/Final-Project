package model

import (
	"time"

	"gorm.io/gorm"
)

type SocialMedia struct {
	ID             uint64         `gorm:"primaryKey" json:"id"`
	Name           string         `json:"name"`
	SocialMediaURL string         `json:"social_media_url"`
	UserID         uint64         `json:"user_id"`
	CreatedAt      time.Time      `json:"created_at"`
	UpdatedAt      time.Time      `json:"updated_at"`
	DeletedAt      gorm.DeletedAt `json:"-" gorm:"column:deleted_at"`
}

type SocialMediaPost struct {
	Name           string `json:"name"`
	SocialMediaURL string `json:"social_media_url"`
}
