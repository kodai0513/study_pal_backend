package workbook_members

import (
	"study-pal-backend/app/domains/models/value_objects/shared/ids"

	"github.com/google/uuid"
)

type WorkbookMemberId struct {
	value uuid.UUID
}

func CreateWorkbookMemberId() WorkbookMemberId {
	id := ids.CreateId()
	return WorkbookMemberId{value: id.Value()}
}

func NewWorkbookMemberId(value uuid.UUID) WorkbookMemberId {
	return WorkbookMemberId{value: value}
}

func (w *WorkbookMemberId) Value() uuid.UUID {
	return w.value
}
