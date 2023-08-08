package http

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"

	"github.com/isutare412/bloated/api/pkg/core/port"
)

type tokenHandler struct {
	authService port.AuthService
}

func newTokenHandler(authService port.AuthService) *tokenHandler {
	return &tokenHandler{
		authService: authService,
	}
}

func (h *tokenHandler) router() http.Handler {
	jsonContent := middleware.AllowContentType("application/json")

	r := chi.NewRouter()
	r.Post("/", jsonContent(http.HandlerFunc(h.createTokenFromGoogle)).ServeHTTP)
	return r
}

func (h *tokenHandler) createTokenFromGoogle(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var req createTokenFromGoogleRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		responseError(w, r, fmt.Errorf("decoding http request body: %w", err))
		return
	}
	if err := req.validate(); err != nil {
		responseError(w, r, fmt.Errorf("validating http request body: %w", err))
		return
	}

	var (
		token string
		err   error
	)
	switch {
	case req.GoogleToken != "":
		token, err = h.authService.IssueCustomTokenFromGoogle(ctx, req.GoogleToken)
	}
	if err != nil {
		responseError(w, r, fmt.Errorf("issuing custom token: %w", err))
		return
	}

	resp := createTokenFromGoogleResponse{CustomToken: token}
	responseJSON(w, r, &resp)
}
