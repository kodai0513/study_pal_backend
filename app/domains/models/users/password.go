package users

import (
	"study-pal-backend/app/utils/password_hashes"

	validation "github.com/go-ozzo/ozzo-validation"
)

type Password struct {
	value string
}

func NewHashPassword(value string) (Password, error) {
	err := validation.Validate(value,
		validation.Required,
	)
	if err != nil {
		return Password{value: ""}, err
	}

	hashPassword, err := password_hashes.ConvertToHashPassword(value)
	if err != nil {
		return Password{value: ""}, err
	}

	return Password{value: hashPassword}, nil
}

func NewPassword(value string) Password {
	return Password{value: value}
}

func (p *Password) Value() string {
	return p.value
}
