package repositories

import (
	"study-pal-backend/app/domains/models/entities"

	"github.com/google/uuid"
)

type DescriptionProblemRepository interface {
	CreateBulk(problems []*entities.DescriptionProblem) []*entities.DescriptionProblem
	Delete(id uuid.UUID, workbookId uuid.UUID)
	ExistByIdAndWorkbookId(id uuid.UUID, workbookId uuid.UUID) bool
	ExistByWorkbookId(id uuid.UUID) bool
	FindByIdAndWorkbookId(id uuid.UUID, workbookId uuid.UUID) *entities.DescriptionProblem
	Update(problem *entities.DescriptionProblem, workbookId uuid.UUID) *entities.DescriptionProblem
}
