package repositories

import (
	"context"
	"study-pal-backend/app/domains/models/articles"
	"study-pal-backend/app/domains/models/shared"
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
		SetPostID(article.PostId()).
		SaveX(a.ctx)
}

func (a *ArticleRepositoryImpl) Update(article *articles.Article) {
	a.client.Article.
		UpdateOneID(article.Id()).
		SetDescription(article.Description()).
		SaveX(a.ctx)
}

func (a *ArticleRepositoryImpl) Delete(id shared.Id) {
	a.client.Article.
		DeleteOneID(id.Value()).
		ExecX(a.ctx)
}

func (a *ArticleRepositoryImpl) FindById(id shared.Id) *articles.Article {
	result := a.client.Article.
		Query().
		Where(article.IDEQ(id.Value())).
		FirstX(a.ctx)

	if result == nil {
		return nil
	}

	resultId, _ := shared.NewId(result.ID)
	resultDescription, _ := articles.NewDescription(result.Description)
	resultPostId, _ := articles.NewPostId(result.PostID)

	return articles.NewArticle(resultId, resultDescription, resultPostId)
}
