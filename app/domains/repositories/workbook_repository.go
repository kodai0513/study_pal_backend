package repositories

import "study-pal-backend/app/domains/models/entities"

type WorkbookRepository interface {
	Create(workbook *entities.Workbook)
}
