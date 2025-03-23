package entities

import (
	"study-pal-backend/app/domains/models/value_objects/users"

	"github.com/google/uuid"
)

type User struct {
	id       users.UserId
	email    users.Email
	name     users.Name
	nickName users.NickName
	password users.Password
}

func NewUser(id users.UserId, email users.Email, name users.Name, nickName users.NickName, password users.Password) *User {
	return &User{
		id:       id,
		email:    email,
		name:     name,
		nickName: nickName,
		password: password,
	}
}

func (u *User) Id() uuid.UUID {
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
