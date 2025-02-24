package users

import (
	"study-pal-backend/app/utils/application_errors"
	"study-pal-backend/app/utils/validations"
)

type Name struct {
	value string
}

func NewName(value string) (*Name, application_errors.ApplicationError) {
	isRequired, err := validations.IsRequired(value)
	if !isRequired {
		return nil, application_errors.NewClientInputValidationApplicationError(err)
	}

	isAlphanumeric, err := validations.IsAlphanumeric(value)
	if !isAlphanumeric {
		return nil, application_errors.NewClientInputValidationApplicationError(err)
	}

	isCharactersMaxLength, err := validations.IsCharactersMaxLength(value, 20)
	if !isCharactersMaxLength {
		return nil, application_errors.NewClientInputValidationApplicationError(err)
	}

	return &Name{value: value}, nil
}

func (n *Name) Value() string {
	return n.value
}
