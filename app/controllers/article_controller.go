package controllers

import (
	"net/http"
	"study-pal-backend/app/app_types"
	"study-pal-backend/app/controllers/shared/mappers"
	"study-pal-backend/app/infrastructures/repositories"
	"study-pal-backend/app/usecases/article"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type ArticleController struct {
	AppData *app_types.AppData
}

type CreateArticle struct {
	Description string `json:"description"`
}

type CreateArticleRequest struct {
	CreateArticle
}

type CreateArticleResponse struct {
	CreateArticle
}

// article godoc
//
//	@Summary		投稿作成API
//	@Description	投稿を作成します
//	@Tags			article
//	@Accept			json
//	@Produce		json
//	@Param			request	body		CreateArticleRequest	true	"投稿作成リクエスト"
//	@Success		201		{object}	CreateArticleResponse
//	@Failure		400		{object}	app_types.ErrorResponse
//	@Failure		401		{object}	app_types.ErrorResponse
//	@Failure		500		{object}	app_types.ErrorResponse
//	@Router			/articles [post]
func (a *ArticleController) Create(c *gin.Context) {
	var request CreateArticleRequest
	err := c.BindJSON(&request)
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
	action := article.CreateAction{
		ArticleRepository: repositories.NewArticleRepositoryImpl(a.AppData.Client(), c),
	}
	articleDto, usecaseErrGroup := action.Execute(
		&article.CreateActionCommand{
			Description: request.Description,
			UserId:      userId.(uuid.UUID),
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
		&CreateArticleResponse{
			CreateArticle: CreateArticle{
				Description: articleDto.Description,
			},
		},
	)
}

type UpdateArticleRequest struct {
	Description string `json:"description"`
}

type UpdateArticleResponse struct {
	Description string    `json:"description"`
	UserId      uuid.UUID `json:"user_id"`
}

// article godoc
//
//	@Summary		投稿更新API
//	@Description	投稿を更新します
//	@Tags			article
//	@Accept			json
//	@Produce		json
//	@Param			article_id	path		string					true	"Article ID"
//	@Param			request		body		UpdateArticleRequest	true	"投稿更新リクエスト"
//	@Success		200			{object}	nil
//	@Failure		400			{object}	app_types.ErrorResponse
//	@Failure		401			{object}	app_types.ErrorResponse
//	@Failure		500			{object}	app_types.ErrorResponse
//	@Router			/articles/{article_id} [put]
func (a *ArticleController) Update(c *gin.Context) {
	var request UpdateArticleRequest
	err := c.BindJSON(&request)
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
	articleIdParam := c.Param("article_id")
	articleId, err := uuid.Parse(articleIdParam)
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
	action := &article.UpdateAction{
		ArticleRepository: repositories.NewArticleRepositoryImpl(a.AppData.Client(), c),
	}
	articleDto, usecaseErrGroup := action.Execute(
		&article.UpdateActionCommand{
			ArticleId:   articleId,
			Description: request.Description,
			UserId:      userId.(uuid.UUID),
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
		http.StatusOK,
		&UpdateArticleResponse{
			Description: articleDto.Description,
		},
	)
}

// article godoc
//
//	@Summary		投稿削除API
//	@Description	投稿を削除します
//	@Tags			article
//	@Accept			json
//	@Produce		json
//	@Param			article_id	path		string	true	"Article ID"
//	@Success		204			{object}	nil
//	@Failure		400			{object}	app_types.ErrorResponse
//	@Failure		401			{object}	app_types.ErrorResponse
//	@Failure		500			{object}	app_types.ErrorResponse
//	@Router			/articles/{article_id} [delete]
func (a *ArticleController) Delete(c *gin.Context) {
	articleIdParam := c.Param("article_id")
	articleId, err := uuid.Parse(articleIdParam)
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
	action := article.DeleteAction{
		ArticleRepository: repositories.NewArticleRepositoryImpl(a.AppData.Client(), c),
	}
	usecaseErrGroup := action.Execute(
		&article.DeleteActionCommand{
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
		http.StatusNoContent,
		gin.H{},
	)
}
