package workbooks

import (
	"study-pal-backend/app/domains/models/shared"
)

type WorkbookId struct {
	value int
}

func NewWorkbookId(value int) (WorkbookId, error) {
	id, err := shared.NewId(value)
	if err != nil {
		return WorkbookId{value: 0}, err
	}
	return WorkbookId{value: id.Value()}, nil
}

func (w *WorkbookId) Value() int {

	return w.value
}
