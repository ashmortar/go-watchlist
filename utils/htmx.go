package utils

import "github.com/labstack/echo/v4"

func IsHtmxRequest(h echo.Context) bool {
	return h.Request().Header.Get("HX-Request") == "true"
}

func SetHxReplaceUrl(h echo.Context, url string) {
	h.Response().Header().Set("HX-Replace-Url", url)
}
