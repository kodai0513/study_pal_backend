package entities

import (
	"study-pal-backend/app/domains/models/value_objects/workbook_categories"

	"github.com/google/uuid"
)

type WorkbookCategory struct {
	id   uuid.UUID
	name workbook_categories.Name
	//problems                        []*Problem
	//workbookCategoryClassifications []*WorkbookCategoryClassification
	workbookId uuid.UUID
}

func NewWorkbookCategory(id uuid.UUID, name workbook_categories.Name, workbookId uuid.UUID) *WorkbookCategory {
	return &WorkbookCategory{
		id:         id,
		name:       name,
		workbookId: workbookId,
	}
}

func (w *WorkbookCategory) Name() string {
	return w.name.Value()
}

func (w *WorkbookCategory) WorkbookId() uuid.UUID {
	return w.workbookId
}
