package http

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"

	"github.com/isutare412/bloated/api/pkg/core/port"
	"github.com/isutare412/bloated/api/pkg/validation"
)

type tokenHandler struct {
	validator   *validation.Validator
	queryGetter *queryGetter
	authService port.AuthService
}

func newTokenHandler(
	validator *validation.Validator,
	queryGetter *queryGetter,
	authService port.AuthService,
) *tokenHandler {
	return &tokenHandler{
		validator:   validator,
		queryGetter: queryGetter,
		authService: authService,
	}
}

func (h *tokenHandler) registerRoutes(r chi.Router) {
	jsonContent := middleware.AllowContentType("application/json")

	r.Post("/tokens", jsonContent(http.HandlerFunc(h.createToken)).ServeHTTP)
	r.Post("/tokens/verify", jsonContent(http.HandlerFunc(h.verifyToken)).ServeHTTP)
}

func (h *tokenHandler) verifyToken(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var req verifyTokenRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		responseError(w, r, fmt.Errorf("decoding http request body: %w", err))
		return
	}
	if err := h.validator.Validate(&req); err != nil {
		responseError(w, r, fmt.Errorf("validating http request body: %w", err))
		return
	}

	customToken, err := h.authService.VerifyCustomToken(ctx, req.CustomToken)
	if err != nil {
		responseError(w, r, fmt.Errorf("verifying custom token: %w", err))
		return
	}

	var resp verifyTokenResponse
	resp.fromCustomToken(customToken)
	responseJSON(w, r, &resp)
}

func (h *tokenHandler) createToken(w http.ResponseWriter, r *http.Request) {
	switch h.queryGetter.issuer(r) {
	case issuerGoogle:
		h.createTokenFromGoogle(w, r)
		return
	}

	ctx := r.Context()

	var req createTokenRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		responseError(w, r, fmt.Errorf("decoding http request body: %w", err))
		return
	}
	if err := h.validator.Validate(&req); err != nil {
		responseError(w, r, fmt.Errorf("validating http request body: %w", err))
		return
	}

	token, err := h.authService.IssueCustomToken(ctx, req.toCustomToken())
	if err != nil {
		responseError(w, r, fmt.Errorf("issuing custom token: %w", err))
		return
	}

	resp := createTokenResponse{CustomToken: token}
	responseJSON(w, r, &resp)
}

func (h *tokenHandler) createTokenFromGoogle(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var req createTokenFromGoogleRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		responseError(w, r, fmt.Errorf("decoding http request body: %w", err))
		return
	}
	if err := h.validator.Validate(&req); err != nil {
		responseError(w, r, fmt.Errorf("validating http request body: %w", err))
		return
	}

	token, err := h.authService.IssueCustomTokenFromGoogle(ctx, req.Token)
	if err != nil {
		responseError(w, r, fmt.Errorf("issuing custom token: %w", err))
		return
	}

	resp := createTokenResponse{CustomToken: token}
	responseJSON(w, r, &resp)
}
