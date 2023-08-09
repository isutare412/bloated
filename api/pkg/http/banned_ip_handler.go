package http

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"

	"github.com/isutare412/bloated/api/pkg/core/port"
)

type bannedIPHandler struct {
	ipService port.IPService
}

func newBannedIPHandler(ipService port.IPService) *bannedIPHandler {
	return &bannedIPHandler{
		ipService: ipService,
	}
}

func (h *bannedIPHandler) registerRoutes(r chi.Router) {
	r.Get("/banned-ips", h.listBannedIPs)
}

func (h *bannedIPHandler) listBannedIPs(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	bannedIPs, err := h.ipService.AllBannedIPs(ctx)
	if err != nil {
		responseError(w, r, fmt.Errorf("getting all banned IPs: %w", err))
		return
	}

	var resp listBannedIPsResponse
	resp.fromEntities(bannedIPs)
	responseJSON(w, r, &resp)
}
