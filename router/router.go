package router

import (
	"github.com/labstack/echo/v4"
	"github.com/rs/zerolog"
)

func InitRoutes(e *echo.Echo) error {
	publicRoutes(e)
	authRoutes(e)
	appRoutes(e)
	e.GET("/ping", func(c echo.Context) error {
		logger := c.Get("logger").(zerolog.Logger)
		logger.Info().Msg("ping")
		return c.String(200, "pong")
	})
	return nil
}
