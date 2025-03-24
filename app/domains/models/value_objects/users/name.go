package users

import (
	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"
)

type Name struct {
	value string
}

func NewName(value string) (Name, error) {
	err := validation.Validate(value,
		validation.Required,
		validation.Length(1, 20),
		is.Alphanumeric,
	)

	if err != nil {
		return Name{value: ""}, err
	}

	return Name{value: value}, nil
}

func (n *Name) Value() string {
	return n.value
}
