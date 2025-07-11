package routes

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func IndexRoutes(e *echo.Echo) {
	routes := e.Group("/api")

	e.GET("/ping", func(c echo.Context) error {
		return c.JSON(http.StatusOK, map[string]string{"message": "ping"})
	})

	AuthRoutes(routes)
	RoleRoutes(routes)
	UserRoleRoutes(routes)
	PostRoutes(routes)
	CommentRoutes(routes)
}
