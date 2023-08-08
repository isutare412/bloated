package jwt

import (
	"time"

	"github.com/golang-jwt/jwt/v5"

	"github.com/isutare412/bloated/api/pkg/core/model"
)

const (
	customIssuer = "redshore-blog"
)

type customClaims struct {
	jwt.RegisteredClaims
	Name       string `json:"name"`
	GivenName  string `json:"given_name"`
	FamilyName string `json:"family_name"`
	Picture    string `json:"picture"`
	Email      string `json:"email"`
}

func newCustomClaims(token model.CustomToken, expireAfter time.Duration) customClaims {
	now := time.Now()

	return customClaims{
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    customIssuer,
			IssuedAt:  jwt.NewNumericDate(now),
			NotBefore: jwt.NewNumericDate(now),
			ExpiresAt: jwt.NewNumericDate(now.Add(expireAfter)),
		},
		Name:       token.Name,
		GivenName:  token.GivenName,
		FamilyName: token.FamilyName,
		Picture:    token.Picture,
		Email:      token.Email,
	}
}

func (c *customClaims) toToken() model.CustomToken {
	return model.CustomToken{
		Name:       c.Name,
		GivenName:  c.GivenName,
		FamilyName: c.FamilyName,
		Picture:    c.Picture,
		Email:      c.Email,
	}
}

type googleClaims struct {
	jwt.RegisteredClaims
	Name          string `json:"name"`
	GivenName     string `json:"given_name"`
	FamilyName    string `json:"family_name"`
	Picture       string `json:"picture"`
	Email         string `json:"email"`
	EmailVerified bool   `json:"email_verified"`
}

func (c *googleClaims) toToken() model.GoogleIDToken {
	return model.GoogleIDToken{
		Name:          c.Name,
		GivenName:     c.GivenName,
		FamilyName:    c.FamilyName,
		Picture:       c.Picture,
		Email:         c.Email,
		EmailVerified: c.EmailVerified,
	}
}
