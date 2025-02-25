package articles

import (
	"study-pal-backend/app/utils/application_errors"

	validation "github.com/go-ozzo/ozzo-validation"
)

type Description struct {
	value string
}

func NewDescription(value string) (*Description, application_errors.ApplicationError) {
	err := validation.Validate(value,
		validation.Required,
		validation.Length(1, 400),
	)
	if err != nil {
		return nil, application_errors.NewClientInputValidationApplicationError(err)
	}

	return &Description{value: value}, nil
}

func (d *Description) Value() string {
	return d.value
}
