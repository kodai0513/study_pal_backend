package users

import (
	"study-pal-backend/app/utils/application_errors"
	"study-pal-backend/app/utils/password_hashes"

	validation "github.com/go-ozzo/ozzo-validation"
)

type Password struct {
	value string
}

func NewPassword(value string) (*Password, application_errors.ApplicationError) {
	err := validation.Validate(value,
		validation.Required,
	)
	if err != nil {
		return nil, application_errors.NewClientInputValidationApplicationError(err)
	}

	hashPassword, err := password_hashes.ConvertToHashPassword(value)
	if err != nil {
		return nil, application_errors.NewClientInputValidationApplicationError(err)
	}

	return &Password{value: hashPassword}, nil
}

func (p *Password) Value() string {
	return p.value
}
