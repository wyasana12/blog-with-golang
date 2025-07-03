package main

import (
	"blog-go/config"
	"blog-go/middleware"
	"blog-go/routes"

	"github.com/labstack/echo/v4"
)

func main() {
	config.ConnectDB()
	config.Init()

	e := echo.New()

	e.Use(middleware.CORSMiddleware())

	routes.IndexRoutes(e)

	e.Logger.Fatal(e.Start(":8080"))
}
