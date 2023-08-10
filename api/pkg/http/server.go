package http

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"

	"github.com/isutare412/bloated/api/pkg/core/port"
	"github.com/isutare412/bloated/api/pkg/log"
)

type Server struct {
	server *http.Server
}

func NewServer(
	cfg Config,
	authService port.AuthService,
	todoService port.TodoService,
	ipService port.IPService,
) *Server {
	var (
		tokenInjector = newTokenInjector(authService)
		pathGetter    = newPathGetter()
		queryGetter   = newQueryGetter()
	)

	var (
		tokenHandler    = newTokenHandler(queryGetter, authService)
		todoHandler     = newTodoHandler(pathGetter, queryGetter, todoService)
		bannedIPHandler = newBannedIPHandler(ipService)
	)

	r := chi.NewRouter()
	r.Use(
		injectContextBag,
		middleware.RealIP,
		wrapResponseWriter,
		requestLogger,
		recoverPanic,
		tokenInjector.injectToken,
	)

	r.Route("/api/v1", func(r chi.Router) {
		tokenHandler.registerRoutes(r)
		todoHandler.registerRoutes(r)
		bannedIPHandler.registerRoutes(r)
	})

	s := &Server{
		server: &http.Server{
			Addr:    fmt.Sprintf("%s:%d", cfg.Host, cfg.Port),
			Handler: r,
		},
	}

	return s
}

func (s *Server) Start() <-chan error {
	runtimeErrs := make(chan error, 1)

	go func() {
		log.L().Info("Starting HTTP server", "addr", s.server.Addr)

		if err := s.server.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.L().Error("HTTP server failed to listen", "error", err)
			runtimeErrs <- err
		}
	}()

	return runtimeErrs
}

func (s *Server) Shutdown(ctx context.Context) error {
	if err := s.server.Shutdown(ctx); err != nil {
		return fmt.Errorf("shutdowning HTTP server: %w", err)
	}
	return nil
}
