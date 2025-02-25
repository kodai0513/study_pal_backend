package users

type User struct {
	email    Email
	name     Name
	nickName NickName
	password Password
}

func NewUser(email Email, name Name, nickName NickName, password Password) *User {
	return &User{
		email:    email,
		name:     name,
		nickName: nickName,
		password: password,
	}
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
