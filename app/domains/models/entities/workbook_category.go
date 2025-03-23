package entities

import (
	"study-pal-backend/app/domains/models/value_objects/workbook_categories"
	"study-pal-backend/app/domains/models/value_objects/workbooks"

	"github.com/google/uuid"
)

type WorkbookCategory struct {
	id         workbook_categories.WorkbookCategoryId
	name       workbook_categories.Name
	workbookId workbooks.WorkbookId
}

func NewWorkbookCategory(id workbook_categories.WorkbookCategoryId, name workbook_categories.Name, workbookId workbooks.WorkbookId) *WorkbookCategory {
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
	return w.workbookId.Value()
}
