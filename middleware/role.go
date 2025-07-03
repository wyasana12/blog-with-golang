package middleware

import (
	"blog-go/internal/model"
	"net/http"
	"strings"

	"github.com/labstack/echo/v4"
)

func RoleMiddleware(requiredRole ...string) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			user, ok := c.Get("user").(*model.User)
			if !ok || user == nil {
				return c.JSON(http.StatusUnauthorized, echo.Map{"message": "Unauthorized or User Not Found In Context"})
			}

			for _, role := range user.Roles {
				for _, allowed := range requiredRole {
					if strings.EqualFold(role.Name, allowed) {
						return next(c)
					}
				}
			}

			return c.JSON(http.StatusForbidden, echo.Map{"message": "Forbiden Role Not Allowed"})
		}
	}
}
