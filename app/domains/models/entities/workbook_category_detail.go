package entities

import (
	"study-pal-backend/app/domains/models/value_objects/workbook_category_details"

	"github.com/google/uuid"
)

type WorkbookCategoryDetail struct {
	id                 uuid.UUID
	name               workbook_category_details.Name
	workbookCategoryId uuid.UUID
}

func NewWorkbookCategoryDetail(
	id uuid.UUID,
	name workbook_category_details.Name,
	workbookCategoryId uuid.UUID,
) *WorkbookCategoryDetail {
	return &WorkbookCategoryDetail{
		id:                 id,
		name:               name,
		workbookCategoryId: workbookCategoryId,
	}
}

func (w *WorkbookCategoryDetail) Id() uuid.UUID {
	return w.id
}

func (w *WorkbookCategoryDetail) Name() string {
	return w.name.Value()
}

func (w *WorkbookCategoryDetail) WorkbookCategoryId() uuid.UUID {
	return w.workbookCategoryId
}

func (w *WorkbookCategoryDetail) SetName(name workbook_category_details.Name) {
	w.name = name
}
