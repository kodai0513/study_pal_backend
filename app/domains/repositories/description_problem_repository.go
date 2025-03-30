package repositories

import (
	"study-pal-backend/app/domains/models/entities"

	"github.com/google/uuid"
)

type DescriptionProblemRepository interface {
	CreateBulk([]*entities.DescriptionProblem)
	ExistByWorkbookId(workbookId uuid.UUID) bool
}
