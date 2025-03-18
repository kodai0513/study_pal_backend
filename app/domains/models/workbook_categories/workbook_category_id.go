package workbook_categories

import "study-pal-backend/app/domains/models/shared"

type WorkbookCategoryId struct {
	value int
}

func NewWorkbookCategoryId(value int) (WorkbookCategoryId, error) {
	id, err := shared.NewId(value)
	if err != nil {
		return WorkbookCategoryId{value: 0}, err
	}

	return WorkbookCategoryId{value: id.Value()}, nil
}

func (w *WorkbookCategoryId) Value() int {
	return w.value
}
