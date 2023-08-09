package http

import (
	"fmt"

	"github.com/isutare412/bloated/api/pkg/core/model"
	"github.com/isutare412/bloated/api/pkg/pkgerror"
)

type createTokenFromGoogleRequest struct {
	Token string `json:"token"`
}

func (req *createTokenFromGoogleRequest) validate() error {
	if req.Token == "" {
		return pkgerror.Known{
			Code:   pkgerror.CodeBadRequest,
			Simple: fmt.Errorf("token should not be empty"),
		}
	}
	return nil
}

type createTokenRequest struct {
	Name       string `json:"name"`
	GivenName  string `json:"givenName"`
	FamilyName string `json:"familyName"`
	Picture    string `json:"picture"`
	Email      string `json:"email"`
}

func (req *createTokenRequest) validate() error {
	if req.Name == "" {
		return pkgerror.Known{
			Code:   pkgerror.CodeBadRequest,
			Simple: fmt.Errorf("name should not be empty"),
		}
	}
	if req.Email == "" {
		return pkgerror.Known{
			Code:   pkgerror.CodeBadRequest,
			Simple: fmt.Errorf("email should not be empty"),
		}
	}
	return nil
}

func (req *createTokenRequest) toCustomToken() model.CustomToken {
	return model.CustomToken(*req)
}

type createTokenResponse struct {
	CustomToken string `json:"customToken"`
}
