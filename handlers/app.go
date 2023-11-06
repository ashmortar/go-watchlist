package handlers

import (
	"github.com/ashmortar/go-watchlist/components"
	"github.com/ashmortar/go-watchlist/utils"
	"github.com/labstack/echo/v4"
)

func Dashboard(c echo.Context) error {
	user, err := utils.CurrentUser(c)
	if err != nil {
		return c.Redirect(302, "/")
	}
	utils.SetHxReplaceUrl(c, "/dashboard")
	isHtmx := utils.IsHtmxRequest(c)

	component := components.Dashboard(user)

	if isHtmx {
		return utils.Render(c, component)
	}

	return utils.Render(c, components.Page(
		"Watchlist",
		components.Header("Watchlist", utils.HeaderLinks(c)),
		component,
	))

}
