package routes

import (
	"blog-go/internal/handler"
	"blog-go/middleware"

	"github.com/labstack/echo/v4"
)

func RoleRoutes(g *echo.Group) {
	role := g.Group("/role", middleware.AuthMiddleware, middleware.RoleMiddleware("superadmin"))

	role.GET("/", handler.GetAllRoles)
	role.POST("/", handler.CreateRole)
	role.PUT("/:id", handler.UpdateRole)
	role.DELETE("/:id", handler.DeleteRole)
}
