package description_problems

import (
	"errors"
	"study-pal-backend/app/domains/models/value_objects/description_problems"
	"study-pal-backend/app/domains/repositories"
	"study-pal-backend/app/usecases/shared/usecase_error"

	"github.com/google/uuid"
)

type UpdateActionCommand struct {
	DescriptionProblemId uuid.UUID
	CorrentStatement     string
	Statement            string
}

type UpdateAction struct {
	DescriptionProblemRepository repositories.DescriptionProblemRepository
}

func (a *UpdateAction) Execute(command *UpdateActionCommand) (*DescriptionProblemDto, usecase_error.UsecaseErrorGroup) {
	problem := a.DescriptionProblemRepository.FindById(command.DescriptionProblemId)

	if problem == nil {
		return nil, usecase_error.NewUsecaseErrorGroupWithMessage(usecase_error.NewUsecaseError(usecase_error.QueryDataNotFoundError, errors.New("descriptionProblem not found")))
	}

	usecaseErrGroup := usecase_error.NewUsecaseErrorGroup(usecase_error.InvalidParameter)
	correctStatement, err := description_problems.NewCorrectStatement(command.CorrentStatement)
	if err != nil {
		usecaseErrGroup.AddOnlySameUsecaseError(usecase_error.NewUsecaseError(usecase_error.InvalidParameter, err))
	}
	statement, err := description_problems.NewStatement(command.Statement)
	if err != nil {
		usecaseErrGroup.AddOnlySameUsecaseError(usecase_error.NewUsecaseError(usecase_error.InvalidParameter, err))
	}

	if usecaseErrGroup.IsError() {
		return nil, usecaseErrGroup
	}

	problem.SetCorrectStatement(correctStatement)
	problem.SetStatement(statement)
	updatedProblem := a.DescriptionProblemRepository.Update(problem)

	return &DescriptionProblemDto{
		CorrentStatement: updatedProblem.CorrectStatement(),
		Statement:        updatedProblem.Statement(),
	}, nil
}
