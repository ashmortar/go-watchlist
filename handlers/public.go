package handlers

import (
	"github.com/ashmortar/go-watchlist/components"
	"github.com/ashmortar/go-watchlist/utils"
	"github.com/labstack/echo/v4"
)

const app_title = "Watchlist"

func Index(c echo.Context) error {
	return utils.Render(c, components.Page(
		app_title,
		components.Header(app_title, utils.HeaderLinks(c)),
		components.Home(),
	))
}

func Home(c echo.Context) error {
	utils.SetHxReplaceUrl(c, "/")

	content := components.Home()

	if utils.IsHtmxRequest(c) {
		return utils.Render(c, content)
	}

	return utils.Render(c, components.Page(
		app_title,
		components.Header(app_title, utils.HeaderLinks(c)),
		content,
	))
}

func Contact(c echo.Context) error {
	utils.SetHxReplaceUrl(c, "/contact")
	content := components.Contact()
	if utils.IsHtmxRequest(c) {
		return utils.Render(c, content)
	}

	return utils.Render(c, components.Page(
		app_title,
		components.Header(app_title, utils.HeaderLinks(c)),
		content,
	))
}
