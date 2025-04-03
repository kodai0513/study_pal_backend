package description_problems

import (
	"errors"
	"study-pal-backend/app/domains/repositories"
	"study-pal-backend/app/usecases/shared/usecase_error"

	"github.com/google/uuid"
)

type DeleteActionCommand struct {
	DescriptionProblemId uuid.UUID
}

type DeleteAction struct {
	DescriptionProblemRepository repositories.DescriptionProblemRepository
}

func (a *DeleteAction) Execute(command *DeleteActionCommand) usecase_error.UsecaseErrorGroup {
	existProblem := a.DescriptionProblemRepository.ExistById(command.DescriptionProblemId)
	if !existProblem {
		return usecase_error.NewUsecaseErrorGroupWithMessage(usecase_error.NewUsecaseError(usecase_error.QueryDataNotFoundError, errors.New("descriptionProblem not found")))
	}

	a.DescriptionProblemRepository.Delete(command.DescriptionProblemId)
	return nil
}
