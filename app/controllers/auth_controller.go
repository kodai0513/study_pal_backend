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
	AppData *app_types.AppData
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
	var request RefreshTokenRequest
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
	action := auth.RefreshTokenAction{
		AppData: *a.AppData,
	}
	refreshTokenDto, usecaseErrGroup := action.Execute(
		&auth.RefreshTokenCommand{
			RefreshToken: request.RefreshToken,
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
		&RefreshTokenResponse{
			AccessToken: refreshTokenDto.AccessToken,
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
	var request LoginRequest
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
	repository := repositories.NewUserRepositoryImpl(a.AppData.Client(), c)
	action := auth.LoginAction{
		AppData:        *a.AppData,
		UserRepository: repository,
	}
	loginDto, usecaseErrGroup := action.Execute(
		&auth.LoginCommand{
			Name:     request.Name,
			Password: request.Password,
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
		&LoginResponse{
			AccessToken:  loginDto.AccessToken,
			RefreshToken: loginDto.RefreshToken,
		},
	)
}
