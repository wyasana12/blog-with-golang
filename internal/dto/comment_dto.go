package dto

type CommentRequest struct {
	Content string `json:"content" validate:"required,min=3"`
}

type CommentResponse struct {
	ID      uint       `json:"id"`
	Content string     `json:"string"`
	PostID  uint       `json:"post_id"`
	User    AuthorInfo `json:"user"`
}
