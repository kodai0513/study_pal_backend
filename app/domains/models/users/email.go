package users

import (
	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"
)

type Email struct {
	value string
}

func NewEmail(value string) (Email, error) {
	err := validation.Validate(value,
		validation.Required,
		is.Email,
	)
	if err != nil {
		return Email{value: ""}, err
	}

	return Email{value: value}, nil
}

func (e *Email) Value() string {
	return e.value
}
