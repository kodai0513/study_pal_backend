package controllers

import "study-pal-backend/app/app_types"

type ArticleController struct {
	appData *app_types.AppData
}

func NewArticleController(appData *app_types.AppData) *ArticleController {
	return &ArticleController{
		appData: appData,
	}
}

// auth godoc
//
//	@Summary		トークン再発行取得API
//	@Description	アクセストークンを取得します
//	@Tags			auth
//	@Accept			json
//	@Produce		json
//	@Param			request	body		RefreshTokenRequest	true	"リフレッシュトークン"
//	@Success		200		{object}	RefreshTokenResponse
//	@Failure		400		{object}	app_types.ErrorResponse
//	@Failure		500		{object}	app_types.ErrorResponse
//	@Router			/refresh-token [post]
func (a *ArticleController) Create() {

}

// auth godoc
//
//	@Summary		トークン再発行取得API
//	@Description	アクセストークンを取得します
//	@Tags			auth
//	@Accept			json
//	@Produce		json
//	@Param			request	body		RefreshTokenRequest	true	"リフレッシュトークン"
//	@Success		200		{object}	RefreshTokenResponse
//	@Failure		400		{object}	app_types.ErrorResponse
//	@Failure		500		{object}	app_types.ErrorResponse
//	@Router			/refresh-token [post]
func (a *ArticleController) Update() {

}

// auth godoc
//
//	@Summary		トークン再発行取得API
//	@Description	アクセストークンを取得します
//	@Tags			auth
//	@Accept			json
//	@Produce		json
//	@Param			request	body		RefreshTokenRequest	true	"リフレッシュトークン"
//	@Success		200		{object}	RefreshTokenResponse
//	@Failure		400		{object}	app_types.ErrorResponse
//	@Failure		500		{object}	app_types.ErrorResponse
//	@Router			/refresh-token [post]
func (a *ArticleController) Delete() {

}
