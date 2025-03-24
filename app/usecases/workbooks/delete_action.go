package workbooks

import (
	"errors"
	"study-pal-backend/app/domains/models/entities"
	"study-pal-backend/app/domains/repositories"
	"study-pal-backend/app/usecases/shared/usecase_error"

	"github.com/google/uuid"
	"github.com/samber/lo"
)

type DeleteActionCommand struct {
	userId     uuid.UUID
	workbookId uuid.UUID
}

func NewDeleteActionCommand(userId uuid.UUID, workbookId uuid.UUID) *DeleteActionCommand {
	return &DeleteActionCommand{
		userId:     userId,
		workbookId: workbookId,
	}
}

type DeleteAction struct {
	workbookRepository repositories.WorkbookRepository
}

func NewDeleteAction(workbookRepository repositories.WorkbookRepository) *DeleteAction {
	return &DeleteAction{
		workbookRepository: workbookRepository,
	}
}

func (a *DeleteAction) Execute(command *DeleteActionCommand) usecase_error.UsecaseErrorGroup {
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
			usecase_error.NewUsecaseError(usecase_error.UnPermittedOperation, errors.New("you are not authorized to delete that workbook")),
		)
	}

	a.workbookRepository.Delete(workbook.Id())
	return nil
}
