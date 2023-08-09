package http

import (
	"fmt"
	"net/http"

	"github.com/isutare412/bloated/api/pkg/pkgerror"
)

type issuer string

const (
	issuerNone   = ""
	issuerGoogle = "google"
)

type queryGetter struct{}

func newQueryGetter() *queryGetter { return &queryGetter{} }

func (q *queryGetter) userID(r *http.Request) (string, error) {
	userID := r.URL.Query().Get("userid")
	if userID == "" {
		return "", pkgerror.Known{
			Code:   pkgerror.CodeBadRequest,
			Simple: fmt.Errorf("query parameter 'userid' was not given"),
		}
	}

	return userID, nil
}

func (q *queryGetter) issuer(r *http.Request) issuer {
	return issuer(r.URL.Query().Get("issuer"))
}
