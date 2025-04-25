package repositories

import (
	"study-pal-backend/app/domains/models/entities"

	"github.com/google/uuid"
)

type WorkbookRepository interface {
	Create(workbook *entities.Workbook) *entities.Workbook
	Delete(id uuid.UUID)
	ExistById(id uuid.UUID) bool
	FindById(id uuid.UUID) *entities.Workbook
	Update(workbook *entities.Workbook) *entities.Workbook
}
