package users

import (
	validation "github.com/go-ozzo/ozzo-validation"
)

type NickName struct {
	value string
}

func NewNickName(value string) (NickName, error) {
	err := validation.Validate(value,
		validation.Length(1, 20),
	)

	if err != nil {
		return NickName{value: ""}, err
	}

	return NickName{value: value}, nil
}

func (n *NickName) Value() string {
	return n.value
}
