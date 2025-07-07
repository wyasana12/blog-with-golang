package routes

import (
	"blog-go/internal/handler"
	"blog-go/middleware"

	"github.com/labstack/echo/v4"
)

func PostRoutes(g *echo.Group) {
	post := g.Group("/post")

	post.GET("/", handler.GetAllPublishedPosts)
	post.GET("/:id", handler.GetDetailPublishedPost)

	mypost := g.Group("/user/post", middleware.AuthMiddleware)
	mypost.GET("/", handler.GetAllMyPost)
	mypost.POST("/", handler.CreatePost)
	mypost.GET("/:id", handler.GetDetailMyPost)
	mypost.PUT("/:id", handler.UpdatePost)
	mypost.DELETE("/:id", handler.DeletePost)

	g.GET("/:username/posts", handler.GetPublishedPostByUsername)
}
