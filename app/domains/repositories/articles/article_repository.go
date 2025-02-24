package articles

import (
	"study-pal-backend/app/domains/models/articles"
	"study-pal-backend/app/utils/application_errors"
)

type ArticleRepository interface {
	Save(article articles.Article) (articles.Article, application_errors.ApplicationError)
}
