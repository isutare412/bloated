package http

import (
	"fmt"

	"github.com/isutare412/bloated/api/pkg/pkgerror"
)

type createTokenFromGoogleRequest struct {
	GoogleToken string `json:"googleToken"`
}

func (req *createTokenFromGoogleRequest) validate() error {
	if req.GoogleToken == "" {
		return pkgerror.Known{
			Code:   pkgerror.CodeBadRequest,
			Simple: fmt.Errorf("googleToken should not be empty"),
		}
	}
	return nil
}

type createTokenFromGoogleResponse struct {
	CustomToken string `json:"customToken"`
}
