package true_or_false_problems

import "study-pal-backend/app/usecases/shared/usecase_error"

type CreateActionCommand struct {
}

type CreateAction struct {
}


func (a *CreateAction) Execute(command *CreateActionCommand) (*TrueOrFalseProblemDto, usecase_error.UsecaseErrorGroup) {
	return nil, nil
}