package repositories

import (
	"context"
	"study-pal-backend/app/domains/models/articles"
	"study-pal-backend/app/domains/models/users"
	"study-pal-backend/app/domains/repositories"
	"study-pal-backend/ent"
	"study-pal-backend/ent/article"
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

func (a *ArticleRepositoryImpl) Save(article *articles.Article) {
	a.client.Article.
		Create().
		SetDescription(article.Description()).
		SetPostID(article.UserId()).
		SaveX(a.ctx)
}

func (a *ArticleRepositoryImpl) Update(article *articles.Article) {
	a.client.Article.
		UpdateOneID(article.Id()).
		SetDescription(article.Description()).
		SaveX(a.ctx)
}

func (a *ArticleRepositoryImpl) Delete(id int) {
	a.client.Article.
		DeleteOneID(id).
		ExecX(a.ctx)
}

func (a *ArticleRepositoryImpl) FindById(id int) *articles.Article {
	result := a.client.Article.
		Query().
		Where(article.IDEQ(id)).
		FirstX(a.ctx)

	if result == nil {
		return nil
	}

	resultId, _ := articles.NewArticleId(result.ID)
	resultDescription, _ := articles.NewDescription(result.Description)
	userId, _ := users.NewUserId(result.PostID)

	return articles.NewArticle(resultId, resultDescription, userId)
}
