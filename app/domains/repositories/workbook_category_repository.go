package repositories

import (
	"study-pal-backend/app/domains/models/entities"

	"github.com/google/uuid"
)

type WorkbookCategoryRepository interface {
	FindByWorkbookId(workbookId uuid.UUID) []*entities.WorkbookCategory
	UpsertAndDeleteBulk(categories []*entities.WorkbookCategory, workbookId uuid.UUID) []*entities.WorkbookCategory
}
