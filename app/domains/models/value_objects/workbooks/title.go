package workbooks

import (
	validation "github.com/go-ozzo/ozzo-validation"
)

type Title struct {
	value string
}

func NewTitle(value string) (Title, error) {
	err := validation.Validate(value,
		validation.Required,
		validation.Length(1, 100),
	)

	if err != nil {
		return Title{value: ""}, err
	}

	return Title{value: value}, nil
}

func (t *Title) Value() string {
	return t.value
}
