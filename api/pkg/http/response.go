package http

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/isutare412/bloated/api/pkg/contextbag"
	"github.com/isutare412/bloated/api/pkg/log"
	"github.com/isutare412/bloated/api/pkg/pkgerror"
)

const bagKeyErrorReponse = "httpErrorResponse"

type errorResponse struct {
	Message string `json:"message"`
}

func responseError(w http.ResponseWriter, r *http.Request, err error) {
	var (
		code = http.StatusInternalServerError
		msg  = err.Error()
	)

	if kerr := pkgerror.AsKnown(err); kerr != nil {
		code = kerr.Code.HTTPStatusCode()
		msg = kerr.SimpleError()
	}

	if code >= http.StatusInternalServerError {
		log.WithOperation("responseError").Error(err.Error())
	}

	errResp := errorResponse{Message: msg}
	contextbag.Bag(r.Context()).Set(bagKeyErrorReponse, errResp)

	errBytes, err := json.Marshal(&errResp)
	if err != nil {
		panic(fmt.Errorf("marshaling http error response: %w", err))
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(errBytes)
}

func responseJSON(w http.ResponseWriter, r *http.Request, body any) {
	bodyBytes, err := json.Marshal(body)
	if err != nil {
		panic(fmt.Errorf("marshaling http json response: %w", err))
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(bodyBytes)
}
