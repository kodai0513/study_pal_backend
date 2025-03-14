package articles

import (
	"errors"
	"study-pal-backend/app/domains/models/articles"
	"study-pal-backend/app/domains/models/shared"
	"study-pal-backend/app/domains/repositories"
	"study-pal-backend/app/usecases/shared/usecase_errors"
)

type DeleteActionCommand struct {
	postId    int
	articleId int
}

func NewDeleteActionCommand(articleId int, postId int) *DeleteActionCommand {
	return &DeleteActionCommand{
		articleId: articleId,
		postId:    postId,
	}
}

type DeleteAction struct {
	articleRepository repositories.ArticleRepository
}

func NewDeleteAction(articleRepository repositories.ArticleRepository) *DeleteAction {
	return &DeleteAction{
		articleRepository: articleRepository,
	}
}

func (c *DeleteAction) Execute(command *DeleteActionCommand) usecase_errors.UsecaseErrorGroup {
	usecaseErrGroup := usecase_errors.NewUsecaseErrorGroup(usecase_errors.InvalidParameter)
	articleId, err := shared.NewId(command.articleId)
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
			usecase_errors.NewUsecaseError(usecase_errors.UnPermittedOperation, errors.New("you are not authorized to delete that article")),
		)
	}

	c.articleRepository.Delete(*articleId)

	return nil
}
