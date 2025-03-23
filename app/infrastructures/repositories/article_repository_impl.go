package repositories

import (
	"context"
	"study-pal-backend/app/domains/models/entities"
	"study-pal-backend/app/domains/models/value_objects/articles"
	"study-pal-backend/app/domains/models/value_objects/users"
	"study-pal-backend/app/domains/repositories"
	"study-pal-backend/ent"
	"study-pal-backend/ent/article"

	"github.com/google/uuid"
)

type ArticleRepositoryImpl struct {
	ctx    context.Context
	client *ent.Client
}

func NewArticleRepositoryImpl(ctx context.Context, client *ent.Client) repositories.ArticleRepository {
	return &ArticleRepositoryImpl{
		ctx:    ctx,
		client: client,
	}
}

func (a *ArticleRepositoryImpl) Create(article *entities.Article) {
	a.client.Article.
		Create().
		SetID(article.Id()).
		SetDescription(article.Description()).
		SetPostID(article.UserId()).
		SaveX(a.ctx)
}

func (a *ArticleRepositoryImpl) Update(article *entities.Article) {
	a.client.Article.
		UpdateOneID(article.Id()).
		SetDescription(article.Description()).
		SaveX(a.ctx)
}

func (a *ArticleRepositoryImpl) Delete(id uuid.UUID) {
	a.client.Article.
		DeleteOneID(id).
		ExecX(a.ctx)
}

func (a *ArticleRepositoryImpl) FindById(id uuid.UUID) *entities.Article {
	result := a.client.Article.
		Query().
		Where(article.IDEQ(id)).
		FirstX(a.ctx)

	if result == nil {
		return nil
	}

	resultId := articles.NewArticleId(result.ID)
	resultDescription, _ := articles.NewDescription(result.Description)
	userId := users.NewUserId(result.PostID)

	return entities.NewArticle(resultId, resultDescription, userId)
}
