package workbooks

import (
	"study-pal-backend/app/domains/models/entities"
	"study-pal-backend/app/domains/models/value_objects/workbooks"
	"study-pal-backend/app/domains/repositories"
	"study-pal-backend/app/usecases/shared/usecase_error"

	"github.com/google/uuid"
)

type CreateActionCommand struct {
	Description string
	Title       string
	UserId      uuid.UUID
}

type CreateAction struct {
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
	a.WorkbookRepository.Create(workbook)
	return &WorkbookDto{
			Id:          workbook.Id(),
			Description: workbook.Description(),
			IsPublic:    workbook.IsPublic(),
			Title:       workbook.Title(),
			UserId:      workbook.UserId(),
		},
		nil
}
