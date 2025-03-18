package users

import (
	"study-pal-backend/app/domains/models/shared"
)

type UserId struct {
	value int
}

func NewUserId(value int) (UserId, error) {
	id, err := shared.NewId(value)
	if err != nil {
		return UserId{value: 0}, err
	}
	return UserId{value: id.Value()}, nil
}

func (u *UserId) Value() int {
	return u.value
}
