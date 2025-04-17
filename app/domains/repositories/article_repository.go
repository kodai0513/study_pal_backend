package repositories

import (
	"study-pal-backend/app/domains/models/entities"

	"github.com/google/uuid"
)

type ArticleRepository interface {
	Create(*entities.Article) *entities.Article
	Delete(uuid.UUID)
	ExistById(uuid.UUID) bool
	FindById(uuid.UUID) *entities.Article
	Update(*entities.Article) *entities.Article
}
