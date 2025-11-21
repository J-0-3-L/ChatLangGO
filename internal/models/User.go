package models

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model //Aqui viene integrado ID,cretedAt,UpdateAt

	Username  string `json:"username" gorm:"unique;not null"`
	Email     string `json:"email" gorm:"unique;not null"`
	Password  string `json:"-"`
	AvatarURL string `json:"avatar_url"`
}
