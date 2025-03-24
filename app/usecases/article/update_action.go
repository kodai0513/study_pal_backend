package article

import (
	"errors"
	"study-pal-backend/app/domains/models/entities"
	"study-pal-backend/app/domains/models/value_objects/articles"
	"study-pal-backend/app/domains/repositories"
	"study-pal-backend/app/usecases/shared/usecase_error"

	"github.com/google/uuid"
)

type UpdateActionCommand struct {
	ArticleId   uuid.UUID
	Description string
	UserId      uuid.UUID
}

type UpdateAction struct {
	ArticleRepository repositories.ArticleRepository
}

func (c *UpdateAction) Execute(command *UpdateActionCommand) (*ArticleDto, usecase_error.UsecaseErrorGroup) {
	usecaseErrGroup := usecase_error.NewUsecaseErrorGroup(usecase_error.InvalidParameter)
	description, err := articles.NewDescription(command.Description)
	if err != nil {
		usecaseErrGroup.AddOnlySameUsecaseError(usecase_error.NewUsecaseError(usecase_error.InvalidParameter, err))
	}

	if usecaseErrGroup.IsError() {
		return nil, usecaseErrGroup
	}

	targetArticle := c.ArticleRepository.FindById(command.ArticleId)
	if targetArticle == nil {
		return nil, usecase_error.NewUsecaseErrorGroupWithMessage(usecase_error.NewUsecaseError(usecase_error.QueryDataNotFoundError, errors.New("article not found")))
	}

	if command.UserId != targetArticle.UserId() {
		return nil, usecase_error.NewUsecaseErrorGroupWithMessage(
			usecase_error.NewUsecaseError(usecase_error.UnPermittedOperation, errors.New("you are not authorized to edit that article")),
		)
	}

	article := entities.NewArticle(command.ArticleId, description, command.UserId)
	resultArticle := c.ArticleRepository.Update(article)

	return &ArticleDto{
			Id:          resultArticle.Id(),
			Description: resultArticle.Description(),
			UserId:      resultArticle.UserId(),
		},
		nil
}
