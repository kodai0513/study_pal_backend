package repositories

import (
	"study-pal-backend/app/domains/models/articles"
)

type ArticleRepository interface {
	Save(article *articles.Article)
	Update(article *articles.Article)
	Delete(id int)
	FindById(id int) *articles.Article
}
