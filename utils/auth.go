package utils

import (
	"github.com/ashmortar/go-watchlist/models"
	"github.com/labstack/echo/v4"
)

func CurrentUser(c echo.Context) (models.User, error) {
	jwtCookie, err := c.Cookie("jwt")
	if err != nil {
		return models.User{}, err
	}

	user, err := ParseGoogleJwtClaims(jwtCookie.Value)
	if err != nil {
		return models.User{}, err
	}
	return user, nil
}

func HeaderLinks(c echo.Context) []models.HeaderLink {
	_, err := CurrentUser(c)
	if err != nil {
		return models.PublicLinks
	}
	return models.AuthLinks
}
