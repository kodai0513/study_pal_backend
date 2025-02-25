package users

import (
	"study-pal-backend/app/utils/application_errors"

	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"
)

type Email struct {
	value string
}

func NewEmail(value string) (*Email, application_errors.ApplicationError) {
	err := validation.Validate(value,
		validation.Required,
		is.Email,
	)
	if err != nil {
		return nil, application_errors.NewClientInputValidationApplicationError(err)
	}

	return &Email{value: value}, nil
}

func (e *Email) Value() string {
	return e.value
}
