package jwt

import "time"

type CustomClientConfig struct {
	TokenTTL   time.Duration
	PrivateKey []byte
	PublicKey  []byte
}
