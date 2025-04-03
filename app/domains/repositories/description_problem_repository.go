package repositories

import (
	"study-pal-backend/app/domains/models/entities"

	"github.com/google/uuid"
)

type DescriptionProblemRepository interface {
	CreateBulk([]*entities.DescriptionProblem) []*entities.DescriptionProblem
	Delete(uuid.UUID)
	ExistById(uuid.UUID) bool
	ExistByWorkbookId(uuid.UUID) bool
	FindById(uuid.UUID) *entities.DescriptionProblem
	Update(*entities.DescriptionProblem) *entities.DescriptionProblem
}
