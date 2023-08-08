package jwt

import "time"

type CustomClientConfig struct {
	TokenTTL       time.Duration
	PrivateKeyPath string
	PublicKeyPath  string
}
