package main

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/sumrid/docker_compose_example/service1/controller"
)

func main() {
	e := echo.New()

	// Use middleware
	e.Use(middleware.Logger())

	// Add route
	e.GET("/", controller.Index)
	e.GET("/app2", controller.GetApp2)

	// Start server
	e.Logger.Fatal(e.Start(":80"))
}
