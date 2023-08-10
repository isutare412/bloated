package config

import (
	"time"

	"github.com/isutare412/bloated/api/pkg/log"
)

type Config struct {
	App         string         `mapstructure:"app" validate:"required"`
	Environment string         `mapstructure:"environment" validate:"required"`
	Version     string         `mapstructure:"version" validate:"required"`
	Log         LogConfig      `mapstructure:"log" validate:"required"`
	HTTP        HTTPConfig     `mapstructure:"http" validate:"required"`
	Postgres    PostgresConfig `mapstructure:"postgres" validate:"required"`
	JWT         JWTConfig      `mapstructure:"jwt" validate:"required"`
	Service     ServiceConfig  `mapstructure:"service" validate:"required"`
}

type LogConfig struct {
	Development bool       `mapstructure:"development"`
	Format      log.Format `mapstructure:"format" validate:"required,oneof=json text"`
	Level       log.Level  `mapstructure:"level" validate:"required,oneof=debug info warn error"`
	StackTrace  bool       `mapstructure:"stackTrace"`
	Caller      bool       `mapstructure:"caller"`
}

type HTTPConfig struct {
	Host string `mapstructure:"host" validate:"required"`
	Port int    `mapstructure:"port" validate:"required"`
}

type PostgresConfig struct {
	Host     string `mapstructure:"host" validate:"required"`
	Port     int    `mapstructure:"port" validate:"required"`
	Database string `mapstructure:"database" validate:"required"`
	User     string `mapstructure:"user" validate:"required"`
	Password string `mapstructure:"password" validate:"required"`
}

type JWTConfig struct {
	TokenTTL   time.Duration `mapstructure:"tokenTTL" validate:"required"`
	PrivateKey string        `mapstructure:"privateKey" validate:"required"`
	PublicKey  string        `mapstructure:"publicKey" validate:"required"`
}

type ServiceConfig struct {
	MaxTodoCountPerUser int `mapstructure:"maxTodoCountPerUser" validate:"required"`
}
