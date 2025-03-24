package problems

import "study-pal-backend/app/usecases/shared/usecase_error"

type UpdateActionCommand struct {
}

type UpdateAction struct {
}

func (a *UpdateAction) Execute(command *UpdateActionCommand) usecase_error.UsecaseErrorGroup {
	return nil
}
