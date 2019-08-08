package main

import (
	"github.com/jeremylombogia/lapanganku-api/field"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

var err error

func main() {
	// Echo instance
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Routes
	e.GET("/field", field.IndexField)
	e.POST("/field", field.PostField)

	// Start server
	e.Logger.Fatal(e.Start(":1323"))
}
