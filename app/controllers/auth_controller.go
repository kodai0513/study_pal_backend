package controllers

import (
	"net/http"
	"study-pal-backend/app/app_types"
	"study-pal-backend/app/controllers/shared/mappers"
	"study-pal-backend/app/infrastructures/repositories"
	"study-pal-backend/app/usecases/auth"

	"github.com/gin-gonic/gin"
)

type AuthController struct {
	appData *app_types.AppData
}

func NewAuthController(appData *app_types.AppData) *AuthController {
	return &AuthController{
		appData: appData,
	}
}

type RefreshTokenRequest struct {
	RefreshToken string `json:"refresh_token"`
}

type RefreshTokenResponse struct {
	AccessToken string `json:"access_token"`
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
func (a *AuthController) RefreshToken(c *gin.Context) {
	var refreshTokenRequest RefreshTokenRequest
	c.BindJSON(&refreshTokenRequest)
	refreshTokenDto, usecaseErrGroup := auth.NewRefreshTokenAction(*a.appData).Execute(auth.NewRefreshTokenCommand(refreshTokenRequest.RefreshToken))

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
		&RefreshTokenResponse{
			AccessToken: refreshTokenDto.AccessToken(),
		},
	)
}

type LoginRequest struct {
	Name     string `json:"name"`
	Password string `json:"password"`
}

type LoginResponse struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

// auth godoc
//
//	@Summary		ログイン情報取得API
//	@Description	ログイン情報を取得します
//	@Tags			auth
//	@Accept			json
//	@Produce		json
//	@Param			request	body		LoginRequest	true	"ログイン情報"
//	@Success		200		{object}	LoginResponse
//	@Failure		400		{object}	app_types.ErrorResponse
//	@Failure		500		{object}	app_types.ErrorResponse
//	@Router			/login [post]
func (a *AuthController) Login(c *gin.Context) {
	var loginRequest LoginRequest
	c.BindJSON(&loginRequest)
	loginDto, usecaseErrGroup := auth.NewLoginAction(*a.appData, repositories.NewUserRepositoryImpl(a.appData.Client(), c)).Execute(
		auth.NewLoginCommand(loginRequest.Name, loginRequest.Password),
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
		&LoginResponse{
			AccessToken:  loginDto.AccessToken(),
			RefreshToken: loginDto.RefreshToken(),
		},
	)
}
