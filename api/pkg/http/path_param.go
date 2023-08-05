package http

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"

	"github.com/isutare412/bloated/api/pkg/pkgerror"
)

type pathGetter struct{}

func newPathGetter() *pathGetter { return &pathGetter{} }

func (p *pathGetter) todoID(r *http.Request) (int, error) {
	todoIDStr := chi.URLParam(r, "todoID")
	if todoIDStr == "" {
		return 0, fmt.Errorf("todo ID not found from URL path")
	}

	todoID, err := strconv.Atoi(todoIDStr)
	if err != nil {
		return 0, pkgerror.Known{
			Code:   pkgerror.CodeBadRequest,
			Simple: fmt.Errorf("todo ID %s is not valid", todoIDStr),
		}
	}

	return todoID, nil
}
