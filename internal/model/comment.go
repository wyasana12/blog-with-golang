package model

import "gorm.io/gorm"

type Comment struct {
	gorm.Model
	Content string `json:"content"`
	PostID  uint   `json:"post_id"`
	Post    Post   `gorm:"foreignKey:PostID"`
	UserID  uint   `json:"user_id"`
	User    User   `gorm:"foreignKey:UserID"`
}
