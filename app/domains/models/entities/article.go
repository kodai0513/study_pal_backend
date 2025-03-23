package entities

import (
	"study-pal-backend/app/domains/models/value_objects/articles"
	"study-pal-backend/app/domains/models/value_objects/users"

	"github.com/google/uuid"
)

type Article struct {
	id          articles.ArticleId
	description articles.Description
	userId      users.UserId
}

func NewArticle(id articles.ArticleId, description articles.Description, userId users.UserId) *Article {
	return &Article{
		id:          id,
		description: description,
		userId:      userId,
	}
}

func (a *Article) Id() uuid.UUID {
	return a.id.Value()
}

func (a *Article) Description() string {
	return a.description.Value()
}

func (a *Article) UserId() uuid.UUID {
	return a.userId.Value()
}
