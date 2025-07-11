package handler

import (
	"blog-go/config"
	"blog-go/internal/dto"
	"blog-go/internal/model"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

func mapToCommentResponse(comment model.Comment) dto.CommentResponse {
	return dto.CommentResponse{
		ID:      comment.ID,
		Content: comment.Content,
		PostID:  comment.PostID,
		User: dto.AuthorInfo{
			ID:       comment.User.ID,
			Username: comment.User.Username,
		},
	}
}

func GetAllCommentByIdPost(c echo.Context) error {
	postID := c.Param("id")

	var comment []model.Comment
	if err := config.DB.Preload("User").Where("post_id = ?", postID).Order("created_at ASC").Find(&comment).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"message": "Failed To Fetch Comment", "error": err.Error()})
	}

	var res []dto.CommentResponse
	for _, p := range comment {
		res = append(res, mapToCommentResponse(p))
	}

	return c.JSON(http.StatusOK, res)
}

func CreateComment(c echo.Context) error {
	postID, _ := strconv.Atoi(c.Param("id"))

	var req dto.CommentRequest
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"message": "Invalid Input", "error": err.Error()})
	}

	if err := config.Validate.Struct(&req); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"message": "Validation Error", "error": err.Error()})
	}

	user := c.Get("user").(model.User)

	comment := model.Comment{
		Content: req.Content,
		PostID:  uint(postID),
		UserID:  user.ID,
	}

	if err := config.DB.Create(&comment).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"message": "Failed To Created Comment", "error": err.Error()})
	}

	return c.JSON(http.StatusCreated, echo.Map{"message": "Success To Created Comment"})
}

func UpdateComment(c echo.Context) error {
	CommentID := c.Param("id")
	user := c.Get("user").(model.User)

	var comment model.Comment
	if err := config.DB.First(&comment, CommentID).Error; err != nil {
		return c.JSON(http.StatusNotFound, echo.Map{"message": "Comment Not Found"})
	}

	if comment.UserID != user.ID {
		return c.JSON(http.StatusForbidden, echo.Map{"message": "You Can't Edit This Comment"})
	}

	var req dto.CommentRequest
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"message": "Invalid Input", "error": err.Error()})
	}

	if err := config.Validate.Struct(&req); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"message": "Validation Error", "error": err.Error()})
	}

	comment.Content = req.Content
	if err := config.DB.Save(&comment).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"message": "Update Failed", "error": err.Error()})
	}

	return c.JSON(http.StatusOK, echo.Map{"message": "Success To Updated Comment"})
}

func DeleteComment(c echo.Context) error {
	commentID := c.Param("id")
	user := c.Get("user").(model.User)

	var comment model.Comment
	if err := config.DB.First(&comment, commentID).Error; err != nil {
		return c.JSON(http.StatusNotFound, echo.Map{"message": "Comment Not Found"})
	}

	if comment.UserID != user.ID {
		return c.JSON(http.StatusForbidden, echo.Map{"message": "You Can't Delete This Comment"})
	}

	if err := config.DB.Unscoped().Delete(&comment).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"message": "Update Failed", "error": err.Error()})
	}

	return c.JSON(http.StatusOK, echo.Map{"message": "Success To Deleted Comment"})
}
