package model

type CustomToken struct {
	Name       string
	GivenName  string
	FamilyName string
	Picture    string
	Email      string
}

type GoogleIDToken struct {
	Name          string
	GivenName     string
	FamilyName    string
	Picture       string
	Email         string
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
