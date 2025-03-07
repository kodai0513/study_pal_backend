package auths

import (
	"errors"
	"study-pal-backend/app/app_types"
	"study-pal-backend/app/domains/repositories/users"
	"study-pal-backend/app/utils/application_errors"
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

func (l *LoginCommand) Name() string {
	return l.name
}

func (l *LoginCommand) Password() string {
	return l.password
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
	userRepository users.UserRepository
}

func NewLoginAction(appData app_types.AppData, userRepository users.UserRepository) *LoginAction {
	return &LoginAction{
		appData:        appData,
		userRepository: userRepository,
	}
}

func (l *LoginAction) Execute(command *LoginCommand) (*LoginDto, application_errors.ApplicationError) {
	user, err := l.userRepository.FindByName(command.Name())
	if err != nil {
		return nil, application_errors.NewDatabaseConnectionApplicationError(err)
	}

	if user == nil {
		return nil, application_errors.NewDataNotFoundApplicationError(errors.New("user not found"))
	}

	err = password_hashes.CheckPasswordHash(command.password, user.Password())
	if err != nil {
		return nil, application_errors.NewClientInputValidationApplicationError(err)
	}

	accessToken, err := study_pal_jwts.CreateAccessToken(l.appData.JwtSecretKey(), user.Id())
	if err != nil {
		return nil, application_errors.NewFatalApplicationError(err)
	}

	refreshToken, err := study_pal_jwts.CreateRefreshToken(l.appData.JwtSecretKey(), user.Id())
	if err != nil {
		return nil, application_errors.NewFatalApplicationError(err)
	}

	return NewLoginDto(accessToken, refreshToken), nil
}
