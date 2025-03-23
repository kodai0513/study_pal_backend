package workbooks

import "study-pal-backend/app/usecases/shared/usecase_error"

type DeleteActionCommand struct {
}

func NewDeleteActionCommand() *DeleteActionCommand {
	return &DeleteActionCommand{
	}
}

type DeleteAction struct {
}

func NewDeleteAction() *DeleteAction {
	return &DeleteAction{
	}
}

func (a *DeleteAction) Execute(command *DeleteActionCommand) usecase_error.UsecaseErrorGroup {
	return nil
}