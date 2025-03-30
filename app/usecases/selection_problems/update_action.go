package selection_problems

import "study-pal-backend/app/usecases/shared/usecase_error"

type UpdateActionCommand struct {
}

type UpdateAction struct {
}


func (a *UpdateAction) Execute(command *UpdateActionCommand) (*SelectionProblemDto, usecase_error.UsecaseErrorGroup) {
	return nil, nil
}