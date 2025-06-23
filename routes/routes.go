package routes

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func Init(e *echo.Echo) {
	e.GET("/ping", func(c echo.Context) error {
		return c.JSON(http.StatusOK, map[string]string{"message": "ping"})
	})
}
