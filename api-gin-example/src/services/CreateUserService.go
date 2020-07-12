package services

import (
	"time"

	"github.com/bernas1104/goexamples/api-gin-example/src/entities"
	"github.com/bernas1104/goexamples/api-gin-example/src/providers"
	"github.com/bernas1104/goexamples/api-gin-example/src/repositories"
)

type CreateUserService interface {
	CreateUser(user entities.User) (entities.User, error)
}

type createUserService struct {
	usersRepository repositories.UsersRepository
	hashProvider    providers.HashProvider
}

func NewCreateUserService() CreateUserService {
	return &createUserService{
		usersRepository: repositories.NewUsersRepository(),
		hashProvider:    providers.NewBCryptHashProvider(),
	}
}

func (service *createUserService) CreateUser(user entities.User) (entities.User, error) {
	user.Password = service.hashProvider.GenerateHash(user.Password)

	user.CreatedAt = time.Now()
	user.UpdatedAt = time.Now()

	user, err := service.usersRepository.Create(user)

	if err != nil {
		return user, err
	}

	return user, nil
}
