package users

import "study-pal-backend/app/domains/repositories/users"

type UsernameDuplicationCheckDomainService struct {
	userRepository users.UserRepository
}

func NewUsernameDuplicationCheckDomainService(userRepository users.UserRepository) *UsernameDuplicationCheckDomainService {
	return &UsernameDuplicationCheckDomainService{
		userRepository: userRepository,
	}
}

func (u *UsernameDuplicationCheckDomainService) Execute(name string) bool {
	user, _ := u.userRepository.FindByName(name)
	return user != nil
}
