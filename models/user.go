package models

import (
	"time"
)

type User struct {
	ID        uint      `json:"id" gorm:"primary_key"`
	Username  string    `json:"username" binding:"required"`
	Email     string    `json:"email" gorm:"unique" binding:"required"`
	Password  string    `json:"password" binding:"required,min=6"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Photos    []Photo   `json:"photos" gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}
