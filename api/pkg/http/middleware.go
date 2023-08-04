package http

import (
	"net/http"
	"time"

	"github.com/go-chi/chi/v5/middleware"
	"go.uber.org/zap"

	"github.com/isutare412/bloated/api/pkg/contextbag"
	"github.com/isutare412/bloated/api/pkg/log"
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

		accessLog := log.A().With(
			zap.String("method", r.Method),
			zap.String("url", r.URL.String()),
			zap.String("addr", r.RemoteAddr),
			zap.String("proto", r.Proto),
			zap.Int64("contentLength", r.ContentLength),
			zap.String("userAgent", r.UserAgent()),
			zap.Int("status", statusCode),
			zap.Int("bodyBytes", ww.BytesWritten()),
			zap.Duration("elapsed", elapsedTime),
		)

		if ct := r.Header.Get("Content-Type"); ct != "" {
			accessLog = accessLog.With(zap.String("contentType", ct))
		}

		if errResp, ok := contextbag.Bag(r.Context()).Get(bagKeyErrorReponse).(errorResponse); ok {
			accessLog = accessLog.With(zap.String("error", errResp.Message))
		}

		accessLog.Info()
	}

	return http.HandlerFunc(fn)
}

func recoverPanic(next http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if r := recover(); r != nil {
				log.WithOperation("recoverHTTPPanic").With(zap.Stack("stackTrace")).Errorf("panicked: %v", r)
				w.WriteHeader(http.StatusInternalServerError)
			}
		}()

		next.ServeHTTP(w, r)
	}

	return http.HandlerFunc(fn)
}
