package workbook_members

import "study-pal-backend/app/usecases/shared/usecase_error"

type DeleteActionCommand struct {
}

type DeleteAction struct {
}

func (a *DeleteAction) Execute(command *DeleteActionCommand) usecase_error.UsecaseErrorGroup {
	return nil
}