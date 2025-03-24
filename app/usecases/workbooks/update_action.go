package workbooks

import (
	"errors"
	"study-pal-backend/app/domains/models/entities"
	"study-pal-backend/app/domains/models/value_objects/workbooks"
	"study-pal-backend/app/domains/repositories"
	"study-pal-backend/app/usecases/shared/usecase_error"

	"github.com/google/uuid"
	"github.com/samber/lo"
)

type UpdateActionCommand struct {
	description string
	title       string
	userId      uuid.UUID
	workbookId  uuid.UUID
}

func NewUpdateActionCommand(description string, title string, userId uuid.UUID, workbookId uuid.UUID) *UpdateActionCommand {
	return &UpdateActionCommand{
		description: description,
		title:       title,
		userId:      userId,
		workbookId:  workbookId,
	}
}

type UpdateAction struct {
	workbookRepository repositories.WorkbookRepository
}

func NewUpdateAction(workbookRepository repositories.WorkbookRepository) *UpdateAction {
	return &UpdateAction{
		workbookRepository: workbookRepository,
	}
}

func (a *UpdateAction) Execute(command *UpdateActionCommand) usecase_error.UsecaseErrorGroup {
	usecaseErrGroup := usecase_error.NewUsecaseErrorGroup(usecase_error.InvalidParameter)
	description, err := workbooks.NewDescription(command.description)
	if err != nil {
		usecaseErrGroup.AddOnlySameUsecaseError(usecase_error.NewUsecaseError(usecase_error.InvalidParameter, err))
	}
	title, err := workbooks.NewTitle(command.title)
	if err != nil {
		usecaseErrGroup.AddOnlySameUsecaseError(usecase_error.NewUsecaseError(usecase_error.InvalidParameter, err))
	}

	if usecaseErrGroup.IsError() {
		return usecaseErrGroup
	}

	workbook := a.workbookRepository.FindById(command.workbookId)
	if workbook == nil {
		return usecase_error.NewUsecaseErrorGroupWithMessage(usecase_error.NewUsecaseError(usecase_error.QueryDataNotFoundError, errors.New("workbook not found")))
	}

	isSelfAdminUser := len(lo.Filter(
		workbook.WorkbookMembers(),
		func(workbookMember *entities.WorkbookMember, index int) bool {
			return workbookMember.UserId() == command.userId && workbookMember.IsAdmin()
		},
	)) > 0
	if !isSelfAdminUser {
		return usecase_error.NewUsecaseErrorGroupWithMessage(
			usecase_error.NewUsecaseError(usecase_error.UnPermittedOperation, errors.New("you are not authorized to update that workbook")),
		)
	}

	workbook.SetDescription(description)
	workbook.SetTitle(title)
	a.workbookRepository.Update(workbook)

	return nil
}
