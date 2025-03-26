package entities

import (
	"study-pal-backend/app/domains/models/value_objects/users"

	"github.com/google/uuid"
)

type User struct {
	id       uuid.UUID
	email    users.Email
	name     users.Name
	nickName users.NickName
	password users.Password
}

func NewUser(id uuid.UUID, email users.Email, name users.Name, nickName users.NickName, password users.Password) *User {
	return &User{
		id:       id,
		email:    email,
		name:     name,
		nickName: nickName,
		password: password,
	}
}

func (u *User) Id() uuid.UUID {
	return u.id
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

func (u *User) SetEmail(email users.Email) {
	u.email = email
}

func (u *User) SetName(name users.Name) {
	u.name = name
}

func (u *User) SetNickName(nickName users.NickName) {
	u.nickName = nickName
}

func (u *User) SetPassword(password users.Password) {
	u.password = password
}
