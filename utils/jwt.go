package utils

import (
	"context"
	"fmt"

	"github.com/ashmortar/go-watchlist/models"
	jwt "github.com/golang-jwt/jwt/v4"
	"github.com/mitchellh/mapstructure"
)

var ctx = context.Background()

// JWT best practices: https://www.ietf.org/rfc/rfc8725.html

func validateAndGetJwtClaims(jwtStr, algorithm string, keyfunc jwt.Keyfunc) (jwt.MapClaims, error) {
	token, err := jwt.NewParser(jwt.WithValidMethods([]string{algorithm})).ParseWithClaims(jwtStr, jwt.MapClaims{}, keyfunc)
	if err != nil {
		return nil, fmt.Errorf("failed to parse JWT: %w", err)
	}

	//nolint:forcetypeassert // we specify this when we call ParseWithClaims
	claims := token.Claims.(jwt.MapClaims)
	return claims, nil
}

func ParseGoogleJwtClaims(jwtStr string) (*models.User, error) {
	kf, err := getJwksKeyfuncForIssuer(ctx, "https://accounts.google.com")
	if err != nil {
		return nil, err
	}

	claims, err := validateAndGetJwtClaims(jwtStr, "RS256", kf.Keyfunc)
	if err != nil {
		return nil, err
	}
	googleClaims := &models.User{}

	er := mapstructure.Decode(claims, googleClaims)
	if er != nil {
		return nil, er
	}

	return googleClaims, nil
}
