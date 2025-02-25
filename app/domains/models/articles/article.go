package articles

type Article struct {
	description Description
	postId      PostId
}

func NewArticle(description Description, postId PostId) *Article {
	return &Article{
		description: description,
		postId:      postId,
	}
}

func (a *Article) Description() string {
	return a.description.Value()
}

func (a *Article) PostId() int {
	return a.postId.Value()
}
