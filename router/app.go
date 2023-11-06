package router

import (
	"github.com/ashmortar/go-watchlist/handlers"
	"github.com/labstack/echo/v4"
)

func appRoutes(e *echo.Echo) error {
	e.GET("/dashboard", handlers.Dashboard)

	e.GET("/lists", handlers.Lists)
	e.POST("/lists", handlers.PostList)
	e.DELETE("/lists/:id", handlers.DeleteList)

	return nil
}
