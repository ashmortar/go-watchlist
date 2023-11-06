package router

import (
	"github.com/ashmortar/go-watchlist/handlers"
	"github.com/labstack/echo/v4"
)

func publicRoutes(e *echo.Echo) error {
	e.GET("/", handlers.Index)
	e.GET("/home", handlers.Home)
	e.GET("/contact", handlers.Contact)

	return nil
}
