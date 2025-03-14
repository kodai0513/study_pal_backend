package articles

import (
	"study-pal-backend/app/domains/models/articles"
	"study-pal-backend/app/domains/repositories"
	"study-pal-backend/app/usecases/shared/usecase_errors"
)

type CreateActionCommand struct {
	description string
	postId      int
}

func NewCreateActionCommand(description string, postId int) *CreateActionCommand {
	return &CreateActionCommand{
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

func (c *CreateAction) Execute(command *CreateActionCommand) usecase_errors.UsecaseErrorGroup {
	usecaseErrGroup := usecase_errors.NewUsecaseErrorGroup(usecase_errors.InvalidParameter)
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

	article := articles.NewArticle(nil, description, postId)
	c.articleRepository.Save(article)

	return nil
}
