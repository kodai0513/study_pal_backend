package entities

import (
	"study-pal-backend/app/domains/models/value_objects/workbook_categories"
	"study-pal-backend/app/domains/models/value_objects/workbook_category_classifications"

	"github.com/google/uuid"
)

type WorkbookCategoryClassification struct {
	id                 workbook_category_classifications.WorkbookCategoryClassificationId
	name               workbook_categories.Name
	workbookCategoryId workbook_categories.WorkbookCategoryId
}

func NewWorkbookCategoryClassification(
	id workbook_category_classifications.WorkbookCategoryClassificationId,
	name workbook_categories.Name,
	workbookCategoryId workbook_categories.WorkbookCategoryId,
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
	return w.workbookCategoryId.Value()
}
