package http

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"go.uber.org/fx"

	"github.com/isutare412/bloated/api/pkg/log"
)

type Server struct {
	server *http.Server
}

func NewServer(
	lifecycle fx.Lifecycle,
	shutdowner fx.Shutdowner,
	cfg Config,
	todoHandler *todoHandler,
	bannedIPHandler *bannedIPHandler,
) *Server {
	r := chi.NewRouter()
	r.Use(
		injectContextBag,
		middleware.RealIP,
		wrapResponseWriter,
		requestLogger,
		recoverPanic,
	)

	r.Route("/api/v1", func(r chi.Router) {
		r.Mount("/todos", todoHandler.router())
		r.Mount("/banned-ips", bannedIPHandler.router())
	})

	s := &Server{
		server: &http.Server{
			Addr:    fmt.Sprintf("%s:%d", cfg.Host, cfg.Port),
			Handler: r,
		},
	}

	lifecycle.Append(fx.Hook{
		OnStart: func(ctx context.Context) error { return s.Start(ctx, shutdowner) },
		OnStop:  s.Stop,
	})

	return s
}

func (s *Server) Start(ctx context.Context, shutdowner fx.Shutdowner) error {
	go func() {
		log.WithOperation("httpStart").Infof("Starting HTTP server at %s", s.server.Addr)

		if err := s.server.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.WithOperation("httpListenAndServe").Errorf("Failed to listen: %v", err)
			shutdowner.Shutdown()
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
