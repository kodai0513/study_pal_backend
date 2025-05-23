package description_problems

import (
	"errors"
	"study-pal-backend/app/domains/models/entities"
	"study-pal-backend/app/domains/models/value_objects/description_problems"
	"study-pal-backend/app/domains/repositories"
	"study-pal-backend/app/usecases/shared/permission_guard"
	"study-pal-backend/app/usecases/shared/trancaction"
	"study-pal-backend/app/usecases/shared/usecase_error"

	"github.com/google/uuid"
)

type UpdateActionCommand struct {
	DescriptionProblemId uuid.UUID
	CorrentStatement     string
	Statement            string
	UserId               uuid.UUID
	WorkbookId           uuid.UUID
}

type UpdateAction struct {
	DescriptionProblemRepository repositories.DescriptionProblemRepository
	PermissionGuard              permission_guard.WorkbookPermissionGuard
	Tx                           trancaction.Tx
}

func (a *UpdateAction) Execute(command *UpdateActionCommand) (*DescriptionProblemDto, usecase_error.UsecaseErrorGroup) {
	err := a.PermissionGuard.Check("update:description-problems", command.UserId, command.WorkbookId)
	if err != nil {
		return nil, usecase_error.NewUsecaseErrorGroupWithMessage(usecase_error.NewUsecaseError(usecase_error.UnPermittedOperation, err))
	}
	problem := a.DescriptionProblemRepository.FindByIdAndWorkbookId(command.DescriptionProblemId, command.WorkbookId)
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

	var updatedProblem *entities.DescriptionProblem
	err = trancaction.WithTx(a.Tx, func() {
		updatedProblem = a.DescriptionProblemRepository.Update(problem, command.WorkbookId)
	})

	if err != nil {
		return nil, usecase_error.NewUsecaseErrorGroupWithMessage(usecase_error.NewUsecaseError(usecase_error.DatabaseError, err))
	}

	return &DescriptionProblemDto{
		CorrentStatement: updatedProblem.CorrectStatement(),
		Statement:        updatedProblem.Statement(),
	}, nil
}
