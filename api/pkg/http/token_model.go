package http

import (
	"github.com/isutare412/bloated/api/pkg/core/model"
)

type verifyTokenRequest struct {
	CustomToken string `json:"customToken" validate:"required"`
}

type verifyTokenResponse struct {
	UserID     string `json:"userId"`
	Name       string `json:"name"`
	GivenName  string `json:"givenName,omitempty"`
	FamilyName string `json:"familyName,omitempty"`
	Picture    string `json:"picture,omitempty"`
	Email      string `json:"email"`
}

func (resp *verifyTokenResponse) fromCustomToken(token model.CustomToken) {
	*resp = verifyTokenResponse(token)
}

type createTokenFromGoogleRequest struct {
	Token string `json:"token" validate:"required"`
}

type createTokenRequest struct {
	Name       string `json:"name" validate:"required"`
	GivenName  string `json:"givenName"`
	FamilyName string `json:"familyName"`
	Picture    string `json:"picture"`
	Email      string `json:"email" validate:"required"`
}

func (req *createTokenRequest) toCustomToken() model.CustomToken {
	return model.CustomToken{
		Name:       req.Name,
		GivenName:  req.GivenName,
		FamilyName: req.FamilyName,
		Picture:    req.Picture,
		Email:      req.Email,
	}
}

type createTokenResponse struct {
	CustomToken string `json:"customToken"`
}
