package handlers

import (
	"net/http"

	"github.com/ashmortar/go-watchlist/components"
	"github.com/ashmortar/go-watchlist/models"
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

func Lists(c echo.Context) error {
	user, err := utils.CurrentUser(c)
	if err != nil {
		return c.Redirect(302, "/")
	}

	lists := models.GetLists(user)

	component := components.Lists(lists)

	isHtmx := utils.IsHtmxRequest(c)
	if isHtmx {
		return utils.Render(c, component)
	}

	return utils.Render(c, components.Page(
		"Watchlist",
		components.Header("Watchlist", utils.HeaderLinks(c)),
		component,
	))
}

func PostList(c echo.Context) error {
	user, err := utils.CurrentUser(c)
	if err != nil {
		return c.Redirect(302, "/")
	}

	name := c.FormValue("name")
	list := models.CreateList(name, user)

	if utils.IsHtmxRequest(c) {
		component := components.List(list)
		return utils.Render(c, component)
	}

	return c.NoContent(http.StatusCreated)
}

func DeleteList(c echo.Context) error {
	user, err := utils.CurrentUser(c)
	if err != nil {
		return c.Redirect(302, "/")
	}

	id := c.Param("id")
	err = models.DeleteList(id, user)
	if err != nil {
		return c.NoContent(http.StatusNotFound)
	}

	if utils.IsHtmxRequest(c) {
		return c.NoContent(http.StatusOK)
	}

	return c.NoContent(http.StatusOK)
}
