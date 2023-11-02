package handlers

import (
	"github.com/ashmortar/go-watchlist/components"
	"github.com/labstack/echo/v4"
)

func Home(c echo.Context) error {
	return components.Page().Render(c.Request().Context(), c.Response().Writer)
}
