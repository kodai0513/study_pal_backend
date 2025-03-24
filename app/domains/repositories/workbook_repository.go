package repositories

import (
	"study-pal-backend/app/domains/models/entities"

	"github.com/google/uuid"
)

type WorkbookRepository interface {
	Create(workbook *entities.Workbook)
	Delete(workbookId uuid.UUID)
	FindById(workbookId uuid.UUID) *entities.Workbook
	Update(workbook *entities.Workbook)
}
