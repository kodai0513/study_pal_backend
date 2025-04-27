package workbook_invitation_members

import "study-pal-backend/app/usecases/shared/usecase_error"

type IndexActionCommand struct {
}

type IndexAction struct {
}


func (a *IndexAction) Execute(command *IndexActionCommand) (*WorkbookInvitationMemberDto, usecase_error.UsecaseErrorGroup) {
	return nil, nil
}