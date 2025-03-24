package repositories

import (
	"study-pal-backend/app/domains/models/entities"

	"github.com/google/uuid"
)

type ProblemRepository interface {
	CreateBulk([]*entities.Problem)
	ExistByWorkbookId(workbookId uuid.UUID) bool
}
