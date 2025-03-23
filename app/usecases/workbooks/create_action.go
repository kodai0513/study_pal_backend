package workbooks

import (
	"study-pal-backend/app/domains/models/entities"
	"study-pal-backend/app/domains/models/value_objects/users"
	"study-pal-backend/app/domains/models/value_objects/workbooks"
	"study-pal-backend/app/domains/repositories"
	"study-pal-backend/app/usecases/shared/usecase_error"

	"github.com/google/uuid"
)

type CreateActionCommand struct {
	description string
	title       string
	userId      uuid.UUID
}

func NewCreateActionCommand(description string, title string, userId uuid.UUID) *CreateActionCommand {
	return &CreateActionCommand{
		description: description,
		title:       title,
		userId:      userId,
	}
}

type CreateAction struct {
	workbookRepository repositories.WorkbookRepository
}

func NewCreateAction(workbookRepository repositories.WorkbookRepository) *CreateAction {
	return &CreateAction{
		workbookRepository: workbookRepository,
	}
}

func (a *CreateAction) Execute(command *CreateActionCommand) usecase_error.UsecaseErrorGroup {
	usecaseErrGroup := usecase_error.NewUsecaseErrorGroup(usecase_error.InvalidParameter)
	workbookId := workbooks.CreateWorkbookId()
	description, err := workbooks.NewDescription(command.description)
	if err != nil {
		usecaseErrGroup.AddOnlySameUsecaseError(usecase_error.NewUsecaseError(usecase_error.InvalidParameter, err))
	}
	title, err := workbooks.NewTitle(command.title)
	if err != nil {
		usecaseErrGroup.AddOnlySameUsecaseError(usecase_error.NewUsecaseError(usecase_error.InvalidParameter, err))
	}
	userId := users.NewUserId(command.userId)

	if usecaseErrGroup.IsError() {
		return usecaseErrGroup
	}

	workbook := entities.CreateWorkbook(workbookId, description, userId, title)
	a.workbookRepository.Create(workbook)
	return nil
}
