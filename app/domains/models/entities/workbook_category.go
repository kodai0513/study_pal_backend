package entities

import (
	"study-pal-backend/app/domains/models/value_objects/workbook_categories"

	"github.com/google/uuid"
)

type WorkbookCategory struct {
	id         uuid.UUID
	children   []*WorkbookCategory
	name       workbook_categories.Name
	workbookId uuid.UUID
}

func CreateWorkbookCategory(id uuid.UUID, name workbook_categories.Name, workbookId uuid.UUID) *WorkbookCategory {
	return &WorkbookCategory{
		id:         id,
		name:       name,
		workbookId: workbookId,
	}
}

func NewWorkbookCategory(
	id uuid.UUID,
	children []*WorkbookCategory,
	name workbook_categories.Name,
	workbookId uuid.UUID,
) *WorkbookCategory {
	return &WorkbookCategory{
		id:         id,
		children:   children,
		name:       name,
		workbookId: workbookId,
	}
}

func (w *WorkbookCategory) Id() uuid.UUID {
	return w.id
}

func (w *WorkbookCategory) Children() []*WorkbookCategory {
	return w.children
}

func (w *WorkbookCategory) Name() string {
	return w.name.Value()
}

func (w *WorkbookCategory) WorkbookId() uuid.UUID {
	return w.workbookId
}

func (w *WorkbookCategory) AddChild(workbookCategory *WorkbookCategory) {
	w.children = append(w.children, workbookCategory)
}

func (w *WorkbookCategory) SetName(name workbook_categories.Name) {
	w.name = name
}
