package description_problems

import (
	"errors"
	"study-pal-backend/app/domains/repositories"
	"study-pal-backend/app/usecases/shared/permission_guard"
	"study-pal-backend/app/usecases/shared/trancaction"
	"study-pal-backend/app/usecases/shared/usecase_error"

	"github.com/google/uuid"
)

type DeleteActionCommand struct {
	DescriptionProblemId uuid.UUID
	UserId               uuid.UUID
	WorkbookId           uuid.UUID
}

type DeleteAction struct {
	DescriptionProblemRepository repositories.DescriptionProblemRepository
	PermissionGuard              permission_guard.WorkbookPermissionGuard
	Tx                           trancaction.Tx
}

func (a *DeleteAction) Execute(command *DeleteActionCommand) usecase_error.UsecaseErrorGroup {
	err := a.PermissionGuard.Check("delete:description-problems", command.UserId, command.WorkbookId)
	if err != nil {
		return usecase_error.NewUsecaseErrorGroupWithMessage(usecase_error.NewUsecaseError(usecase_error.UnPermittedOperation, err))
	}
	existProblem := a.DescriptionProblemRepository.ExistByIdAndWorkbookId(command.DescriptionProblemId, command.WorkbookId)
	if !existProblem {
		return usecase_error.NewUsecaseErrorGroupWithMessage(usecase_error.NewUsecaseError(usecase_error.QueryDataNotFoundError, errors.New("descriptionProblem not found")))
	}

	err = trancaction.WithTx(a.Tx, func() {
		a.DescriptionProblemRepository.Delete(command.DescriptionProblemId, command.WorkbookId)
	})

	if err != nil {
		return usecase_error.NewUsecaseErrorGroupWithMessage(usecase_error.NewUsecaseError(usecase_error.DatabaseError, err))
	}

	return nil
}
