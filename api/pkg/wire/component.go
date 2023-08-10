package wire

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/isutare412/bloated/api/pkg/config"
	"github.com/isutare412/bloated/api/pkg/core/service/auth"
	"github.com/isutare412/bloated/api/pkg/core/service/ip"
	"github.com/isutare412/bloated/api/pkg/core/service/todo"
	"github.com/isutare412/bloated/api/pkg/http"
	"github.com/isutare412/bloated/api/pkg/jwt"
	"github.com/isutare412/bloated/api/pkg/log"
	"github.com/isutare412/bloated/api/pkg/postgres"
	"github.com/isutare412/bloated/api/pkg/validation"
)

type Components struct {
	pgConn   *postgres.Connection
	userRepo *postgres.UserRepository
	ipRepo   *postgres.IPRepository
	todoRepo *postgres.TodoRepository

	jwtGoogleClient *jwt.GoogleClient
	jwtCustomClient *jwt.CustomClient

	authService *auth.Service
	todoService *todo.Service
	ipService   *ip.Service

	httpServer *http.Server
}

func NewComponents(cfgHub *config.Hub, validator *validation.Validator) (*Components, error) {
	start := time.Now()
	defer func() {
		log.L().Info("Wired components", "elapsed", time.Since(start))
	}()

	pgConn, err := postgres.NewConnection(cfgHub.PostgresClientConfig())
	if err != nil {
		return nil, fmt.Errorf("creating PostgreSQL connection: %w", err)
	}

	var (
		userRepo = postgres.NewUserRepository(pgConn)
		ipRepo   = postgres.NewIPRepository(pgConn)
		todoRepo = postgres.NewTodoRepository(pgConn)
	)

	jwtGoogleClinet, err := jwt.NewGoogleClient()
	if err != nil {
		return nil, fmt.Errorf("creating JWT google client: %w", err)
	}

	jwtCustomClient, err := jwt.NewCustomClient(cfgHub.JWTCustomClientConfig())
	if err != nil {
		return nil, fmt.Errorf("creating JWT custom client: %w", err)
	}

	var (
		authService = auth.NewService(validator, jwtCustomClient, jwtGoogleClinet, userRepo)
		todoService = todo.NewService(cfgHub.TodoServiceConfig(), pgConn, todoRepo)
		ipService   = ip.NewService(pgConn, ipRepo)
	)

	httpServer := http.NewServer(cfgHub.HTTPConfig(), validator, authService, todoService, ipService)

	return &Components{
		pgConn:   pgConn,
		userRepo: userRepo,
		ipRepo:   ipRepo,
		todoRepo: todoRepo,

		jwtGoogleClient: jwtGoogleClinet,
		jwtCustomClient: jwtCustomClient,

		authService: authService,
		todoService: todoService,
		ipService:   ipService,

		httpServer: httpServer,
	}, nil
}

func (c *Components) Start() {
	httpErrs := c.httpServer.Start()

	signals := make(chan os.Signal, 1)
	signal.Notify(signals, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)

	select {
	case s := <-signals:
		log.L().Info("Received signal", "signal", s.String())
	case err := <-httpErrs:
		log.L().Error("Shutdown request from HTTP server", "error", err)
	}
}

func (c *Components) Shutdown() {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	start := time.Now()
	defer func() {
		log.L().Info("Shutdowned components", "elapsed", time.Since(start))
	}()

	if err := c.httpServer.Shutdown(ctx); err != nil {
		log.L().Error("Failed to shutdown HTTP server", "error", err)
	}

	if err := c.pgConn.Shutdown(ctx); err != nil {
		log.L().Error("Failed to shutdown PostgreSQL connection", "error", err)
	}
}
