package config

import (
	"github.com/isutare412/bloated/api/pkg/core/service/todo"
	"github.com/isutare412/bloated/api/pkg/http"
	"github.com/isutare412/bloated/api/pkg/jwt"
	"github.com/isutare412/bloated/api/pkg/log"
	"github.com/isutare412/bloated/api/pkg/postgres"
)

type Hub struct {
	Cfg Config
}

func NewHub(cfg Config) *Hub {
	return &Hub{Cfg: cfg}
}

func (h *Hub) LogConfig() log.Config {
	return log.Config{
		Format: h.Cfg.Log.Format,
		Level:  h.Cfg.Log.Level,
		Caller: h.Cfg.Log.Caller,
	}
}

func (h *Hub) HTTPConfig() http.Config {
	return http.Config{
		Host: h.Cfg.HTTP.Host,
		Port: h.Cfg.HTTP.Port,
	}
}

func (h *Hub) PostgresClientConfig() postgres.ClientConfig {
	return postgres.ClientConfig{
		Host:     h.Cfg.Postgres.Host,
		Port:     h.Cfg.Postgres.Port,
		User:     h.Cfg.Postgres.User,
		Password: h.Cfg.Postgres.Password,
		DBName:   h.Cfg.Postgres.Database,
	}
}

func (h *Hub) JWTCustomClientConfig() jwt.CustomClientConfig {
	return jwt.CustomClientConfig{
		TokenTTL:       h.Cfg.JWT.TokenTTL,
		PrivateKeyPath: h.Cfg.JWT.PrivateKey,
		PublicKeyPath:  h.Cfg.JWT.PublicKey,
	}
}

func (h *Hub) TodoServiceConfig() todo.Config {
	return todo.Config{
		MaxTodoCountPerUser: h.Cfg.Service.MaxTodoCountPerUser,
	}
}
