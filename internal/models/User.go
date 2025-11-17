package models

import (
	"time"
)

type User struct {
	ID        uint   `json:"id" gorm:"primaryKey"`
	Username  string `json:"username" gorm:"unique;not null"`
	Email     string `json:"email" gorm:"unique;not null"`
	Password  string `json:"password" gorm:"unique;not null"`
	AvatarURL string `json:"avatar_url"`

	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
