package dto

import "time"

type PostRequest struct {
	Title   string `json:"title" validate:"required,min=3"`
	Content string `json:"content" validate:"required"`
	Status  string `json:"status" validate:"required,oneof=draft published"`
}

type PostResponse struct {
	ID          uint       `json:"id"`
	Title       string     `json:"title"`
	Content     string     `json:"content"`
	Status      string     `json:"status"`
	PublishedAt *time.Time `json:"published_at,omitempty"`
	Author      AuthorInfo `json:"author"`
	CreatedAt   time.Time  `json:"created_at"`
	UpdatedAt   time.Time  `json:"updated_at"`
}
