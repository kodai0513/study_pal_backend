package users

import (
	"study-pal-backend/app/utils/application_errors"
	"study-pal-backend/app/utils/validations"
)

type NickName struct {
	value string
}

func NewNickName(value string) (*NickName, application_errors.ApplicationError) {

	isCharactersMaxLength, err := validations.IsCharactersMaxLength(value, 20)
	if !isCharactersMaxLength {
		return nil, application_errors.NewClientInputValidationApplicationError(err)
	}

	return &NickName{value: value}, nil
}

func (n *NickName) Value() string {
	return n.value
}
