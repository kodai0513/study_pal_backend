package workbook_categories

import "github.com/google/uuid"

type WorkbookCategoryDto struct {
	WorkbookCategoryId uuid.UUID
	Children           []*WorkbookCategoryDto
	Name               string
}
