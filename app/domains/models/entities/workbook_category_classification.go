package entities

import (
	"study-pal-backend/app/domains/models/value_objects/workbook_categories"

	"github.com/google/uuid"
)

type WorkbookCategoryClassification struct {
	id   uuid.UUID
	name workbook_categories.Name
	//problems           []*Problem
	workbookCategoryId uuid.UUID
}

func NewWorkbookCategoryClassification(
	id uuid.UUID,
	name workbook_categories.Name,
	workbookCategoryId uuid.UUID,
) *WorkbookCategoryClassification {
	return &WorkbookCategoryClassification{
		id:                 id,
		name:               name,
		workbookCategoryId: workbookCategoryId,
	}
}

func (w *WorkbookCategoryClassification) Name() string {
	return w.name.Value()
}

func (w *WorkbookCategoryClassification) WorkbookCategoryId() uuid.UUID {
	return w.workbookCategoryId
}
