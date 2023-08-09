package model

import (
	"fmt"

	"github.com/isutare412/bloated/api/pkg/core/ent"
	"github.com/isutare412/bloated/api/pkg/core/enum"
	"github.com/isutare412/bloated/api/pkg/pkgerror"
)

type CustomToken struct {
	UserID     string
	Name       string
	GivenName  string
	FamilyName string
	Picture    string
	Email      string
}

func (t *CustomToken) Validate() error {
	if t.Name == "" {
		return pkgerror.Known{
			Code:   pkgerror.CodeBadRequest,
			Simple: fmt.Errorf("name claim does not exist"),
		}
	}

	if t.Email == "" {
		return pkgerror.Known{
			Code:   pkgerror.CodeBadRequest,
			Simple: fmt.Errorf("email claim does not exist"),
		}
	}
	return nil
}

func (t *CustomToken) ToUser(iss enum.Issuer) *ent.User {
	return &ent.User{
		Email:      t.Email,
		UserName:   t.Name,
		GivenName:  t.GivenName,
		FamilyName: t.FamilyName,
		PhotoURL:   t.Picture,
		Origin:     iss,
	}
}

type GoogleIDToken struct {
	Name          string
	GivenName     string
	FamilyName    string
	Picture       string
	Email         string
	EmailVerified bool
}

func (t *GoogleIDToken) Validate() error {
	if t.Name == "" {
		return pkgerror.Known{
			Code:   pkgerror.CodeBadRequest,
			Simple: fmt.Errorf("name claim does not exist"),
		}
	}

	if t.Email == "" {
		return pkgerror.Known{
			Code:   pkgerror.CodeBadRequest,
			Simple: fmt.Errorf("email claim does not exist"),
		}
	}
	return nil
}

func (t *GoogleIDToken) ToCustomToken() CustomToken {
	return CustomToken{
		Name:       t.Name,
		GivenName:  t.GivenName,
		FamilyName: t.FamilyName,
		Picture:    t.Picture,
		Email:      t.Email,
	}
}
