package auths

import (
	"errors"
	"study-pal-backend/app/app_types"
	"study-pal-backend/app/utils/application_errors"
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

func (l *RefreshTokenAction) Execute(command *RefreshTokenCommand) (*RefreshTokenDto, application_errors.ApplicationError) {
	isValid, userId, err := study_pal_jwts.VerifyToken(l.appData.JwtSecretKey(), command.refreshToken)
	if err != nil {
		return nil, application_errors.NewFatalApplicationError(err)
	}

	if !isValid {
		return nil, application_errors.NewClientInputValidationApplicationError(errors.New("invalid token"))
	}

	accessToken, err := study_pal_jwts.CreateAccessToken(l.appData.JwtSecretKey(), userId)
	if err != nil {
		return nil, application_errors.NewFatalApplicationError(err)
	}

	return &RefreshTokenDto{
		accessToken: accessToken,
	}, nil
}
