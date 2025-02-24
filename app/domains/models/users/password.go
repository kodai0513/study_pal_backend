package users

import (
	"study-pal-backend/app/utils/application_errors"
	"study-pal-backend/app/utils/password_hashes"
	"study-pal-backend/app/utils/validations"
)

type Password struct {
	value string
}

func NewPassword(value string) (*Password, application_errors.ApplicationError) {
	isRequired, err := validations.IsRequired(value)
	if !isRequired {
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
