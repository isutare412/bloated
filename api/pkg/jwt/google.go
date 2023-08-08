package jwt

import (
	"fmt"
	"time"

	"github.com/MicahParks/keyfunc/v2"
	"github.com/golang-jwt/jwt/v5"

	"github.com/isutare412/bloated/api/pkg/core/model"
	"github.com/isutare412/bloated/api/pkg/log"
)

const (
	googleJWKsURL   = "https://www.googleapis.com/oauth2/v3/certs"
	googleJWTIssuer = "https://accounts.google.com"
)

type GoogleClient struct {
	jwks *keyfunc.JWKS
}

func NewGoogleClient() (*GoogleClient, error) {
	client := &GoogleClient{}

	jwks, err := keyfunc.Get(googleJWKsURL, keyfunc.Options{
		RefreshInterval:     10 * time.Minute,
		RefreshErrorHandler: client.logJWKRefreshError,
	})
	if err != nil {
		return nil, fmt.Errorf("creating JWKS: %w", err)
	}
	client.jwks = jwks

	return client, nil
}

func (c *GoogleClient) VerifyGoogleIDToken(tokenString string) (model.GoogleIDToken, error) {
	token, err := jwt.ParseWithClaims(
		tokenString, &googleClaims{},
		c.jwks.Keyfunc, jwt.WithIssuedAt(), jwt.WithIssuer(googleJWTIssuer))
	if err != nil {
		return model.GoogleIDToken{}, fmt.Errorf("parsing JWT: %w", err)
	}

	if !token.Valid {
		return model.GoogleIDToken{}, fmt.Errorf("invalid JWT token")
	}

	claims, ok := token.Claims.(*googleClaims)
	if !ok {
		return model.GoogleIDToken{}, fmt.Errorf("unexpected claims format")
	}

	return claims.toToken(), nil
}

func (c *GoogleClient) logJWKRefreshError(err error) {
	log.WithOperation("jwkRefresh").Errorf("Failed to refresh Google JWKS: %v", err)
}
