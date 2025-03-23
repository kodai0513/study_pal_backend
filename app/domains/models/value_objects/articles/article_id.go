package articles

import (
	"study-pal-backend/app/domains/models/value_objects/shared/ids"

	"github.com/google/uuid"
)

type ArticleId struct {
	value uuid.UUID
}

func CreateArticleId() ArticleId {
	id := ids.CreateId()
	return ArticleId{value: id.Value()}
}

func NewArticleId(value uuid.UUID) ArticleId {
	return ArticleId{value: value}
}

func (a *ArticleId) Value() uuid.UUID {
	return a.value
}
