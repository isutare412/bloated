package validation

import (
	"errors"
	"fmt"

	"github.com/go-playground/locales/en_US"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	translations "github.com/go-playground/validator/v10/translations/en"

	"github.com/isutare412/bloated/api/pkg/pkgerror"
)

type Validator struct {
	validate   *validator.Validate
	translator ut.Translator
}

func NewValidator() (*Validator, error) {
	en := en_US.New()
	uni := ut.New(en)

	validate := validator.New()
	if err := translations.RegisterDefaultTranslations(validate, uni.GetFallback()); err != nil {
		return nil, fmt.Errorf("registering translation: %w", err)
	}

	return &Validator{
		validate:   validate,
		translator: uni.GetFallback(),
	}, nil
}

func (v *Validator) Validate(st any) error {
	if err := v.validate.Struct(st); err != nil {
		verrs := err.(validator.ValidationErrors)
		return pkgerror.Known{
			Code:   pkgerror.CodeBadRequest,
			Simple: errors.New(verrs[0].Translate(v.translator)),
		}
	}
	return nil
}
