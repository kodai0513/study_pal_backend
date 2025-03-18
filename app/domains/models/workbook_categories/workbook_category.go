package workbook_categories

import "study-pal-backend/app/domains/models/workbooks"

type WorkbookCategory struct {
	name       Name
	workbookId workbooks.WorkbookId
}

func NewWorkbookCategory(name Name, workbookId workbooks.WorkbookId) *WorkbookCategory {
	return &WorkbookCategory{
		name:       name,
		workbookId: workbookId,
	}
}

func (w *WorkbookCategory) Name() string {
	return w.name.Value()
}

func (w *WorkbookCategory) WorkbookId() int {
	return w.workbookId.Value()
}
