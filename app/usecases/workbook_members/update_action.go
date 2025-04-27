package workbook_members

import "study-pal-backend/app/usecases/shared/usecase_error"

type UpdateActionCommand struct {
}

type UpdateAction struct {
}


func (a *UpdateAction) Execute(command *UpdateActionCommand) (*WorkbookMemberDto, usecase_error.UsecaseErrorGroup) {
	return nil, nil
}