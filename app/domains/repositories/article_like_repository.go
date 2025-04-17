package repositories

import (
	"study-pal-backend/app/domains/models/entities"

	"github.com/google/uuid"
)

type ArticleLikeRepository interface {
	Create(*entities.ArticleLike)
	Delete(uuid.UUID)
	ExistById(uuid.UUID) bool
	ExistByArticleIdAndUserId(uuid.UUID, uuid.UUID) bool
}
