package users

type User struct {
	id       UserId
	email    Email
	name     Name
	nickName NickName
	password Password
}

func NewUser(id UserId, email Email, name Name, nickName NickName, password Password) *User {
	return &User{
		id:       id,
		email:    email,
		name:     name,
		nickName: nickName,
		password: password,
	}
}

func (u *User) Id() int {
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
