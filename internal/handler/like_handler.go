package handler

import (
	"blog-go/config"
	"blog-go/internal/model"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

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
