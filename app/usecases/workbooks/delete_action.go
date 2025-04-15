package workbooks

import (
	"errors"
	"study-pal-backend/app/domains/models/entities"
	"study-pal-backend/app/domains/repositories"
	"study-pal-backend/app/usecases/shared/trancaction"
	"study-pal-backend/app/usecases/shared/usecase_error"

	"github.com/google/uuid"
	"github.com/samber/lo"
)

type DeleteActionCommand struct {
	UserId     uuid.UUID
	WorkbookId uuid.UUID
}

type DeleteAction struct {
	Tx                 trancaction.Tx
	WorkbookRepository repositories.WorkbookRepository
}

func (a *DeleteAction) Execute(command *DeleteActionCommand) usecase_error.UsecaseErrorGroup {
	workbook := a.WorkbookRepository.FindById(command.WorkbookId)
	if workbook == nil {
		return usecase_error.NewUsecaseErrorGroupWithMessage(usecase_error.NewUsecaseError(usecase_error.QueryDataNotFoundError, errors.New("workbook not found")))
	}

	isSelfAdminUser := len(lo.Filter(
		workbook.WorkbookMembers(),
		func(workbookMember *entities.WorkbookMember, index int) bool {
			return workbookMember.UserId() == command.UserId && workbookMember.IsAdmin()
		},
	)) > 0

	if !isSelfAdminUser {
		return usecase_error.NewUsecaseErrorGroupWithMessage(
			usecase_error.NewUsecaseError(usecase_error.UnPermittedOperation, errors.New("you are not authorized to delete that workbook")),
		)
	}

	err := trancaction.WithTx(a.Tx, func() {
		a.WorkbookRepository.Delete(workbook.Id())
	})

	if err != nil {
		return usecase_error.NewUsecaseErrorGroupWithMessage(usecase_error.NewUsecaseError(usecase_error.DatabaseError, err))
	}

	return nil
}
