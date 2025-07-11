package routes

import (
	"blog-go/internal/handler"
	"blog-go/middleware"

	"github.com/labstack/echo/v4"
)

func CommentRoutes(g *echo.Group) {
	g.GET("/post/:id/comments", handler.GetAllCommentByIdPost)

	comment := g.Group("/comments", middleware.AuthMiddleware)
	comment.PUT("/:id", handler.UpdateComment)
	comment.DELETE("/:id", handler.DeleteComment)
	g.POST("/post/:id/comments", handler.CreateComment, middleware.AuthMiddleware)
}
