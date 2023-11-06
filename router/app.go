package router

import (
	"github.com/ashmortar/go-watchlist/handlers"
	"github.com/labstack/echo/v4"
)

func appRoutes(e *echo.Echo) error {
	e.GET("/dashboard", handlers.Dashboard)

	return nil
}
