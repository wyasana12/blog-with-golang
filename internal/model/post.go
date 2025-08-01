package model

import (
	"time"

	"gorm.io/gorm"
)

type Post struct {
	gorm.Model
	Title           string     `json:"title"`
	Content         string     `json:"content" gorm:"type:text"`
	Status          string     `json:"status" gorm:"default:'draft'"`
	PublishedAt     *time.Time `json:"published_at"`
	AuthorID        uint       `json:"author_id"`
	Author          User       `gorm:"foreignKey:AuthorID"`
	DisableComments bool       `json:"disable_comments" gorm:"default:false"`
	HideLikes       bool       `json:"hide_likes" gorm:"default:false"`
	Likes           []Like     `json:"likes,omitempty"`
	Comments        []Comment  `json:"comments,omitempty"`
}
