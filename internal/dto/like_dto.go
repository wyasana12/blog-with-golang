package dto

type LikeResponse struct {
	ID   uint       `json:"id"`
	User AuthorInfo `json:"user"`
}
