package routes

import (
	"blog-go/internal/handler"
	"blog-go/middleware"

	"github.com/labstack/echo/v4"
)

func UserRoleRoutes(g *echo.Group) {
	manage := g.Group("/users", middleware.AuthMiddleware, middleware.RoleMiddleware("superadmin"))

	manage.POST("/:id/roles/sync", handler.AssignRoleToUser)
	manage.DELETE("/:id/roles/:roleId", handler.RevokeRoleFromUser)
}
