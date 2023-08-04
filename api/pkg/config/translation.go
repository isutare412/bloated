package config

import (
	"github.com/isutare412/bloated/api/pkg/http"
	"github.com/isutare412/bloated/api/pkg/log"
	"github.com/isutare412/bloated/api/pkg/postgres"
)

func NewLogConfig(cfg Config) log.Config {
	return log.Config{
		Development: cfg.Log.Development,
		Format:      cfg.Log.Format,
		Level:       cfg.Log.Level,
		StackTrace:  cfg.Log.StackTrace,
		Caller:      cfg.Log.Caller,
	}
}

func NewHTTPConfig(cfg Config) http.Config {
	return http.Config{
		Host: cfg.HTTP.Host,
		Port: cfg.HTTP.Port,
	}
}

func NewPostgresClientConfig(cfg Config) postgres.ClientConfig {
	return postgres.ClientConfig{
		Host:     cfg.Postgres.Host,
		Port:     cfg.Postgres.Port,
		User:     cfg.Postgres.User,
		Password: cfg.Postgres.Password,
		DBName:   cfg.Postgres.Database,
	}
}
