package articles

import "study-pal-backend/app/domains/models/users"

type Article struct {
	id          ArticleId
	description Description
	userId      users.UserId
}

func NewArticle(id ArticleId, description Description, userId users.UserId) *Article {
	return &Article{
		id:          id,
		description: description,
		userId:      userId,
	}
}

func (a *Article) Id() int {
	return a.id.Value()
}

func (a *Article) Description() string {
	return a.description.Value()
}

func (a *Article) UserId() int {
	return a.userId.Value()
}
