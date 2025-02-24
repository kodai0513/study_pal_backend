package articles

import (
	"study-pal-backend/app/utils/application_errors"
	"study-pal-backend/app/utils/validations"
)

type Description struct {
	value string
}

func NewDescription(value string) (*Description, application_errors.ApplicationError) {
	isRequired, err := validations.IsRequired(value)
	if !isRequired {
		return nil, application_errors.NewClientInputValidationApplicationError(err)
	}

	isCharactersMaxLength, err := validations.IsCharactersMaxLength(value, 400)
	if !isCharactersMaxLength {
		return nil, application_errors.NewClientInputValidationApplicationError(err)
	}

	return &Description{value: value}, nil
}

func (d *Description) Value() string {
	return d.value
}
