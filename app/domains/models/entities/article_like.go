package entities

import (
	"github.com/google/uuid"
)

type ArticleLike struct {
	id        uuid.UUID
	articleId uuid.UUID
	userId    uuid.UUID
}

func NewArticleLike(id uuid.UUID, articleId uuid.UUID, userId uuid.UUID) *ArticleLike {
	return &ArticleLike{
		id:        id,
		articleId: articleId,
		userId:    userId,
	}
}

func (a *ArticleLike) Id() uuid.UUID {
	return a.id
}

func (a *ArticleLike) ArticleId() uuid.UUID {
	return a.articleId
}

func (a *ArticleLike) UserId() uuid.UUID {
	return a.userId
}
