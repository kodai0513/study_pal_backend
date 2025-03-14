package articles

import (
	"errors"
	"study-pal-backend/app/domains/models/articles"
	"study-pal-backend/app/domains/models/shared"
	"study-pal-backend/app/domains/repositories"
	"study-pal-backend/app/usecases/shared/usecase_errors"
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

func (c *UpdateAction) Execute(command *UpdateActionCommand) usecase_errors.UsecaseErrorGroup {
	usecaseErrGroup := usecase_errors.NewUsecaseErrorGroup(usecase_errors.InvalidParameter)
	articleId, err := shared.NewId(command.articleId)
	if err != nil {
		usecaseErrGroup.AddOnlySameUsecaseError(usecase_errors.NewUsecaseError(usecase_errors.InvalidParameter, err))
	}
	description, err := articles.NewDescription(command.description)
	if err != nil {
		usecaseErrGroup.AddOnlySameUsecaseError(usecase_errors.NewUsecaseError(usecase_errors.InvalidParameter, err))
	}
	postId, err := articles.NewPostId(command.postId)
	if err != nil {
		usecaseErrGroup.AddOnlySameUsecaseError(usecase_errors.NewUsecaseError(usecase_errors.InvalidParameter, err))
	}

	if usecaseErrGroup.IsError() {
		return usecaseErrGroup
	}

	targetArticle := c.articleRepository.FindById(*articleId)
	if targetArticle == nil {
		return usecase_errors.NewUsecaseErrorGroupWithMessage(usecase_errors.NewUsecaseError(usecase_errors.QueryDataNotFoundError, errors.New("article not found")))
	}

	if postId.Value() != targetArticle.PostId() {
		return usecase_errors.NewUsecaseErrorGroupWithMessage(
			usecase_errors.NewUsecaseError(usecase_errors.UnPermittedOperation, errors.New("you are not authorized to edit that article")),
		)
	}

	article := articles.NewArticle(articleId, description, postId)
	c.articleRepository.Update(article)

	return nil
}
