package auth

import (
	"study-pal-backend/app/app_types"
	"study-pal-backend/app/usecases/shared/usecase_error"
	"study-pal-backend/app/utils/study_pal_jwts"
)

type RefreshTokenCommand struct {
	RefreshToken string
}

type RefreshTokenDto struct {
	AccessToken string
}

type RefreshTokenAction struct {
	AppData app_types.AppData
}

func (l *RefreshTokenAction) Execute(command *RefreshTokenCommand) (*RefreshTokenDto, usecase_error.UsecaseErrorGroup) {
	userId, err := study_pal_jwts.VerifyToken(l.AppData.JwtSecretKey(), command.RefreshToken)
	if err != nil {
		return nil, usecase_error.NewUsecaseErrorGroupWithMessage(usecase_error.NewUsecaseError(usecase_error.UnPermittedOperation, err))
	}

	accessToken := study_pal_jwts.CreateAccessToken(l.AppData.JwtSecretKey(), userId)

	return &RefreshTokenDto{
			AccessToken: accessToken,
		},
		nil
}
