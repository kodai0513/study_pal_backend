package repositories

import (
	"study-pal-backend/app/domains/models/articles"
	"study-pal-backend/app/domains/models/shared"
)

type ArticleRepository interface {
	Save(article *articles.Article)
	Update(article *articles.Article)
	Delete(id shared.Id)
	FindById(id shared.Id) *articles.Article
}
