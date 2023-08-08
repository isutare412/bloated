package jwt

import (
	"crypto/rsa"
	"fmt"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"

	"github.com/isutare412/bloated/api/pkg/core/model"
)

type CustomClient struct {
	privateKey *rsa.PrivateKey
	publicKey  *rsa.PublicKey
	tokenTTL   time.Duration
}

func NewCustomClient(cfg CustomClientConfig) (*CustomClient, error) {
	privateKeyBytes, err := os.ReadFile(cfg.PrivateKeyPath)
	if err != nil {
		return nil, fmt.Errorf("reading private key %s: %w", cfg.PrivateKeyPath, err)
	}
	privateKey, err := jwt.ParseRSAPrivateKeyFromPEM(privateKeyBytes)
	if err != nil {
		return nil, fmt.Errorf("parsing private key: %w", err)
	}

	publicKeyBytes, err := os.ReadFile(cfg.PublicKeyPath)
	if err != nil {
		return nil, fmt.Errorf("reading private key %s: %w", cfg.PublicKeyPath, err)
	}
	publicKey, err := jwt.ParseRSAPublicKeyFromPEM(publicKeyBytes)
	if err != nil {
		return nil, fmt.Errorf("parsing public key: %w", err)
	}

	return &CustomClient{
		privateKey: privateKey,
		publicKey:  publicKey,
		tokenTTL:   cfg.TokenTTL,
	}, nil
}

func (c *CustomClient) SignCustomToken(customToken model.CustomToken) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodRS256, newCustomClaims(customToken, c.tokenTTL))

	tokenString, err := token.SignedString(c.privateKey)
	if err != nil {
		return "", fmt.Errorf("signing JWT token: %w", err)
	}

	return tokenString, nil
}

func (c *CustomClient) VerifyCustomToken(tokenString string) (model.CustomToken, error) {
	token, err := jwt.ParseWithClaims(
		tokenString, &customClaims{},
		func(token *jwt.Token) (any, error) {
			if _, ok := token.Method.(*jwt.SigningMethodRSA); !ok {
				return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
			}
			return c.publicKey, nil
		},
		jwt.WithIssuedAt(), jwt.WithIssuer(customIssuer))
	if err != nil {
		return model.CustomToken{}, fmt.Errorf("parsing JWT: %w", err)
	}

	if !token.Valid {
		return model.CustomToken{}, fmt.Errorf("invalid JWT token")
	}

	claims, ok := token.Claims.(*customClaims)
	if !ok {
		return model.CustomToken{}, fmt.Errorf("unexpected claims format")
	}

	return claims.toToken(), nil
}
