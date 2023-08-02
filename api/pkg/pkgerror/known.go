package pkgerror

import (
	"errors"
	"fmt"
)

type Known struct {
	Code   Code
	Origin error
	Simple error
}

func (k Known) Error() string {
	switch {
	case k.Origin != nil:
		return k.Origin.Error()
	case k.Simple != nil:
		return k.Simple.Error()
	}
	return k.errorCodeMessage()
}

func (k Known) SimpleError() string {
	if k.Simple != nil {
		return k.Simple.Error()
	}
	return k.errorCodeMessage()
}

func (k Known) Unwrap() error {
	return k.Origin
}

func (k Known) errorCodeMessage() string {
	return fmt.Sprintf("known error code (%d)", k.Code)
}

func AsKnown(err error) *Known {
	var kerr Known
	if !errors.As(err, &kerr) {
		return nil
	}
	return &kerr
}
