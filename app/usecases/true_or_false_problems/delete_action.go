package true_or_false_problems

import (
	"errors"
	"study-pal-backend/app/domains/repositories"
	"study-pal-backend/app/usecases/shared/permission_guard"
	"study-pal-backend/app/usecases/shared/trancaction"
	"study-pal-backend/app/usecases/shared/usecase_error"

	"github.com/google/uuid"
)

type DeleteActionCommand struct {
	UserId               uuid.UUID
	TrueOrFalseProblemId uuid.UUID
	WorkbookId           uuid.UUID
}

type DeleteAction struct {
	PermissionGuard              permission_guard.WorkbookPermissionGuard
	TrueOrFalseProblemRepository repositories.TrueOrFalseProblemRepository
	Tx                           trancaction.Tx
}

func (a *DeleteAction) Execute(command *DeleteActionCommand) usecase_error.UsecaseErrorGroup {
	err := a.PermissionGuard.Check("delete:true-or-false-problems", command.UserId, command.WorkbookId)
	if err != nil {
		return usecase_error.NewUsecaseErrorGroupWithMessage(usecase_error.NewUsecaseError(usecase_error.UnPermittedOperation, err))
	}
	existProblem := a.TrueOrFalseProblemRepository.ExistByIdAndWorkbookId(command.TrueOrFalseProblemId, command.WorkbookId)
	if !existProblem {
		return usecase_error.NewUsecaseErrorGroupWithMessage(usecase_error.NewUsecaseError(usecase_error.QueryDataNotFoundError, errors.New("trueOrFalseProblem not found")))
	}

	err = trancaction.WithTx(a.Tx, func() {
		a.TrueOrFalseProblemRepository.Delete(command.TrueOrFalseProblemId, command.WorkbookId)
	})

	if err != nil {
		return usecase_error.NewUsecaseErrorGroupWithMessage(usecase_error.NewUsecaseError(usecase_error.DatabaseError, err))
	}

	return nil
}
