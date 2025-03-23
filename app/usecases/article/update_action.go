package article

import (
	"errors"
	"study-pal-backend/app/domains/models/entities"
	"study-pal-backend/app/domains/models/value_objects/articles"
	"study-pal-backend/app/domains/models/value_objects/users"
	"study-pal-backend/app/domains/repositories"
	"study-pal-backend/app/usecases/shared/usecase_error"

	"github.com/google/uuid"
)

type updateActionCommand struct {
	articleId   uuid.UUID
	description string
	postId      uuid.UUID
}

func NewUpdateActionCommand(articleId uuid.UUID, description string, postId uuid.UUID) *updateActionCommand {
	return &updateActionCommand{
		articleId:   articleId,
		description: description,
		postId:      postId,
	}
}

type UpdateAction struct {
	articleRepository repositories.ArticleRepository
}

func NewUpdateAction(articleRepository repositories.ArticleRepository) *UpdateAction {
	return &UpdateAction{
		articleRepository: articleRepository,
	}
}

func (c *UpdateAction) Execute(command *updateActionCommand) usecase_error.UsecaseErrorGroup {
	usecaseErrGroup := usecase_error.NewUsecaseErrorGroup(usecase_error.InvalidParameter)
	articleId := articles.NewArticleId(command.articleId)
	description, err := articles.NewDescription(command.description)
	if err != nil {
		usecaseErrGroup.AddOnlySameUsecaseError(usecase_error.NewUsecaseError(usecase_error.InvalidParameter, err))
	}
	userId := users.NewUserId(command.postId)

	if usecaseErrGroup.IsError() {
		return usecaseErrGroup
	}

	targetArticle := c.articleRepository.FindById(articleId.Value())
	if targetArticle == nil {
		return usecase_error.NewUsecaseErrorGroupWithMessage(usecase_error.NewUsecaseError(usecase_error.QueryDataNotFoundError, errors.New("article not found")))
	}

	if userId.Value() != targetArticle.UserId() {
		return usecase_error.NewUsecaseErrorGroupWithMessage(
			usecase_error.NewUsecaseError(usecase_error.UnPermittedOperation, errors.New("you are not authorized to edit that article")),
		)
	}

	article := entities.NewArticle(articleId, description, userId)
	c.articleRepository.Update(article)

	return nil
}
