package workbook_categories

import "study-pal-backend/app/usecases/shared/usecase_error"

type IndexActionCommand struct {
}

type IndexAction struct {
}


func (a *IndexAction) Execute(command *IndexActionCommand) (*WorkbookCategoryDto, usecase_error.UsecaseErrorGroup) {
	return nil, nil
}