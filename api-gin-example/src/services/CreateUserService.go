package services

import (
	"github.com/bernas1104/goexamples/api-gin-example/src/entities"
	"github.com/bernas1104/goexamples/api-gin-example/src/repositories"
)

type CreateUserService interface {
	CreateUser(user entities.User) entities.User
}

type createUserService struct {
	usersRepository repositories.UsersRepository
}

func NewCreateUserService() CreateUserService {
	return &createUserService{
		usersRepository: repositories.NewUsersRepository(),
	}
}

func (service *createUserService) CreateUser(user entities.User) entities.User {
	user = service.usersRepository.Save(user)

	return user
}
