package assets

import (
	"github.com/labstack/echo/v4"
)

func Serve(e *echo.Echo) {
	e.Static("/css", "assets/css")
	e.Static("/js", "assets/js")
	e.Static("/", "assets/favicon")
	e.Static("/images", "assets/images")
}
