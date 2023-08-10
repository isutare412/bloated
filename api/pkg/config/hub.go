package config

import (
	"github.com/isutare412/bloated/api/pkg/core/service/todo"
	"github.com/isutare412/bloated/api/pkg/http"
	"github.com/isutare412/bloated/api/pkg/jwt"
	"github.com/isutare412/bloated/api/pkg/log"
	"github.com/isutare412/bloated/api/pkg/postgres"
)

type Hub struct {
	cfg Config
}

func NewHub(cfg Config) *Hub {
	return &Hub{cfg: cfg}
}

func (h *Hub) LogConfig() log.Config {
	return log.Config{
		Format: h.cfg.Log.Format,
		Level:  h.cfg.Log.Level,
		Caller: h.cfg.Log.Caller,
	}
}

func (h *Hub) HTTPConfig() http.Config {
	return http.Config{
		Host: h.cfg.HTTP.Host,
		Port: h.cfg.HTTP.Port,
	}
}

func (h *Hub) PostgresClientConfig() postgres.ClientConfig {
	return postgres.ClientConfig{
		Host:     h.cfg.Postgres.Host,
		Port:     h.cfg.Postgres.Port,
		User:     h.cfg.Postgres.User,
		Password: h.cfg.Postgres.Password,
		DBName:   h.cfg.Postgres.Database,
	}
}

func (h *Hub) JWTCustomClientConfig() jwt.CustomClientConfig {
	return jwt.CustomClientConfig{
		TokenTTL:       h.cfg.JWT.TokenTTL,
		PrivateKeyPath: h.cfg.JWT.PrivateKey,
		PublicKeyPath:  h.cfg.JWT.PublicKey,
	}
}

func (h *Hub) TodoServiceConfig() todo.Config {
	return todo.Config{
		MaxTodoCountPerUser: h.cfg.Service.MaxTodoCountPerUser,
	}
}
