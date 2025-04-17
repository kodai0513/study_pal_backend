package article_likes

import (
	"errors"
	"study-pal-backend/app/domains/models/entities"
	"study-pal-backend/app/domains/repositories"
	"study-pal-backend/app/usecases/shared/trancaction"
	"study-pal-backend/app/usecases/shared/usecase_error"

	"github.com/google/uuid"
)

type CreateActionCommand struct {
	ArticleId uuid.UUID
	UserId    uuid.UUID
}

type CreateAction struct {
	ArticleRepository     repositories.ArticleRepository
	ArticleLikeRepository repositories.ArticleLikeRepository
	Tx                    trancaction.Tx
}

func (a *CreateAction) Execute(command *CreateActionCommand) usecase_error.UsecaseErrorGroup {
	existArticle := a.ArticleRepository.ExistById(command.ArticleId)
	if !existArticle {
		return usecase_error.NewUsecaseErrorGroupWithMessage(usecase_error.NewUsecaseError(usecase_error.QueryDataNotFoundError, errors.New("article not found")))
	}

	existArticleLike := a.ArticleLikeRepository.ExistByArticleIdAndUserId(command.ArticleId, command.UserId)
	if existArticleLike {
		return usecase_error.NewUsecaseErrorGroupWithMessage(usecase_error.NewUsecaseError(usecase_error.InvalidParameter, errors.New("articleLike already exists")))
	}

	trancaction.WithTx(a.Tx, func() {
		a.ArticleLikeRepository.Create(entities.NewArticleLike(uuid.New(), command.ArticleId, command.UserId))
	})

	return nil
}
