package auth

import (
	"errors"
	"study-pal-backend/app/app_types"
	"study-pal-backend/app/domains/repositories"
	"study-pal-backend/app/usecases/shared/usecase_error"
	"study-pal-backend/app/utils/password_hashes"
	"study-pal-backend/app/utils/study_pal_jwts"
)

type LoginCommand struct {
	Name     string
	Password string
}

type LoginDto struct {
	AccessToken  string
	RefreshToken string
}

type LoginAction struct {
	AppData        app_types.AppData
	UserRepository repositories.UserRepository
}

func (l *LoginAction) Execute(command *LoginCommand) (*LoginDto, usecase_error.UsecaseErrorGroup) {
	user := l.UserRepository.FindByName(command.Name)

	if user == nil {
		return nil, usecase_error.NewUsecaseErrorGroupWithMessage(usecase_error.NewUsecaseError(usecase_error.QueryDataNotFoundError, errors.New("user not found")))
	}

	err := password_hashes.CheckPasswordHash(command.Password, user.Password())
	if err != nil {
		return nil, usecase_error.NewUsecaseErrorGroupWithMessage(usecase_error.NewUsecaseError(usecase_error.InvalidParameter, err))
	}

	accessToken := study_pal_jwts.CreateAccessToken(l.AppData.JwtSecretKey(), user.Id())
	refreshToken := study_pal_jwts.CreateRefreshToken(l.AppData.JwtSecretKey(), user.Id())

	return &LoginDto{
			AccessToken:  accessToken,
			RefreshToken: refreshToken,
		},
		nil
}
