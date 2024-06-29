package models

import (
	"time"
)

type Photo struct {
	ID        uint      `json:"id" gorm:"primary_key"`
	Title     string    `json:"title" binding:"required"`
	Caption   string    `json:"caption"`
	PhotoUrl  string    `json:"photo_url" binding:"required"`
	UserID    uint      `json:"user_id" binding:"required"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
