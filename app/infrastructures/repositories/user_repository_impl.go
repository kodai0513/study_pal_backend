package repositories

import (
	"context"
	"study-pal-backend/app/domains/models/users"
	"study-pal-backend/app/domains/models/value_object_commons"
	user_repository "study-pal-backend/app/domains/repositories/users"
	"study-pal-backend/ent"
	"study-pal-backend/ent/user"
)

type UserRepositoryImpl struct {
	ctx    context.Context
	client *ent.Client
}

func NewUserRepositoryImpl(ctx context.Context, client *ent.Client) user_repository.UserRepository {
	return &UserRepositoryImpl{
		ctx:    ctx,
		client: client,
	}
}

func (u *UserRepositoryImpl) FindByName(name string) (*users.User, error) {
	result, err := u.client.User.
		Query().
		Where(user.NameEQ(name)).
		First(u.ctx)

	if err != nil {
		return nil, err
	}

	id, _ := value_object_commons.NewId(result.ID)
	email, _ := users.NewEmail(result.Email)
	resultName, _ := users.NewName(result.Name)
	nickName, _ := users.NewNickName(result.NickName)
	password := users.NewPassword(result.Password)

	return users.NewUser(id, email, resultName, nickName, password), nil
}
