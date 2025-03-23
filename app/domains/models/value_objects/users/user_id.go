package users

import (
	"study-pal-backend/app/domains/models/value_objects/shared/ids"

	"github.com/google/uuid"
)

type UserId struct {
	value uuid.UUID
}

func CreateUserId() UserId {
	id := ids.CreateId()
	return UserId{value: id.Value()}
}

func NewUserId(value uuid.UUID) UserId {
	return UserId{value: value}
}

func (u *UserId) Value() uuid.UUID {
	return u.value
}
