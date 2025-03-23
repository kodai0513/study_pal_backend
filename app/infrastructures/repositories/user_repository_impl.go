package repositories

import (
	"context"
	"study-pal-backend/app/domains/models/entities"
	"study-pal-backend/app/domains/models/value_objects/users"
	"study-pal-backend/app/domains/repositories"
	"study-pal-backend/ent"
	"study-pal-backend/ent/user"
)

type UserRepositoryImpl struct {
	client *ent.Client
	ctx    context.Context
}

func NewUserRepositoryImpl(client *ent.Client, ctx context.Context) repositories.UserRepository {
	return &UserRepositoryImpl{
		client: client,
		ctx:    ctx,
	}
}

func (u *UserRepositoryImpl) FindByName(name string) *entities.User {
	result := u.client.User.
		Query().
		Where(user.NameEQ(name)).
		FirstX(u.ctx)

	if result == nil {
		return nil
	}

	id := users.NewUserId(result.ID)
	email, _ := users.NewEmail(result.Email)
	resultName, _ := users.NewName(result.Name)
	nickName, _ := users.NewNickName(result.NickName)
	password := users.NewPassword(result.Password)

	return entities.NewUser(id, email, resultName, nickName, password)
}
