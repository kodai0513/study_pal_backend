package article

import (
	"errors"
	"study-pal-backend/app/domains/models/users"
	"study-pal-backend/app/domains/repositories"
	"study-pal-backend/app/usecases/shared/usecase_error"
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

func (c *DeleteAction) Execute(command *DeleteActionCommand) usecase_error.UsecaseErrorGroup {
	usecaseErrGroup := usecase_error.NewUsecaseErrorGroup(usecase_error.InvalidParameter)
	userId, err := users.NewUserId(command.postId)
	if err != nil {
		usecaseErrGroup.AddOnlySameUsecaseError(usecase_error.NewUsecaseError(usecase_error.InvalidParameter, err))
	}

	if usecaseErrGroup.IsError() {
		return usecaseErrGroup
	}

	targetArticle := c.articleRepository.FindById(command.articleId)

	if targetArticle == nil {
		return usecase_error.NewUsecaseErrorGroupWithMessage(usecase_error.NewUsecaseError(usecase_error.QueryDataNotFoundError, errors.New("article not found")))
	}

	if userId.Value() != targetArticle.UserId() {
		return usecase_error.NewUsecaseErrorGroupWithMessage(
			usecase_error.NewUsecaseError(usecase_error.UnPermittedOperation, errors.New("you are not authorized to delete that article")),
		)
	}

	c.articleRepository.Delete(command.articleId)

	return nil
}
