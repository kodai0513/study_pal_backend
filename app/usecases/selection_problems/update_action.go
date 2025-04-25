package selection_problems

import (
	"errors"
	"study-pal-backend/app/domains/models/entities"
	"study-pal-backend/app/domains/models/value_objects/selection_problem_answers"
	"study-pal-backend/app/domains/models/value_objects/selection_problems"
	"study-pal-backend/app/domains/repositories"
	"study-pal-backend/app/usecases/shared/permission_guard"
	"study-pal-backend/app/usecases/shared/trancaction"
	"study-pal-backend/app/usecases/shared/usecase_error"

	"github.com/google/uuid"
	"github.com/samber/lo"
)

type SelectionProblemAnswer struct {
	IsCorrect bool
	Statement string
}

type UpdateActionCommand struct {
	SelectionProblemAnswers []*SelectionProblemAnswer
	SelectionProblemId      uuid.UUID
	Statement               string
	UserId                  uuid.UUID
	WorkbookId              uuid.UUID
}

type UpdateAction struct {
	PermissionGuard            permission_guard.WorkbookPermissionGuard
	SelectionProblemRepository repositories.SelectionProblemRepository
	Tx                         trancaction.Tx
}

func (a *UpdateAction) Execute(command *UpdateActionCommand) (*SelectionProblemDto, usecase_error.UsecaseErrorGroup) {
	err := a.PermissionGuard.Check("update:selection-problems", command.UserId, command.WorkbookId)
	if err != nil {
		return nil, usecase_error.NewUsecaseErrorGroupWithMessage(usecase_error.NewUsecaseError(usecase_error.UnPermittedOperation, err))
	}
	problem := a.SelectionProblemRepository.FindByIdAndWorkbookId(command.SelectionProblemId, command.WorkbookId)

	if problem == nil {
		return nil, usecase_error.NewUsecaseErrorGroupWithMessage(usecase_error.NewUsecaseError(usecase_error.QueryDataNotFoundError, errors.New("selectionProblem not found")))
	}

	invalidUsecaseErrGroup := usecase_error.NewUsecaseErrorGroup(usecase_error.InvalidParameter)
	statement, err := selection_problems.NewStatement(command.Statement)
	if err != nil {
		invalidUsecaseErrGroup.AddOnlySameUsecaseError(usecase_error.NewUsecaseError(usecase_error.InvalidParameter, err))
	}
	problem.SetStatement(statement)

	answerEntities := lo.Map(command.SelectionProblemAnswers, func(answer *SelectionProblemAnswer, _ int) *entities.SelectionProblemAnswer {
		statement, err := selection_problem_answers.NewStatement(answer.Statement)
		if err != nil {
			invalidUsecaseErrGroup.AddOnlySameUsecaseError(usecase_error.NewUsecaseError(usecase_error.InvalidParameter, err))
		}
		return entities.NewSelectionProblemAnswer(
			uuid.New(),
			answer.IsCorrect,
			problem.Id(),
			statement,
		)
	})

	err = problem.ReplaceSelectionProblemAnswer(answerEntities)
	if err != nil {
		invalidUsecaseErrGroup.AddOnlySameUsecaseError(usecase_error.NewUsecaseError(usecase_error.InvalidParameter, err))
	}

	if invalidUsecaseErrGroup.IsError() {
		return nil, invalidUsecaseErrGroup
	}

	var updatedProblem *entities.SelectionProblem
	err = trancaction.WithTx(a.Tx, func() {
		updatedProblem = a.SelectionProblemRepository.Update(problem, command.WorkbookId)
	})

	if err != nil {
		return nil, usecase_error.NewUsecaseErrorGroupWithMessage(usecase_error.NewUsecaseError(usecase_error.DatabaseError, err))
	}

	answerDtos := lo.Map(updatedProblem.SelectionProblemAnswers(), func(answer *entities.SelectionProblemAnswer, _ int) *SelectionProblemAnswerDto {
		return &SelectionProblemAnswerDto{
			IsCorrect:                answer.IsCorrect(),
			SelectionProblemAnswerId: answer.Id(),
			Statement:                answer.Statement(),
		}
	})

	return &SelectionProblemDto{
		SelectionProblemAnswers: answerDtos,
		Statement:               updatedProblem.Statement(),
	}, nil
}
