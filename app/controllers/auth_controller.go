package controllers

import (
	"net/http"
	"study-pal-backend/app/app_types"
	"study-pal-backend/app/infrastructures/repositories"
	"study-pal-backend/app/usecases/auths"
	"study-pal-backend/app/utils/application_errors"

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
	refreshTokenDto, err := auths.NewRefreshTokenAction(*a.appData).Execute(auths.NewRefreshTokenCommand(refreshTokenRequest.RefreshToken))
	if err != nil && err.Kind() == application_errors.ClientInputValidation {
		c.SecureJSON(
			http.StatusBadRequest,
			app_types.NewErrorResponse([]string{err.Error()}),
		)
		return
	}

	if err != nil && err.Kind() == application_errors.FatalError {
		c.SecureJSON(
			http.StatusInternalServerError,
			app_types.NewErrorResponse([]string{err.Error()}),
		)
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
	loginDto, err := auths.NewLoginAction(*a.appData, repositories.NewUserRepositoryImpl(c, a.appData.Client())).Execute(
		auths.NewLoginCommand(loginRequest.Name, loginRequest.Password),
	)

	if err != nil && (err.Kind() == application_errors.DatabaseConnection || err.Kind() == application_errors.FatalError) {
		c.SecureJSON(
			http.StatusInternalServerError,
			app_types.NewErrorResponse([]string{err.Error()}),
		)
		return
	}

	if err != nil && (err.Kind() == application_errors.ClientInputValidation) {
		c.SecureJSON(
			http.StatusBadRequest,
			app_types.NewErrorResponse([]string{err.Error()}),
		)
		return
	}

	if err != nil && err.Kind() == application_errors.DataNotFound {
		c.SecureJSON(
			http.StatusNotFound,
			app_types.NewErrorResponse([]string{err.Error()}),
		)
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
