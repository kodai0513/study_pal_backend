package controllers

import (
	"net/http"
	"study-pal-backend/app/app_types"
	"study-pal-backend/app/controllers/shared/mappers"
	"study-pal-backend/app/infrastructures/repositories"
	"study-pal-backend/app/usecases/article_likes"
	"study-pal-backend/app/usecases/shared/trancaction"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type ArticleLikeController struct {
	AppData *app_types.AppData
}

// article-like godoc
//
//	@Summary	API
//	@Description
//	@Tags		article-like
//	@Accept		json
//	@Produce	json
//	@Param		article_id	path		string	true	"Article ID"
//	@Success	201			{object}	nil
//	@Failure	400			{object}	app_types.ErrorResponse
//	@Failure	401			{object}	app_types.ErrorResponse
//	@Failure	404			{object}	app_types.ErrorResponse
//	@Failure	500			{object}	app_types.ErrorResponse
//	@Router		/articles/{article_id}/likes [post]
func (a *ArticleLikeController) Create(c *gin.Context) {
	articleId, err := uuid.Parse(c.Param("article_id"))
	if err != nil {
		c.SecureJSON(
			http.StatusBadRequest,
			&app_types.ErrorResponse{
				Errors: []string{err.Error()},
			},
		)
		c.Abort()
		return
	}
	userId, _ := c.Get("user_id")

	tx, err := a.AppData.Client().Tx(c)
	if err != nil {
		panic(err)
	}
	action := &article_likes.CreateAction{
		ArticleRepository:     repositories.NewArticleRepositoryImpl(tx, c),
		ArticleLikeRepository: repositories.NewArticleLikeRepositoryImpl(tx, c),
		Tx:                    trancaction.NewTx(tx),
	}
	usecaseErrGroup := action.Execute(
		&article_likes.CreateActionCommand{
			ArticleId: articleId,
			UserId:    userId.(uuid.UUID),
		},
	)

	if usecaseErrGroup != nil && usecaseErrGroup.IsError() {
		c.SecureJSON(
			mappers.UsecaseErrorToHttpStatus(usecaseErrGroup),
			&app_types.ErrorResponse{
				Errors: usecaseErrGroup.Errors(),
			},
		)
		c.Abort()
		return
	}

	c.SecureJSON(
		http.StatusCreated,
		gin.H{},
	)
}

// article-like godoc
//
//	@Summary	API
//	@Description
//	@Tags		article-like
//	@Accept		json
//	@Produce	json
//	@Param		article_id		path		string	true	"Article ID"
//	@Param		article_like_id	path		string	true	"ArticleLike ID"
//	@Success	204				{object}	nil
//	@Failure	400				{object}	app_types.ErrorResponse
//	@Failure	401				{object}	app_types.ErrorResponse
//	@Failure	404				{object}	app_types.ErrorResponse
//	@Failure	500				{object}	app_types.ErrorResponse
//	@Router		/articles/{article_id}/likes/{article_like_id} [delete]
func (a *ArticleLikeController) Delete(c *gin.Context) {
	articleId, err := uuid.Parse(c.Param("article_id"))
	if err != nil {
		c.SecureJSON(
			http.StatusBadRequest,
			&app_types.ErrorResponse{
				Errors: []string{err.Error()},
			},
		)
		c.Abort()
		return
	}
	articleLikeId, err := uuid.Parse(c.Param("article_like_id"))
	if err != nil {
		c.SecureJSON(
			http.StatusBadRequest,
			&app_types.ErrorResponse{
				Errors: []string{err.Error()},
			},
		)
		c.Abort()
		return
	}
	userId, _ := c.Get("user_id")

	tx, err := a.AppData.Client().Tx(c)
	if err != nil {
		panic(err)
	}
	action := &article_likes.DeleteAction{
		ArticleRepository:     repositories.NewArticleRepositoryImpl(tx, c),
		ArticleLikeRepository: repositories.NewArticleLikeRepositoryImpl(tx, c),
		Tx:                    trancaction.NewTx(tx),
	}
	usecaseErrGroup := action.Execute(
		&article_likes.DeleteActionCommand{
			ArticleId:     articleId,
			ArticleLikeId: articleLikeId,
			UserId:        userId.(uuid.UUID),
		},
	)

	if usecaseErrGroup != nil && usecaseErrGroup.IsError() {
		c.SecureJSON(
			mappers.UsecaseErrorToHttpStatus(usecaseErrGroup),
			&app_types.ErrorResponse{
				Errors: usecaseErrGroup.Errors(),
			},
		)
		c.Abort()
		return
	}

	c.SecureJSON(
		http.StatusNoContent,
		gin.H{},
	)
}
