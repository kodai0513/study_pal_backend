package repositories

import (
	"context"
	"study-pal-backend/app/domains/models/entities"
	"study-pal-backend/app/domains/repositories"
	"study-pal-backend/ent"
	"study-pal-backend/ent/articlelike"

	"github.com/google/uuid"
)

type ArticleLikeRepositoryImpl struct {
	tx  *ent.Tx
	ctx context.Context
}

func NewArticleLikeRepositoryImpl(tx *ent.Tx, ctx context.Context) repositories.ArticleLikeRepository {
	return &ArticleLikeRepositoryImpl{
		tx:  tx,
		ctx: ctx,
	}
}

func (a *ArticleLikeRepositoryImpl) Create(articleLike *entities.ArticleLike) {
	a.tx.ArticleLike.Create().
		SetID(articleLike.Id()).
		SetArticleID(articleLike.ArticleId()).
		SetUserID(articleLike.UserId()).
		SaveX(a.ctx)
}

func (a *ArticleLikeRepositoryImpl) Delete(articleLikeId uuid.UUID) {
	a.tx.ArticleLike.DeleteOneID(articleLikeId).ExecX(a.ctx)
}

func (a *ArticleLikeRepositoryImpl) ExistById(articleLikeId uuid.UUID) bool {
	return a.tx.ArticleLike.Query().Where(articlelike.IDEQ(articleLikeId)).ExistX(a.ctx)
}

func (a *ArticleLikeRepositoryImpl) ExistByArticleIdAndUserId(articleId uuid.UUID, userId uuid.UUID) bool {
	return a.tx.ArticleLike.Query().Where(
		articlelike.ArticleIDEQ(articleId),
		articlelike.UserIDEQ(userId),
	).ExistX(a.ctx)
}
