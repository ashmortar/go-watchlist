package middleware

import (
	"github.com/ashmortar/go-watchlist/models"
	"github.com/ashmortar/go-watchlist/utils"
	"github.com/labstack/echo/v4"
)

func SessionMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		session, err := utils.SessionStore.Get(c.Request(), "session")
		if err != nil {
			return err
		}

		if session.Values["user"] != nil {
			c.Set("user", session.Values["user"])
		} else {
			c.Set("user", models.NewUser())
		}

		return next(c)
	}
}
