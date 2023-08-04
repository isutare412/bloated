package http

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	"go.uber.org/fx"

	"github.com/go-chi/chi/v5"

	"github.com/isutare412/bloated/api/pkg/core/port"
	"github.com/isutare412/bloated/api/pkg/log"
)

type Server struct {
	server *http.Server
}

func NewServer(
	lc fx.Lifecycle,
	shut fx.Shutdowner,
	cfg Config,
	todoSvc port.TodoService,
	ipSvc port.IPService,
) *Server {
	r := chi.NewRouter()

	s := &Server{
		server: &http.Server{
			Addr:    fmt.Sprintf("%s:%d", cfg.Host, cfg.Port),
			Handler: r,
		},
	}

	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error { return s.Start(ctx, shut) },
		OnStop:  s.Stop,
	})

	return s
}

func (s *Server) Start(ctx context.Context, shut fx.Shutdowner) error {
	go func() {
		log.WithOperation("httpStart").Infof("Starting HTTP server at %s", s.server.Addr)

		if err := s.server.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.WithOperation("httpListenAndServe").Errorf("Failed to listen: %v", err)
			shut.Shutdown()
		}
	}()
	return nil
}

func (s *Server) Stop(ctx context.Context) error {
	if err := s.server.Shutdown(ctx); err != nil {
		return fmt.Errorf("shutdowning HTTP server: %w", err)
	}
	return nil
}
