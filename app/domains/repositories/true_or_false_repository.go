package repositories

import (
	"study-pal-backend/app/domains/models/entities"

	"github.com/google/uuid"
)

type TrueOrFalseProblemRepository interface {
	CreateBulk(problems []*entities.TrueOrFalseProblem) []*entities.TrueOrFalseProblem
	Delete(id uuid.UUID, workbookId uuid.UUID)
	ExistByIdAndWorkbookId(id uuid.UUID, workbookId uuid.UUID) bool
	FindByIdAndWorkbookId(id uuid.UUID, workbookId uuid.UUID) *entities.TrueOrFalseProblem
	Update(problem *entities.TrueOrFalseProblem, workbookId uuid.UUID) *entities.TrueOrFalseProblem
}
