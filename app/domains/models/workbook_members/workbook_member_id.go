package workbook_members

import "study-pal-backend/app/domains/models/shared"

type WorkbookMemberId struct {
	value int
}

func NewWorkbookMemberId(value int) (WorkbookMemberId, error) {
	id, err := shared.NewId(value)
	if err != nil {
		return WorkbookMemberId{value: 0}, err
	}

	return WorkbookMemberId{value: id.Value()}, nil
}

func (w *WorkbookMemberId) Value() int {
	return w.value
}
