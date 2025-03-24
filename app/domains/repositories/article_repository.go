package repositories

import (
	"study-pal-backend/app/domains/models/entities"

	"github.com/google/uuid"
)

type ArticleRepository interface {
	Create(article *entities.Article) *entities.Article
	Delete(id uuid.UUID)
	FindById(id uuid.UUID) *entities.Article
	Update(article *entities.Article) *entities.Article
}
