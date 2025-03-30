package description_problems

import "study-pal-backend/app/usecases/shared/usecase_error"

type CreateActionCommand struct {
}

type CreateAction struct {
}


func (a *CreateAction) Execute(command *CreateActionCommand) (*DescriptionProblemDto, usecase_error.UsecaseErrorGroup) {
	return nil, nil
}