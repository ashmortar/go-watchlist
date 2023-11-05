package utils

import (
	"context"
	"encoding/json"
	"net/http"
	"sync"
	"time"

	"github.com/MicahParks/keyfunc"
	"github.com/rs/zerolog/log"
)

var keyfuncByIssuerMutex sync.Mutex
var keyfuncByIssuer = map[string]*keyfunc.JWKS{}

func getJwksKeyfuncForIssuer(ctx context.Context, iss string) (*keyfunc.JWKS, error) {
	keyfuncByIssuerMutex.Lock()
	defer keyfuncByIssuerMutex.Unlock()

	if kf, ok := keyfuncByIssuer[iss]; ok {
		return kf, nil
	}

	jwksURI, err := fetchJwksURIForIssuer(ctx, iss)
	if err != nil {
		return nil, err
	}

	logger := log.Ctx(ctx)
	logger.Info().
		Str("iss", iss).
		Str("jwks_uri", jwksURI).
		Msg("JWKS URL for issuer")

	kf, err := keyfunc.Get(jwksURI, keyfunc.Options{
		Ctx: ctx,
		RefreshErrorHandler: func(err error) {
			logger.Err(err).
				Str("jwks_uri", jwksURI).
				Msg("JWKS keyfunc refresh error")
		},
		RefreshInterval: time.Hour,
		RefreshTimeout:  10 * time.Second,
	})
	if err != nil {
		return nil, err
	}

	keyfuncByIssuer[iss] = kf
	return kf, nil
}

func fetchJwksURIForIssuer(ctx context.Context, iss string) (string, error) {
	logger := log.Ctx(ctx)

	openIDURL := iss + "/.well-known/openid-configuration"

	logger.Info().
		Str("url", openIDURL).
		Msg("downloading JWT issuer OpenID configuration")

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, openIDURL, nil)
	if err != nil {
		return "", err
	}
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return "", err
	}
	defer func() {
		_ = resp.Body.Close()
	}()

	var bodyData struct {
		JwksURI string `json:"jwks_uri"`
	}

	err = json.NewDecoder(resp.Body).Decode(&bodyData)
	if err != nil {
		return "", err
	}

	return bodyData.JwksURI, nil
}
