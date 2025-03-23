package workbook_category_classifications

import (
	"study-pal-backend/app/domains/models/value_objects/shared/ids"

	"github.com/google/uuid"
)

type WorkbookCategoryClassificationId struct {
	value uuid.UUID
}

func CreateWorkbookCategoryClassificationId() WorkbookCategoryClassificationId {
	id := ids.CreateId()
	return WorkbookCategoryClassificationId{value: id.Value()}
}

func NewWorkbookCategoryClassificationId(value uuid.UUID) WorkbookCategoryClassificationId {
	return WorkbookCategoryClassificationId{value: value}
}

func (w *WorkbookCategoryClassificationId) Value() uuid.UUID {
	return w.value
}
