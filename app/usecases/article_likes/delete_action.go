package article_likes

import (
	"errors"
	"study-pal-backend/app/domains/repositories"
	"study-pal-backend/app/usecases/shared/trancaction"
	"study-pal-backend/app/usecases/shared/usecase_error"

	"github.com/google/uuid"
)

type DeleteActionCommand struct {
	ArticleId     uuid.UUID
	ArticleLikeId uuid.UUID
	UserId        uuid.UUID
}

type DeleteAction struct {
	ArticleRepository     repositories.ArticleRepository
	ArticleLikeRepository repositories.ArticleLikeRepository
	Tx                    trancaction.Tx
}

func (a *DeleteAction) Execute(command *DeleteActionCommand) usecase_error.UsecaseErrorGroup {
	existArticle := a.ArticleRepository.ExistById(command.ArticleId)
	if !existArticle {
		return usecase_error.NewUsecaseErrorGroupWithMessage(usecase_error.NewUsecaseError(usecase_error.QueryDataNotFoundError, errors.New("article not found")))
	}

	existArticleLike := a.ArticleLikeRepository.ExistById(command.ArticleLikeId)
	if !existArticleLike {
		return usecase_error.NewUsecaseErrorGroupWithMessage(usecase_error.NewUsecaseError(usecase_error.QueryDataNotFoundError, errors.New("articleLike not found")))
	}

	trancaction.WithTx(a.Tx, func() {
		a.ArticleLikeRepository.Delete(command.ArticleLikeId)
	})
	return nil
}
