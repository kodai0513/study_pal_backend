package true_or_false_problems

import (
	"errors"
	"study-pal-backend/app/domains/models/entities"
	"study-pal-backend/app/domains/models/value_objects/true_or_false_problems"
	"study-pal-backend/app/domains/repositories"
	"study-pal-backend/app/usecases/shared/permission_guard"
	"study-pal-backend/app/usecases/shared/trancaction"
	"study-pal-backend/app/usecases/shared/usecase_error"

	"github.com/google/uuid"
)

type UpdateActionCommand struct {
	IsCorrect            bool
	Statement            string
	TrueOrFalseProblemId uuid.UUID
	UserId               uuid.UUID
	WorkbookId           uuid.UUID
}

type UpdateAction struct {
	PermissionGuard              permission_guard.WorkbookPermissionGuard
	TrueOrFalseProblemRepository repositories.TrueOrFalseProblemRepository
	Tx                           trancaction.Tx
}

func (a *UpdateAction) Execute(command *UpdateActionCommand) (*TrueOrFalseProblemDto, usecase_error.UsecaseErrorGroup) {
	err := a.PermissionGuard.Check("update:true-or-false-problems", command.UserId, command.WorkbookId)
	if err != nil {
		return nil, usecase_error.NewUsecaseErrorGroupWithMessage(usecase_error.NewUsecaseError(usecase_error.UnPermittedOperation, err))
	}
	problem := a.TrueOrFalseProblemRepository.FindByIdAndWorkbookId(command.TrueOrFalseProblemId, command.WorkbookId)

	if problem == nil {
		return nil, usecase_error.NewUsecaseErrorGroupWithMessage(usecase_error.NewUsecaseError(usecase_error.QueryDataNotFoundError, errors.New("trueOrFalseProblem not found")))
	}

	statement, err := true_or_false_problems.NewStatement(command.Statement)
	if err != nil {
		return nil, usecase_error.NewUsecaseErrorGroupWithMessage(usecase_error.NewUsecaseError(usecase_error.InvalidParameter, errors.New("trueOrFalseProblem not found")))
	}

	problem.SetIsCorrect(command.IsCorrect)
	problem.SetStatement(statement)

	var updatedProblem *entities.TrueOrFalseProblem
	err = trancaction.WithTx(a.Tx, func() {
		updatedProblem = a.TrueOrFalseProblemRepository.Update(problem, command.WorkbookId)
	})

	if err != nil {
		return nil, usecase_error.NewUsecaseErrorGroupWithMessage(usecase_error.NewUsecaseError(usecase_error.DatabaseError, err))
	}

	return &TrueOrFalseProblemDto{
		IsCorrect: updatedProblem.IsCorrect(),
		Statement: updatedProblem.Statement(),
	}, nil
}
