package article

import (
	"errors"
	"study-pal-backend/app/domains/repositories"
	"study-pal-backend/app/usecases/shared/trancaction"
	"study-pal-backend/app/usecases/shared/usecase_error"

	"github.com/google/uuid"
)

type DeleteActionCommand struct {
	UserId    uuid.UUID
	ArticleId uuid.UUID
}

type DeleteAction struct {
	ArticleRepository repositories.ArticleRepository
	Tx                trancaction.Tx
}

func (c *DeleteAction) Execute(command *DeleteActionCommand) usecase_error.UsecaseErrorGroup {
	targetArticle := c.ArticleRepository.FindById(command.ArticleId)

	if targetArticle == nil {
		return usecase_error.NewUsecaseErrorGroupWithMessage(usecase_error.NewUsecaseError(usecase_error.QueryDataNotFoundError, errors.New("article not found")))
	}

	if command.UserId != targetArticle.UserId() {
		return usecase_error.NewUsecaseErrorGroupWithMessage(
			usecase_error.NewUsecaseError(usecase_error.UnPermittedOperation, errors.New("you are not authorized to delete that article")),
		)
	}

	err := trancaction.WithTx(c.Tx, func() {
		c.ArticleRepository.Delete(command.ArticleId)
	})

	if err != nil {
		return usecase_error.NewUsecaseErrorGroupWithMessage(usecase_error.NewUsecaseError(usecase_error.DatabaseError, err))
	}

	return nil
}
