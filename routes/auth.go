package routes

import (
	"blog-go/internal/handler"
	"blog-go/middleware"

	"github.com/labstack/echo/v4"
)

func AuthRoutes(g *echo.Group) {
	auth := g.Group("/auth")

	auth.POST("/register", handler.Register)
	auth.POST("/login", handler.Login)

	auth.Use(middleware.AuthMiddleware)
	auth.GET("/me", handler.Me)
}
