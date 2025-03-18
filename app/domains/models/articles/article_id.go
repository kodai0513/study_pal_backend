package articles

import (
	"study-pal-backend/app/domains/models/shared"
)

type ArticleId struct {
	value int
}

func NewArticleId(value int) (ArticleId, error) {
	id, err := shared.NewId(value)
	if err != nil {
		return ArticleId{value: id.Value()}, err
	}
	return ArticleId{value: id.Value()}, nil
}

func (a *ArticleId) Value() int {
	if a == nil {
		return 0
	}
	return a.value
}
