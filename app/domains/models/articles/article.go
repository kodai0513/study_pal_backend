package articles

import "study-pal-backend/app/domains/models/shared"

type Article struct {
	id          *shared.Id
	description Description
	postId      PostId
}

func NewArticle(id *shared.Id, description Description, postId PostId) *Article {
	return &Article{
		id:          id,
		description: description,
		postId:      postId,
	}
}

func (a *Article) Id() int {
	if a.id != nil {
		return 0
	}
	return a.id.Value()
}

func (a *Article) Description() string {
	return a.description.Value()
}

func (a *Article) PostId() int {
	return a.postId.Value()
}
