package true_or_false_problems

import (
	"errors"
	"study-pal-backend/app/domains/repositories"
	"study-pal-backend/app/usecases/shared/trancaction"
	"study-pal-backend/app/usecases/shared/usecase_error"

	"github.com/google/uuid"
)

type DeleteActionCommand struct {
	TrueOrFalseProblemId uuid.UUID
}

type DeleteAction struct {
	TrueOrFalseProblemRepository repositories.TrueOrFalseProblemRepository
	Tx                           trancaction.Tx
}

func (a *DeleteAction) Execute(command *DeleteActionCommand) usecase_error.UsecaseErrorGroup {
	existProblem := a.TrueOrFalseProblemRepository.ExistById(command.TrueOrFalseProblemId)
	if !existProblem {
		return usecase_error.NewUsecaseErrorGroupWithMessage(usecase_error.NewUsecaseError(usecase_error.QueryDataNotFoundError, errors.New("trueOrFalseProblem not found")))
	}

	err := trancaction.WithTx(a.Tx, func() {
		a.TrueOrFalseProblemRepository.Delete(command.TrueOrFalseProblemId)
	})

	if err != nil {
		return usecase_error.NewUsecaseErrorGroupWithMessage(usecase_error.NewUsecaseError(usecase_error.DatabaseError, err))
	}

	return nil
}
