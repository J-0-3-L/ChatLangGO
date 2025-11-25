package models

import "gorm.io/gorm"

type Post struct {
	gorm.Model

	Content string `json:"content" gorm:"type:text;not null"`

	UserID uint `json:"user_id"`
	User   User `json:"user" gorm:"foreignKey:UserID"`
}
