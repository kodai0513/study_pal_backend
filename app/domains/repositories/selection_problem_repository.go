package repositories

import (
	"study-pal-backend/app/domains/models/entities"

	"github.com/google/uuid"
)

type SelectionProblemRepository interface {
	CreateBulk(problems []*entities.SelectionProblem) []*entities.SelectionProblem
	Delete(id uuid.UUID, workbookId uuid.UUID)
	FindByIdAndWorkbookId(id uuid.UUID, workbookId uuid.UUID) *entities.SelectionProblem
	ExistByIdAndWorkbookId(id uuid.UUID, workbookId uuid.UUID) bool
	Update(problem *entities.SelectionProblem, workbookId uuid.UUID) *entities.SelectionProblem
}
