package model

import (
	"github.com/isutare412/bloated/api/pkg/core/constant"
	"github.com/isutare412/bloated/api/pkg/core/ent"
)

type CustomToken struct {
	UserID     string
	Name       string `validate:"required"`
	GivenName  string
	FamilyName string
	Picture    string
	Email      string `validate:"required"`
}

func (t *CustomToken) ToUser(iss constant.Issuer) *ent.User {
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
	Name          string `validate:"required"`
	GivenName     string
	FamilyName    string
	Picture       string
	Email         string `validate:"required"`
	EmailVerified bool
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
