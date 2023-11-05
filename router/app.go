package router

import (
	"github.com/labstack/echo/v4"
)

func appRoutes(e *echo.Echo) error {
	e.GET("/dashboard", func(c echo.Context) error {
		return c.String(200, "Dashboard")
	})

	return nil
}
