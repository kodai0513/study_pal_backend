package workbook_categories

import (
	"study-pal-backend/app/domains/models/value_objects/shared/ids"

	"github.com/google/uuid"
)

type WorkbookCategoryId struct {
	value uuid.UUID
}

func CreateWorkbookCategoryId() WorkbookCategoryId {
	id := ids.CreateId()
	return WorkbookCategoryId{value: id.Value()}
}

func NewWorkbookCategoryId(value uuid.UUID) WorkbookCategoryId {
	return WorkbookCategoryId{value: value}
}

func (w *WorkbookCategoryId) Value() uuid.UUID {
	return w.value
}
