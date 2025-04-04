package repositories

import (
	"study-pal-backend/app/domains/models/entities"
)

type UserRepository interface {
	FindByName(string) *entities.User
}
