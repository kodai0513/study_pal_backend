package article

import "github.com/google/uuid"

type ArticleDto struct {
	Id          uuid.UUID
	Description string
	UserId      uuid.UUID
}
