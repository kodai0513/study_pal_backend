package articles

import (
	"study-pal-backend/app/domains/models/articles"
)

type ArticleRepository interface {
	Save(article *articles.Article) (*articles.Article, error)
}
