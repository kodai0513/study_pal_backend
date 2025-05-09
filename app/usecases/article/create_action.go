package article

import (
	"study-pal-backend/app/domains/models/entities"
	"study-pal-backend/app/domains/models/value_objects/articles"
	"study-pal-backend/app/domains/repositories"
	"study-pal-backend/app/usecases/shared/trancaction"
	"study-pal-backend/app/usecases/shared/usecase_error"

	"github.com/google/uuid"
)

type CreateActionCommand struct {
	Description string
	UserId      uuid.UUID
}

type CreateAction struct {
	ArticleRepository repositories.ArticleRepository
	Tx                trancaction.Tx
}

func (c *CreateAction) Execute(command *CreateActionCommand) (*ArticleDto, usecase_error.UsecaseErrorGroup) {
	usecaseErrGroup := usecase_error.NewUsecaseErrorGroup(usecase_error.InvalidParameter)
	description, err := articles.NewDescription(command.Description)
	if err != nil {
		usecaseErrGroup.AddOnlySameUsecaseError(usecase_error.NewUsecaseError(usecase_error.InvalidParameter, err))
	}

	if usecaseErrGroup.IsError() {
		return nil, usecaseErrGroup
	}
	article := entities.NewArticle(uuid.New(), description, command.UserId)

	var resultArticle *entities.Article
	err = trancaction.WithTx(c.Tx, func() {
		resultArticle = c.ArticleRepository.Create(article)
	})

	if err != nil {
		return nil, usecase_error.NewUsecaseErrorGroupWithMessage(usecase_error.NewUsecaseError(usecase_error.DatabaseError, err))
	}

	return &ArticleDto{
			Description: resultArticle.Description(),
		},
		nil
}
