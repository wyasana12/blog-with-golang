package model

import (
	"time"

	"gorm.io/gorm"
)

type Post struct {
	gorm.Model
	Title       string     `json:"title"`
	Content     string     `json:"content" gorm:"type:text"`
	Status      string     `json:"status" gorm:"default:'draft'"`
	PublishedAt *time.Time `json:"published_at"`
	AuthorID    uint       `json:"author_id"`
	Author      User       `gorm:"foreignKey:AuthorID"`
}
