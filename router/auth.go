package router

import (
	"github.com/ashmortar/go-watchlist/handlers"
	"github.com/labstack/echo/v4"
)

func authRoutes(e *echo.Echo) error {

	e.POST("/auth/google/callback", handlers.GoogleCallback)

	e.GET("/logout", handlers.Logout)

	e.GET("/profile", handlers.Profile)

	return nil
}
