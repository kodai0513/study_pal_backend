package selection_problems

import (
	"errors"
	"study-pal-backend/app/domains/repositories"
	"study-pal-backend/app/usecases/shared/trancaction"
	"study-pal-backend/app/usecases/shared/usecase_error"

	"github.com/google/uuid"
)

type DeleteActionCommand struct {
	SelectionProblemId uuid.UUID
}

type DeleteAction struct {
	SelectionProblemRepository repositories.SelectionProblemRepository
	Tx                         trancaction.Tx
}

func (a *DeleteAction) Execute(command *DeleteActionCommand) usecase_error.UsecaseErrorGroup {
	problemExist := a.SelectionProblemRepository.ExistById(command.SelectionProblemId)
	if !problemExist {
		return usecase_error.NewUsecaseErrorGroupWithMessage(usecase_error.NewUsecaseError(usecase_error.QueryDataNotFoundError, errors.New("selectionProblem not found")))
	}

	err := trancaction.WithTx(a.Tx, func() {
		a.SelectionProblemRepository.Delete(command.SelectionProblemId)
	})

	if err != nil {
		return usecase_error.NewUsecaseErrorGroupWithMessage(usecase_error.NewUsecaseError(usecase_error.DatabaseError, err))
	}

	return nil
}
