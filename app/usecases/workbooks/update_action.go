package workbooks

import "study-pal-backend/app/usecases/shared/usecase_error"

type UpdateActionCommand struct {
}

func NewUpdateActionCommand() *UpdateActionCommand {
	return &UpdateActionCommand{
	}
}

type UpdateAction struct {
}

func NewUpdateAction() *UpdateAction {
	return &UpdateAction{
	}
}

func (a *UpdateAction) Execute(command *UpdateActionCommand) usecase_error.UsecaseErrorGroup {
	return nil
}