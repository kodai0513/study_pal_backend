package workbooks

import (
	"errors"
	"study-pal-backend/app/domains/models/value_objects/workbooks"
	"study-pal-backend/app/domains/repositories"
	"study-pal-backend/app/usecases/shared/permission_guard"
	"study-pal-backend/app/usecases/shared/trancaction"
	"study-pal-backend/app/usecases/shared/usecase_error"

	"github.com/google/uuid"
)

type UpdateActionCommand struct {
	Description string
	Title       string
	UserId      uuid.UUID
	WorkbookId  uuid.UUID
}
type UpdateAction struct {
	PermissionGuard    permission_guard.WorkbookPermissionGuard
	Tx                 trancaction.Tx
	WorkbookRepository repositories.WorkbookRepository
}

func (a *UpdateAction) Execute(command *UpdateActionCommand) (*WorkbookDto, usecase_error.UsecaseErrorGroup) {
	err := a.PermissionGuard.Check("delete:workbooks", command.UserId, command.WorkbookId)
	if err != nil {
		return nil, usecase_error.NewUsecaseErrorGroupWithMessage(usecase_error.NewUsecaseError(usecase_error.UnPermittedOperation, err))
	}
	usecaseErrGroup := usecase_error.NewUsecaseErrorGroup(usecase_error.InvalidParameter)
	description, err := workbooks.NewDescription(command.Description)
	if err != nil {
		usecaseErrGroup.AddOnlySameUsecaseError(usecase_error.NewUsecaseError(usecase_error.InvalidParameter, err))
	}
	title, err := workbooks.NewTitle(command.Title)
	if err != nil {
		usecaseErrGroup.AddOnlySameUsecaseError(usecase_error.NewUsecaseError(usecase_error.InvalidParameter, err))
	}

	if usecaseErrGroup.IsError() {
		return nil, usecaseErrGroup
	}

	workbook := a.WorkbookRepository.FindById(command.WorkbookId)
	if workbook == nil {
		return nil, usecase_error.NewUsecaseErrorGroupWithMessage(usecase_error.NewUsecaseError(usecase_error.QueryDataNotFoundError, errors.New("workbook not found")))
	}

	workbook.SetDescription(description)
	workbook.SetTitle(title)

	err = trancaction.WithTx(a.Tx, func() {
		a.WorkbookRepository.Update(workbook)
	})

	if err != nil {
		return nil, usecase_error.NewUsecaseErrorGroupWithMessage(usecase_error.NewUsecaseError(usecase_error.DatabaseError, err))
	}

	return &WorkbookDto{
			Description: workbook.Description(),
			IsPublic:    workbook.IsPublic(),
			Title:       workbook.Title(),
		},
		nil
}
