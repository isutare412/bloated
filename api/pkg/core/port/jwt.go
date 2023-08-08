package port

import "github.com/isutare412/bloated/api/pkg/core/model"

//go:generate mockgen -package portmock -destination portmock/mock_jwt.go github.com/isutare412/bloated/api/pkg/core/port CustomJWTClient,GoogleJWTClient

type CustomJWTClient interface {
	VerifyCustomToken(tokenString string) (model.CustomToken, error)
	SignCustomToken(token model.CustomToken) (string, error)
}

type GoogleJWTClient interface {
	VerifyGoogleIDToken(tokenString string) (model.GoogleIDToken, error)
}
