package users

import "study-pal-backend/app/domains/repositories"

type UsernameDuplicationCheckDomainService struct {
	userRepository repositories.UserRepository
}

func NewUsernameDuplicationCheckDomainService(userRepository repositories.UserRepository) *UsernameDuplicationCheckDomainService {
	return &UsernameDuplicationCheckDomainService{
		userRepository: userRepository,
	}
}

func (u *UsernameDuplicationCheckDomainService) Execute(name string) bool {
	user := u.userRepository.FindByName(name)
	return user != nil
}
