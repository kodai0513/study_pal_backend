package auths

import (
	"errors"
	"study-pal-backend/app/app_types"
	"study-pal-backend/app/domains/repositories"
	"study-pal-backend/app/usecases/shared/usecase_errors"
	"study-pal-backend/app/utils/password_hashes"
	"study-pal-backend/app/utils/study_pal_jwts"
)

type LoginCommand struct {
	name     string
	password string
}

func NewLoginCommand(name string, password string) *LoginCommand {
	return &LoginCommand{
		name:     name,
		password: password,
	}
}

type LoginDto struct {
	accessToken  string
	refreshToken string
}

func NewLoginDto(accessToken string, refreshToken string) *LoginDto {
	return &LoginDto{
		accessToken:  accessToken,
		refreshToken: refreshToken,
	}
}

func (l *LoginDto) AccessToken() string {
	return l.accessToken
}

func (l *LoginDto) RefreshToken() string {
	return l.refreshToken
}

type LoginAction struct {
	appData        app_types.AppData
	userRepository repositories.UserRepository
}

func NewLoginAction(appData app_types.AppData, userRepository repositories.UserRepository) *LoginAction {
	return &LoginAction{
		appData:        appData,
		userRepository: userRepository,
	}
}

func (l *LoginAction) Execute(command *LoginCommand) (*LoginDto, usecase_errors.UsecaseErrorGroup) {
	user := l.userRepository.FindByName(command.name)

	if user == nil {
		return nil, usecase_errors.NewUsecaseErrorGroupWithMessage(usecase_errors.NewUsecaseError(usecase_errors.QueryDataNotFoundError, errors.New("user not found")))
	}

	err := password_hashes.CheckPasswordHash(command.password, user.Password())
	if err != nil {
		return nil, usecase_errors.NewUsecaseErrorGroupWithMessage(usecase_errors.NewUsecaseError(usecase_errors.InvalidParameter, err))
	}

	accessToken := study_pal_jwts.CreateAccessToken(l.appData.JwtSecretKey(), user.Id())
	refreshToken := study_pal_jwts.CreateRefreshToken(l.appData.JwtSecretKey(), user.Id())

	return NewLoginDto(accessToken, refreshToken), nil
}
