package middleware

import (
	"blog-go/config"
	"blog-go/internal/model"
	"net/http"
	"os"
	"strings"

	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
)

func AuthMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		auth := c.Request().Header.Get("Authorization")
		if !strings.HasPrefix(auth, "Bearer ") {
			return c.JSON(http.StatusUnauthorized, echo.Map{"message": "Missing Token"})
		}

		tokenStr := strings.TrimPrefix(auth, "Bearer ")
		token, err := jwt.Parse(tokenStr, func(t *jwt.Token) (interface{}, error) {
			secret := os.Getenv("JWT_SECRET")
			if secret == "" {
				secret = "secret"
			}

			return []byte(secret), nil
		})

		if err != nil || !token.Valid {
			return c.JSON(http.StatusUnauthorized, echo.Map{"message": "Invalid Token"})
		}

		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok || !token.Valid {
			return c.JSON(http.StatusUnauthorized, echo.Map{"message": "Invalid Token"})
		}

		userID := uint(claims["user_id"].(float64))

		var user model.User
		if err := config.DB.Preload("Roles").First(&user, userID).Error; err != nil {
			return c.JSON(http.StatusUnauthorized, echo.Map{"message": "User Not Found"})
		}

		c.Set("user", user)
		return next(c)
	}
}
