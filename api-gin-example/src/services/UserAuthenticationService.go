package services

import (
	"github.com/bernas1104/goexamples/api-gin-example/src/providers"
	"github.com/bernas1104/goexamples/api-gin-example/src/repositories"
)

type UserAuthenticationService interface {
	AuthenticateUser(email string, password string) bool
}

type userAuthenticationService struct {
	usersRepository repositories.UsersRepository
	hashProvider    providers.HashProvider
}

func NewUserAuthenticationService() UserAuthenticationService {
	return &userAuthenticationService{
		usersRepository: repositories.NewUsersRepository(),
		hashProvider:    providers.NewBCryptHashProvider(),
	}
}

func (service *userAuthenticationService) AuthenticateUser(
	email string,
	password string,
) bool {
	user := service.usersRepository.FindByEmail(email)

	if user.ID == 0 {
		return false
	}

	auth := service.hashProvider.Verify(user.Password, password)

	return auth
}
