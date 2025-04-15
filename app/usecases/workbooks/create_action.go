package workbooks

import (
	"study-pal-backend/app/domains/models/entities"
	"study-pal-backend/app/domains/models/value_objects/workbooks"
	"study-pal-backend/app/domains/repositories"
	"study-pal-backend/app/usecases/shared/trancaction"
	"study-pal-backend/app/usecases/shared/usecase_error"

	"github.com/google/uuid"
)

type CreateActionCommand struct {
	Description string
	Title       string
	UserId      uuid.UUID
}

type CreateAction struct {
	Tx                 trancaction.Tx
	WorkbookRepository repositories.WorkbookRepository
}

func (a *CreateAction) Execute(command *CreateActionCommand) (*WorkbookDto, usecase_error.UsecaseErrorGroup) {
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

	workbook := entities.CreateWorkbook(uuid.New(), description, command.UserId, title)

	var createdWorkbook *entities.Workbook
	err = trancaction.WithTx(a.Tx, func() {
		createdWorkbook = a.WorkbookRepository.Create(workbook)
	})

	if err != nil {
		return nil, usecase_error.NewUsecaseErrorGroupWithMessage(usecase_error.NewUsecaseError(usecase_error.DatabaseError, err))
	}
	return &WorkbookDto{
			Description: createdWorkbook.Description(),
			IsPublic:    createdWorkbook.IsPublic(),
			Title:       createdWorkbook.Title(),
		},
		nil
}
