package true_or_false_problems

import (
	"errors"
	"study-pal-backend/app/domains/models/value_objects/true_or_false_problems"
	"study-pal-backend/app/domains/repositories"
	"study-pal-backend/app/usecases/shared/usecase_error"

	"github.com/google/uuid"
)

type UpdateActionCommand struct {
	IsCorrect            bool
	Statement            string
	TrueOrFalseProblemId uuid.UUID
}

type UpdateAction struct {
	TrueOrFalseProblemRepository repositories.TrueOrFalseProblemRepository
}

func (a *UpdateAction) Execute(command *UpdateActionCommand) (*TrueOrFalseProblemDto, usecase_error.UsecaseErrorGroup) {
	problem := a.TrueOrFalseProblemRepository.FindById(command.TrueOrFalseProblemId)

	if problem == nil {
		return nil, usecase_error.NewUsecaseErrorGroupWithMessage(usecase_error.NewUsecaseError(usecase_error.QueryDataNotFoundError, errors.New("trueOrFalseProblem not found")))
	}

	statement, err := true_or_false_problems.NewStatement(command.Statement)
	if err != nil {
		return nil, usecase_error.NewUsecaseErrorGroupWithMessage(usecase_error.NewUsecaseError(usecase_error.InvalidParameter, errors.New("trueOrFalseProblem not found")))
	}

	problem.SetIsCorrect(command.IsCorrect)
	problem.SetStatement(statement)
	updatedProblem := a.TrueOrFalseProblemRepository.Update(problem)
	return &TrueOrFalseProblemDto{
		IsCorrect: updatedProblem.IsCorrect(),
		Statement: updatedProblem.Statement(),
	}, nil
}
