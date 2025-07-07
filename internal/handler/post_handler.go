package handler

import (
	"blog-go/config"
	"blog-go/internal/dto"
	"blog-go/internal/model"
	"net/http"

	"github.com/labstack/echo/v4"
)

func mapToPostResponse(post model.Post) dto.PostResponse {
	return dto.PostResponse{
		ID:          post.ID,
		Title:       post.Title,
		Content:     post.Content,
		PublishedAt: post.PublishedAt,
		Author: dto.AuthorInfo{
			ID:   post.Author.ID,
			Name: post.Author.Name,
		},
		CreatedAt: post.CreatedAt,
		UpdatedAt: post.UpdatedAt,
	}
}

func GetAllPublishedPosts(c echo.Context) error {
	var posts []model.Post

	if err := config.DB.Preload("Author").Where("status = ?", "published").Order("published_at desc").Find(&posts).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"message": "Failed To Fetch Posts", "error": err.Error()})
	}

	var res []dto.PostResponse
	for _, p := range posts {
		res = append(res, mapToPostResponse(p))
	}

	return c.JSON(http.StatusOK, res)
}

func GetDetailPublishedPost(c echo.Context) error {
	id := c.Param("id")

	var post model.Post

	if err := config.DB.Preload("Author").Where("id = ? AND status = ?", id, "published").Find(&post).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"message": "Failed To Fetch Detail Post", "error": err.Error()})
	}

	return c.JSON(http.StatusOK, post)
}

func GetPublishedPostByUsername(c echo.Context) error {
	username := c.Param("username")

	var user model.User
	if err := config.DB.Where("username = ?", username).First(&user).Error; err != nil {
		return c.JSON(http.StatusNotFound, echo.Map{"message": "User Not Found"})
	}

	var posts []model.Post

	if err := config.DB.Preload("Author").Where("author_id = ? AND status = ?", user.ID, "published").Order("published_at desc").Find(&posts).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"message": "Failed To Fetch Post", "error": err.Error()})
	}

	var res []dto.PostResponse
	for _, p := range posts {
		res = append(res, mapToPostResponse(p))
	}

	return c.JSON(http.StatusOK, res)
}
