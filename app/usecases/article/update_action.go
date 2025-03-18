package article

import (
	"errors"
	"study-pal-backend/app/domains/models/articles"
	"study-pal-backend/app/domains/models/users"
	"study-pal-backend/app/domains/repositories"
	"study-pal-backend/app/usecases/shared/usecase_error"
)

type UpdateActionCommand struct {
	articleId   int
	description string
	postId      int
}

func NewUpdateActionCommand(articleId int, description string, postId int) *UpdateActionCommand {
	return &UpdateActionCommand{
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

func (c *UpdateAction) Execute(command *UpdateActionCommand) usecase_error.UsecaseErrorGroup {
	usecaseErrGroup := usecase_error.NewUsecaseErrorGroup(usecase_error.InvalidParameter)
	articleId, err := articles.NewArticleId(command.articleId)
	if err != nil {
		usecaseErrGroup.AddOnlySameUsecaseError(usecase_error.NewUsecaseError(usecase_error.InvalidParameter, err))
	}
	description, err := articles.NewDescription(command.description)
	if err != nil {
		usecaseErrGroup.AddOnlySameUsecaseError(usecase_error.NewUsecaseError(usecase_error.InvalidParameter, err))
	}
	userId, err := users.NewUserId(command.postId)
	if err != nil {
		usecaseErrGroup.AddOnlySameUsecaseError(usecase_error.NewUsecaseError(usecase_error.InvalidParameter, err))
	}

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

	article := articles.NewArticle(articleId, description, userId)
	c.articleRepository.Update(article)

	return nil
}
