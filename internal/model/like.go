package model

import "gorm.io/gorm"

type Like struct {
	gorm.Model
	UserID uint `json:"user_id"`
	PostID uint `json:"post_id"`

	User User `gorm:"foreignKey:UserID"`
	Post Post `gorm:"foreignKey:PostID"`
}
