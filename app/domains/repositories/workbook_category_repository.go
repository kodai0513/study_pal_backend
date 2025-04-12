package repositories

import (
	"study-pal-backend/app/domains/models/entities"

	"github.com/google/uuid"
)

type WorkbookCategoryRepository interface {
	UpsertAndDeleteBulk([]*entities.WorkbookCategory, uuid.UUID) []*entities.WorkbookCategory
}
