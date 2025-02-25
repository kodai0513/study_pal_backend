package users

import (
	"study-pal-backend/app/utils/application_errors"

	validation "github.com/go-ozzo/ozzo-validation"
)

type NickName struct {
	value string
}

func NewNickName(value string) (*NickName, application_errors.ApplicationError) {
	err := validation.Validate(value,
		validation.Length(1, 20),
	)

	if err != nil {
		return nil, application_errors.NewClientInputValidationApplicationError(err)
	}

	return &NickName{value: value}, nil
}

func (n *NickName) Value() string {
	return n.value
}
