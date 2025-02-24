package users

import (
	"study-pal-backend/app/utils/application_errors"
	"study-pal-backend/app/utils/validations"
)

type Email struct {
	value string
}

func NewEmail(value string) (*Email, application_errors.ApplicationError) {
	isRequired, err := validations.IsRequired(value)
	if !isRequired {
		return nil, application_errors.NewClientInputValidationApplicationError(err)
	}

	isEmail, err := validations.IsEmailValid(value)
	if !isEmail {
		return nil, application_errors.NewClientInputValidationApplicationError(err)
	}

	return &Email{value: value}, nil
}

func (e *Email) Value() string {
	return e.value
}
