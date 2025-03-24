package entities

import (
	"study-pal-backend/app/domains/models/value_objects/articles"

	"github.com/google/uuid"
)

type Article struct {
	id          uuid.UUID
	description articles.Description
	userId      uuid.UUID
}

func NewArticle(id uuid.UUID, description articles.Description, userId uuid.UUID) *Article {
	return &Article{
		id:          id,
		description: description,
		userId:      userId,
	}
}

func (a *Article) Id() uuid.UUID {
	return a.id
}

func (a *Article) Description() string {
	return a.description.Value()
}

func (a *Article) UserId() uuid.UUID {
	return a.userId
}
