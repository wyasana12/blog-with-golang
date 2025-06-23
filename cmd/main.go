package main

import (
	"blog-go/config"
	"blog-go/routes"

	"github.com/labstack/echo/v4"
)

func main() {
	config.InitDB()

	e := echo.New()

	routes.Init(e)

	e.Logger.Fatal(e.Start(":8080"))
}
