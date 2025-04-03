package repositories

import (
	"study-pal-backend/app/domains/models/entities"

	"github.com/google/uuid"
)

type SelectionProblemRepository interface {
	CreateBulk([]*entities.SelectionProblem) []*entities.SelectionProblem
	Delete(uuid.UUID)
	FindById(uuid.UUID) *entities.SelectionProblem
	ExistById(uuid.UUID) bool
	Update(*entities.SelectionProblem) *entities.SelectionProblem
}
