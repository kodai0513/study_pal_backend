package article

import (
	"errors"
	"study-pal-backend/app/domains/models/value_objects/users"
	"study-pal-backend/app/domains/repositories"
	"study-pal-backend/app/usecases/shared/usecase_error"

	"github.com/google/uuid"
)

type deleteActionCommand struct {
	postId    uuid.UUID
	articleId uuid.UUID
}

func NewDeleteActionCommand(articleId uuid.UUID, postId uuid.UUID) *deleteActionCommand {
	return &deleteActionCommand{
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

func (c *DeleteAction) Execute(command *deleteActionCommand) usecase_error.UsecaseErrorGroup {
	userId := users.NewUserId(command.postId)

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
