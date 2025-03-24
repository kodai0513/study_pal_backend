package articles

import (
	validation "github.com/go-ozzo/ozzo-validation"
)

type Description struct {
	value string
}

func NewDescription(value string) (Description, error) {
	err := validation.Validate(value,
		validation.Required,
		validation.Length(1, 400),
	)
	if err != nil {
		return Description{value: ""}, err
	}

	return Description{value: value}, nil
}

func (d *Description) Value() string {
	return d.value
}
