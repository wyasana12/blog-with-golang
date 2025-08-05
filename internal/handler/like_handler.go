package handler

import (
	"blog-go/config"
	"blog-go/internal/dto"
	"blog-go/internal/model"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

func mapToLikeResponse(like model.Like) dto.LikeResponse {
	return dto.LikeResponse{
		ID: like.ID,
		User: dto.AuthorInfo{
			ID:       like.User.ID,
			Username: like.User.Username,
		},
	}
}

func ToggleLike(c echo.Context) error {
	user := c.Get("user").(model.User)
	postID, _ := strconv.Atoi(c.Param("id"))

	var existing model.Like
	if err := config.DB.Where("user_id = ? AND post_id = ?", user.ID, postID).First(&existing).Error; err == nil {
		config.DB.Unscoped().Delete(&existing)
		return c.JSON(http.StatusOK, echo.Map{"message": "Unliked"})
	}

	like := model.Like{UserID: user.ID, PostID: uint(postID)}
	if err := config.DB.Create(&like).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"message": "Failed To Like"})
	}

	return c.JSON(http.StatusOK, echo.Map{"message": "Liked"})
}

func GetAllUsersWhoLike(c echo.Context) error {
	postID, _ := strconv.Atoi(c.Param("id"))

	var post model.Post

	if err := config.DB.First(&post, postID).Error; err != nil {
		return c.JSON(http.StatusNotFound, echo.Map{"message": "Post Not Found"})
	}

	var like []model.Like
	if err := config.DB.Preload("User").Where("post_id = ?", postID).Order("updated_at ASC").Find(&like).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"message": "Failed To Fetch Like", "error": err.Error()})
	}

	var res []dto.LikeResponse
	for _, p := range like {
		res = append(res, mapToLikeResponse(p))
	}

	return c.JSON(http.StatusOK, res)
}
