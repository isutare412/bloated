package pkgerror

import "net/http"

type Code int

const (
	CodeUnspecified Code = 0
)

const (
	CodeBadRequest Code = 1000
)

const (
	CodeNotFound Code = 2000
)

const (
	CodeConflict = 3000
)

func (c Code) HTTPStatusCode() int {
	switch {
	case c >= 3000:
		return http.StatusConflict
	case c >= 2000:
		return http.StatusNotFound
	case c >= 1000:
		return http.StatusBadRequest
	default:
		return http.StatusInternalServerError
	}
}
