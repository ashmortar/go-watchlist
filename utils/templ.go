package utils

import (
	"github.com/a-h/templ"
	"github.com/labstack/echo/v4"
)

func Render(c echo.Context, content templ.Component) error {
	return content.Render(
		c.Request().Context(),
		c.Response().Writer,
	)
}
