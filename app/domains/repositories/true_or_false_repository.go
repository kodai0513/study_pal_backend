package repositories

import (
	"study-pal-backend/app/domains/models/entities"

	"github.com/google/uuid"
)

type TrueOrFalseProblemRepository interface {
	CreateBulk([]*entities.TrueOrFalseProblem) []*entities.TrueOrFalseProblem
	Delete(uuid.UUID)
	ExistById(uuid.UUID) bool
	FindById(uuid.UUID) *entities.TrueOrFalseProblem
	Update(*entities.TrueOrFalseProblem) *entities.TrueOrFalseProblem
}
