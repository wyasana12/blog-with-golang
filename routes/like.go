package routes

import (
	"blog-go/internal/handler"
	"blog-go/middleware"

	"github.com/labstack/echo/v4"
)

func LikeRoutes(g *echo.Group) {
	like := g.Group("/post/:id", middleware.AuthMiddleware)
	like.PUT("/like", handler.ToggleLike)
	like.GET("/like", handler.GetAllUsersWhoLike)
}
