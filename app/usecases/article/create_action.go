package article

import (
	"study-pal-backend/app/domains/models/entities"
	"study-pal-backend/app/domains/models/value_objects/articles"
	"study-pal-backend/app/domains/models/value_objects/users"
	"study-pal-backend/app/domains/repositories"
	"study-pal-backend/app/usecases/shared/usecase_error"

	"github.com/google/uuid"
)

type createActionCommand struct {
	description string
	postId      uuid.UUID
}

func NewCreateActionCommand(description string, postId uuid.UUID) *createActionCommand {
	return &createActionCommand{
		description: description,
		postId:      postId,
	}
}

type CreateAction struct {
	articleRepository repositories.ArticleRepository
}

func NewCreateAction(articleRepository repositories.ArticleRepository) *CreateAction {
	return &CreateAction{
		articleRepository: articleRepository,
	}
}

func (c *CreateAction) Execute(command *createActionCommand) usecase_error.UsecaseErrorGroup {
	usecaseErrGroup := usecase_error.NewUsecaseErrorGroup(usecase_error.InvalidParameter)
	description, err := articles.NewDescription(command.description)
	if err != nil {
		usecaseErrGroup.AddOnlySameUsecaseError(usecase_error.NewUsecaseError(usecase_error.InvalidParameter, err))
	}
	userId := users.NewUserId(command.postId)

	if usecaseErrGroup.IsError() {
		return usecaseErrGroup
	}

	articleId := articles.CreateArticleId()
	article := entities.NewArticle(articleId, description, userId)
	c.articleRepository.Create(article)

	return nil
}
