package main

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func main() {
	e := echo.New()

	// Use middleware
	e.Use(middleware.Logger())

	// Add route
	e.GET("/", Index)
	e.GET("/app2", GetApp2)

	// Start server
	e.Logger.Fatal(e.Start(":80"))
}
