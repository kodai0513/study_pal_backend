package controllers

import (
	"net/http"
	"study-pal-backend/app/app_types"
	"study-pal-backend/app/controllers/shared/mappers"
	"study-pal-backend/app/infrastructures/repositories"
	"study-pal-backend/app/usecases/article"
	"study-pal-backend/app/utils/type_converts"

	"github.com/gin-gonic/gin"
)

type ArticleController struct {
	appData *app_types.AppData
}

func NewArticleController(appData *app_types.AppData) *ArticleController {
	return &ArticleController{
		appData: appData,
	}
}

type CreateArticleRequest struct {
	Description string `json:"description"`
}

// article godoc
//
//	@Summary		投稿作成API
//	@Description	投稿を作成します
//	@Tags			article
//	@Accept			json
//	@Produce		json
//	@Param			request	body		CreateArticleRequest	true	"投稿作成リクエスト"
//	@Success		201		{object}	nil
//	@Failure		400		{object}	app_types.ErrorResponse
//	@Failure		401		{object}	app_types.ErrorResponse
//	@Failure		500		{object}	app_types.ErrorResponse
//	@Router			/articles [post]
func (a *ArticleController) Create(c *gin.Context) {
	var request CreateArticleRequest
	c.BindJSON(&request)
	userId, _ := c.Get("user_id")
	usecaseErrGroup := article.NewCreateAction(repositories.NewArticleRepositoryImpl(c, a.appData.Client())).Execute(
		article.NewCreateActionCommand(request.Description, userId.(int)),
	)

	if usecaseErrGroup != nil && usecaseErrGroup.IsError() {
		c.SecureJSON(
			mappers.UsecaseErrorToHttpStatus(usecaseErrGroup),
			app_types.NewErrorResponse(usecaseErrGroup.Errors()),
		)
		c.Abort()
		return
	}

	c.SecureJSON(
		http.StatusCreated,
		gin.H{},
	)
}

type UpdateArticleRequest struct {
	Description string `json:"description"`
}

// article godoc
//
//	@Summary		投稿更新API
//	@Description	投稿を更新します
//	@Tags			article
//	@Accept			json
//	@Produce		json
//	@Param			article_id	path		int						true	"投稿ID"
//	@Param			request		body		UpdateArticleRequest	true	"投稿作成リクエスト"
//	@Success		200			{object}	nil
//	@Failure		400			{object}	app_types.ErrorResponse
//	@Failure		401			{object}	app_types.ErrorResponse
//	@Failure		500			{object}	app_types.ErrorResponse
//	@Router			/articles/{article_id} [put]
func (a *ArticleController) Update(c *gin.Context) {
	var request UpdateArticleRequest
	c.BindJSON(&request)
	articleId := c.Param("article_id")
	userId, _ := c.Get("user_id")
	usecaseErrGroup := article.NewUpdateAction(repositories.NewArticleRepositoryImpl(c, a.appData.Client())).Execute(
		article.NewUpdateActionCommand(type_converts.StringToInt(articleId, 0), request.Description, userId.(int)),
	)

	if usecaseErrGroup != nil && usecaseErrGroup.IsError() {
		c.SecureJSON(
			mappers.UsecaseErrorToHttpStatus(usecaseErrGroup),
			app_types.NewErrorResponse(usecaseErrGroup.Errors()),
		)
		c.Abort()
		return
	}

	c.SecureJSON(
		http.StatusOK,
		gin.H{},
	)
}

// article godoc
//
//	@Summary		投稿削除API
//	@Description	投稿を削除します
//	@Tags			article
//	@Accept			json
//	@Produce		json
//	@Param			article_id	path		int	true	"投稿ID"
//	@Success		204			{object}	nil
//	@Failure		400			{object}	app_types.ErrorResponse
//	@Failure		401			{object}	app_types.ErrorResponse
//	@Failure		500			{object}	app_types.ErrorResponse
//	@Router			/articles/{article_id} [delete]
func (a *ArticleController) Delete(c *gin.Context) {
	articleId := c.Param("article_id")
	userId, _ := c.Get("user_id")
	usecaseErrGroup := article.NewDeleteAction(repositories.NewArticleRepositoryImpl(c, a.appData.Client())).Execute(
		article.NewDeleteActionCommand(type_converts.StringToInt(articleId, 0), userId.(int)),
	)

	if usecaseErrGroup != nil && usecaseErrGroup.IsError() {
		c.SecureJSON(
			mappers.UsecaseErrorToHttpStatus(usecaseErrGroup),
			app_types.NewErrorResponse(usecaseErrGroup.Errors()),
		)
		c.Abort()
		return
	}

	c.SecureJSON(
		http.StatusNoContent,
		gin.H{},
	)
}
