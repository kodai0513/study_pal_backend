package workbooks

import (
	"errors"
	"study-pal-backend/app/domains/repositories"
	"study-pal-backend/app/usecases/shared/permission_guard"
	"study-pal-backend/app/usecases/shared/trancaction"
	"study-pal-backend/app/usecases/shared/usecase_error"

	"github.com/google/uuid"
)

type DeleteActionCommand struct {
	UserId     uuid.UUID
	WorkbookId uuid.UUID
}

type DeleteAction struct {
	PermissionGuard    permission_guard.WorkbookPermissionGuard
	Tx                 trancaction.Tx
	WorkbookRepository repositories.WorkbookRepository
}

func (a *DeleteAction) Execute(command *DeleteActionCommand) usecase_error.UsecaseErrorGroup {
	err := a.PermissionGuard.Check("delete:workbooks", command.UserId, command.WorkbookId)
	if err != nil {
		return usecase_error.NewUsecaseErrorGroupWithMessage(usecase_error.NewUsecaseError(usecase_error.UnPermittedOperation, err))
	}
	workbook := a.WorkbookRepository.FindById(command.WorkbookId)
	if workbook == nil {
		return usecase_error.NewUsecaseErrorGroupWithMessage(usecase_error.NewUsecaseError(usecase_error.QueryDataNotFoundError, errors.New("workbook not found")))
	}

	err = trancaction.WithTx(a.Tx, func() {
		a.WorkbookRepository.Delete(workbook.Id())
	})

	if err != nil {
		return usecase_error.NewUsecaseErrorGroupWithMessage(usecase_error.NewUsecaseError(usecase_error.DatabaseError, err))
	}

	return nil
}
