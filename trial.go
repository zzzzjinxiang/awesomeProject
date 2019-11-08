package main

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"net/http"
)

func mainx() {

	// Echo instance
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Routes
	e.GET("/", hello)
	e.GET("/go", helloWeixin)

	// Start server
	e.Logger.Fatal(e.Start(":2081"))
}

// Handler
func hello(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World!")
}

func helloWeixin(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World!wechat")
}
