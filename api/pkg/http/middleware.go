package http

import (
	"context"
	"fmt"
	"log/slog"
	"net/http"
	"runtime/debug"
	"strings"
	"time"

	"github.com/go-chi/chi/v5/middleware"

	"github.com/isutare412/bloated/api/pkg/contextbag"
	"github.com/isutare412/bloated/api/pkg/core/constant"
	"github.com/isutare412/bloated/api/pkg/core/port"
	"github.com/isutare412/bloated/api/pkg/log"
	"github.com/isutare412/bloated/api/pkg/pkgerror"
)

func injectContextBag(next http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		r = r.WithContext(contextbag.WithBag(r.Context()))
		next.ServeHTTP(w, r)
	}
	return http.HandlerFunc(fn)
}

func wrapResponseWriter(next http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		ww := middleware.NewWrapResponseWriter(w, r.ProtoMajor)
		next.ServeHTTP(ww, r)
	}
	return http.HandlerFunc(fn)
}

func requestLogger(next http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		beforeServing := time.Now()
		next.ServeHTTP(w, r)
		elapsedTime := time.Since(beforeServing)
		ww := w.(middleware.WrapResponseWriter)

		var statusCode = http.StatusOK
		if sc := ww.Status(); sc != 0 {
			statusCode = sc
		}

		accessLog := log.L().With(
			slog.String("method", r.Method),
			slog.String("url", r.URL.String()),
			slog.String("addr", r.RemoteAddr),
			slog.String("proto", r.Proto),
			slog.Int64("contentLength", r.ContentLength),
			slog.String("userAgent", r.UserAgent()),
			slog.Int("status", statusCode),
			slog.Int("bodyBytes", ww.BytesWritten()),
			slog.Duration("elapsed", elapsedTime),
		)

		if ct := r.Header.Get("Content-Type"); ct != "" {
			accessLog = accessLog.With(slog.String("contentType", ct))
		}

		if errResp, ok := contextbag.Bag(r.Context()).Get(constant.BagKeyHTTPErrorReponse).(errorResponse); ok {
			accessLog = accessLog.With(slog.String("error", errResp.Message))
		}

		accessLog.Info("HTTP access")
	}

	return http.HandlerFunc(fn)
}

func recoverPanic(next http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if r := recover(); r != nil {
				log.L().Error("HTTP handler panicked", "recover", r, "stack", debug.Stack())
				w.WriteHeader(http.StatusInternalServerError)
			}
		}()

		next.ServeHTTP(w, r)
	}

	return http.HandlerFunc(fn)
}

type tokenInjector struct {
	authService port.AuthService
}

func newTokenInjector(authService port.AuthService) *tokenInjector {
	return &tokenInjector{
		authService: authService,
	}
}

func (ti *tokenInjector) injectToken(next http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		authorization := r.Header.Get("authorization")
		if authorization != "" {
			if err := ti.verifyAuthorizationHeader(r.Context(), authorization); err != nil {
				responseError(w, r, fmt.Errorf("verifying authorization header: %w", err))
				return
			}
		}

		next.ServeHTTP(w, r)
	}

	return http.HandlerFunc(fn)
}

func (ti *tokenInjector) verifyAuthorizationHeader(ctx context.Context, auth string) error {
	splitted := strings.SplitN(auth, " ", 2)
	if len(splitted) != 2 {
		return pkgerror.Known{
			Code:   pkgerror.CodeBadRequest,
			Simple: fmt.Errorf("invalid authorization header"),
		}
	}

	var (
		authType    = splitted[0]
		tokenString = splitted[1]
	)

	if authType != "Bearer" {
		return pkgerror.Known{
			Code:   pkgerror.CodeBadRequest,
			Simple: fmt.Errorf("unknown authorization type: '%s'", authType),
		}
	}

	token, err := ti.authService.VerifyCustomToken(ctx, tokenString)
	if err != nil {
		return fmt.Errorf("verifying custom token: %w", err)
	}

	contextbag.Bag(ctx).Set(constant.BagKeyCustomToken, token)
	return nil
}
