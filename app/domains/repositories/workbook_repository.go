package repositories

import (
	"study-pal-backend/app/domains/models/entities"

	"github.com/google/uuid"
)

type WorkbookRepository interface {
	Create(*entities.Workbook) *entities.Workbook
	Delete(uuid.UUID)
	ExistById(uuid.UUID) bool
	FindById(uuid.UUID) *entities.Workbook
	Update(*entities.Workbook) *entities.Workbook
}
