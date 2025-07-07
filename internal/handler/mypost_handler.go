package handler

import (
	"blog-go/config"
	"blog-go/internal/dto"
	"blog-go/internal/model"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
)

func GetAllMyPost(c echo.Context) error {
	user := c.Get("user").(model.User)
	status := c.QueryParam("status")

	var posts []model.Post
	query := config.DB.Preload("Author").Where("author_id = ?", user.ID)

	if status != "" {
		query = query.Where("status = ?", status)
	}

	if err := query.Order("created_at desc").Find(&posts).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"message": "Failed To Fetch My Post", "error": err.Error()})
	}

	var res []dto.PostResponse
	for _, p := range posts {
		res = append(res, mapToPostResponse(p))
	}

	return c.JSON(http.StatusOK, res)
}

func GetDetailMyPost(c echo.Context) error {
	user := c.Get("user").(model.User)
	id := c.Param("id")

	var post model.Post
	if err := config.DB.Preload("Author").Where("id = ? AND author_id = ?", id, user.ID).First(&post).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"message": "Failed To Fetch My Detail Post", "error": err.Error()})
	}

	return c.JSON(http.StatusOK, post)
}

func CreatePost(c echo.Context) error {
	user := c.Get("user").(model.User)

	var req dto.PostRequest
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"message": "Invalid Input", "error": err.Error()})
	}

	if err := config.Validate.Struct(req); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"message": "Validation Failed", "error": err.Error()})
	}

	post := model.Post{
		Title:    req.Title,
		Content:  req.Content,
		Status:   req.Status,
		AuthorID: user.ID,
	}

	if req.Status == "published" {
		now := time.Now()
		post.PublishedAt = &now
	}

	if err := config.DB.Create(&post).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"message": "Failed To Create Post"})
	}

	return c.JSON(http.StatusCreated, echo.Map{"message": "Success To Created Post"})
}

func UpdatePost(c echo.Context) error {
	user := c.Get("user").(model.User)
	id := c.Param("id")

	var post model.Post
	if err := config.DB.First(&post, id).Error; err != nil {
		return c.JSON(http.StatusNotFound, echo.Map{"message": "Post Not Found"})
	}

	if post.AuthorID != user.ID {
		return c.JSON(http.StatusForbidden, echo.Map{"message": "You Are Not The Owner"})
	}

	var req dto.PostRequest
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"message": "Invalid Input", "error": err.Error()})
	}

	if err := config.Validate.Struct(req); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"message": "Validation Failed", "error": err.Error()})
	}

	post.Title = req.Title
	post.Content = req.Content
	post.Status = req.Status
	if req.Status == "published" && post.PublishedAt == nil {
		now := time.Now()
		post.PublishedAt = &now
	}

	if err := config.DB.Save(&post).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"message": "Updated Failed", "error": err.Error()})
	}

	return c.JSON(http.StatusOK, echo.Map{"message": "Success To Updated Post"})
}

func DeletePost(c echo.Context) error {
	user := c.Get("user").(model.User)
	id := c.Param("id")

	var post model.Post
	if err := config.DB.First(&post, id).Error; err != nil {
		return c.JSON(http.StatusNotFound, echo.Map{"message": "Post Not Found"})
	}

	if post.AuthorID != user.ID {
		return c.JSON(http.StatusForbidden, echo.Map{"message": "You Are Not The Owner"})
	}

	if err := config.DB.Delete(&post).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"message": "Delete Failed", "error": err.Error()})
	}

	return c.JSON(http.StatusOK, echo.Map{"message": "Success To Deleted Post"})
}
