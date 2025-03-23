package workbooks

import (
	"study-pal-backend/app/domains/models/value_objects/shared/ids"

	"github.com/google/uuid"
)

type WorkbookId struct {
	value uuid.UUID
}

func CreateWorkbookId() WorkbookId {
	id := ids.CreateId()
	return WorkbookId{value: id.Value()}
}

func NewWorkbookId(value uuid.UUID) WorkbookId {
	return WorkbookId{value: value}
}

func (w *WorkbookId) Value() uuid.UUID {
	return w.value
}
