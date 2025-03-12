package users

import (
	"study-pal-backend/app/domains/models/shared"
)

type User struct {
	id       *shared.Id
	email    Email
	name     Name
	nickName NickName
	password Password
}

func NewUser(id *shared.Id, email Email, name Name, nickName NickName, password Password) *User {
	return &User{
		id:       id,
		email:    email,
		name:     name,
		nickName: nickName,
		password: password,
	}
}

func (u *User) Id() int {
	if u.id != nil {
		return 0
	}
	return u.id.Value()
}

func (u *User) Email() string {
	return u.email.Value()
}

func (u *User) Name() string {
	return u.email.Value()
}

func (u *User) NickName() string {
	return u.nickName.Value()
}

func (u *User) Password() string {
	return u.password.Value()
}
