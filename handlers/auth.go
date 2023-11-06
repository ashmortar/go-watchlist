package handlers

import (
	"fmt"
	"net/http"
	"time"

	"github.com/ashmortar/go-watchlist/components"
	"github.com/ashmortar/go-watchlist/utils"
	"github.com/labstack/echo/v4"
	"github.com/rs/zerolog"
	"github.com/spf13/viper"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

func googleOauth2Config() *oauth2.Config {
	oauthConfig := &oauth2.Config{
		RedirectURL:  fmt.Sprintf("http://localhost:%s/auth/google/callback", viper.GetString("PORT")),
		ClientID:     viper.GetString("GOOGLE_CLIENT_ID"),
		ClientSecret: viper.GetString("GOOGLE_CLIENT_SECRET"),
		Scopes: []string{
			"https://www.googleapis.com/auth/userinfo.email",
			"https://www.googleapis.com/auth/userinfo.profile",
			"openid",
		},
		Endpoint: google.Endpoint,
	}
	return oauthConfig
}

func Profile(c echo.Context) error {
	logger := c.Get("logger").(zerolog.Logger)
	isHtmx := utils.IsHtmxRequest(c)
	currentUser, err := utils.CurrentUser(c)
	if err != nil {
		logger.Info().Msgf("user not logged in. Error was: %s", err.Error())
		oauthConfig := googleOauth2Config()
		content := components.SignInWithGoogle(
			oauthConfig.ClientID,
			oauthConfig.RedirectURL,
		)
		if isHtmx {
			return utils.Render(c, content)
		}
		return utils.Render(c, components.Page(
			"Watchlist",
			components.Header("Watchlist", utils.HeaderLinks(c)),
			content,
		))
	}
	logger.Info().Msgf("user logged in. User: %s", currentUser.Subject)
	content := components.UserAvatar(currentUser)
	if isHtmx {
		return utils.Render(c, content)
	}
	return utils.Render(c, components.Page(
		"Watchlist",
		components.Header("Watchlist", utils.HeaderLinks(c)),
		content,
	))
}

func GoogleCallback(c echo.Context) error {
	logger := c.Get("logger").(zerolog.Logger)
	logger.Info().Msg("GoogleCallback")
	jwt := c.FormValue("credential")
	user, err := utils.ParseGoogleJwtClaims(jwt)
	if err != nil {
		logger.Error().Err(err).Msgf("error parsing google jwt claims jwt. Message was: %s", err.Error())
		return c.String(500, "error parsing google jwt claims")
	}

	logger.Info().Msgf("successfully parsed jwt claims. User: %+s", user.Subject)

	c.SetCookie(&http.Cookie{
		Name:     "jwt",
		Value:    jwt,
		Expires:  time.Unix(user.ExpiresAt, 0),
		Path:     "/",
		HttpOnly: true,
		Secure:   viper.GetBool("SECURE_COOKIE"),
		SameSite: http.SameSiteStrictMode,
	})

	return c.Redirect(http.StatusFound, "/dashboard")
}

func Logout(c echo.Context) error {
	logger := c.Get("logger").(zerolog.Logger)
	logger.Info().Msg("Logout")
	c.SetCookie(&http.Cookie{
		Name:     "jwt",
		Value:    "",
		Expires:  time.Unix(0, 0),
		Path:     "/",
		HttpOnly: true,
		Secure:   viper.GetBool("SECURE_COOKIE"),
		SameSite: http.SameSiteStrictMode,
	})
	return c.Redirect(http.StatusFound, "/")
}
