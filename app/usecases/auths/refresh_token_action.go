package auths

import (
	"study-pal-backend/app/app_types"
	"study-pal-backend/app/usecases/shared/usecase_errors"
	"study-pal-backend/app/utils/study_pal_jwts"
)

type RefreshTokenCommand struct {
	refreshToken string
}

func NewRefreshTokenCommand(refreshToken string) *RefreshTokenCommand {
	return &RefreshTokenCommand{
		refreshToken: refreshToken,
	}
}

type RefreshTokenDto struct {
	accessToken string
}

func (l *RefreshTokenDto) AccessToken() string {
	return l.accessToken
}

type RefreshTokenAction struct {
	appData app_types.AppData
}

func NewRefreshTokenAction(appData app_types.AppData) *RefreshTokenAction {
	return &RefreshTokenAction{
		appData: appData,
	}
}

func (l *RefreshTokenAction) Execute(command *RefreshTokenCommand) (*RefreshTokenDto, usecase_errors.UsecaseErrorGroup) {
	userId, err := study_pal_jwts.VerifyToken(l.appData.JwtSecretKey(), command.refreshToken)
	if err != nil {
		return nil, usecase_errors.NewUsecaseErrorGroupWithMessage(usecase_errors.NewUsecaseError(usecase_errors.UnPermittedOperation, err))
	}

	accessToken := study_pal_jwts.CreateAccessToken(l.appData.JwtSecretKey(), userId)

	return &RefreshTokenDto{
		accessToken: accessToken,
	}, nil
}
