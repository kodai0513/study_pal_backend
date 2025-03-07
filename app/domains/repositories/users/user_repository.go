package users

import (
	"study-pal-backend/app/domains/models/users"
)

type UserRepository interface {
	FindByName(name string) (*users.User, error)
}
