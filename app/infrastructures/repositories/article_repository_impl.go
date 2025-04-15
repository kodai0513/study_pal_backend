package repositories

import (
	"context"
	"study-pal-backend/app/domains/models/entities"
	"study-pal-backend/app/domains/models/value_objects/articles"
	"study-pal-backend/app/domains/repositories"
	"study-pal-backend/ent"
	"study-pal-backend/ent/article"

	"github.com/google/uuid"
)

type ArticleRepositoryImpl struct {
	tx  *ent.Tx
	ctx context.Context
}

func NewArticleRepositoryImpl(tx *ent.Tx, ctx context.Context) repositories.ArticleRepository {
	return &ArticleRepositoryImpl{
		tx:  tx,
		ctx: ctx,
	}
}

func (a *ArticleRepositoryImpl) Create(article *entities.Article) *entities.Article {
	result := a.tx.Article.
		Create().
		SetID(article.Id()).
		SetDescription(article.Description()).
		SetPostID(article.UserId()).
		SaveX(a.ctx)

	description, _ := articles.NewDescription(result.Description)
	return entities.NewArticle(result.ID, description, result.UserID)
}

func (a *ArticleRepositoryImpl) Update(article *entities.Article) *entities.Article {
	result := a.tx.Article.
		UpdateOneID(article.Id()).
		SetDescription(article.Description()).
		SaveX(a.ctx)

	description, _ := articles.NewDescription(result.Description)
	return entities.NewArticle(result.ID, description, result.UserID)
}

func (a *ArticleRepositoryImpl) Delete(id uuid.UUID) {
	a.tx.Article.
		DeleteOneID(id).
		ExecX(a.ctx)
}

func (a *ArticleRepositoryImpl) FindById(id uuid.UUID) *entities.Article {
	result := a.tx.Article.
		Query().
		Where(article.IDEQ(id)).
		FirstX(a.ctx)

	if result == nil {
		return nil
	}

	resultDescription, _ := articles.NewDescription(result.Description)

	return entities.NewArticle(result.ID, resultDescription, result.UserID)
}
