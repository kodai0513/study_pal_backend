package users

import (
	"study-pal-backend/app/utils/application_errors"

	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"
)

type Name struct {
	value string
}

func NewName(value string) (*Name, application_errors.ApplicationError) {
	err := validation.Validate(value,
		validation.Required,
		validation.Length(1, 20),
		is.Alphanumeric,
	)

	if err != nil {
		return nil, application_errors.NewClientInputValidationApplicationError(err)
	}

	return &Name{value: value}, nil
}

func (n *Name) Value() string {
	return n.value
}
