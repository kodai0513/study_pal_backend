package workbook_members

import "study-pal-backend/app/usecases/shared/usecase_error"

type CreateActionCommand struct {
}

type CreateAction struct {
}


func (a *CreateAction) Execute(command *CreateActionCommand) (*WorkbookMemberDto, usecase_error.UsecaseErrorGroup) {
	return nil, nil
}