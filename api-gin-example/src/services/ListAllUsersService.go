package services

import (
	"github.com/bernas1104/goexamples/api-gin-example/src/entities"
	"github.com/bernas1104/goexamples/api-gin-example/src/repositories"
)

type ListAllUsersService interface {
	ListAllUsers() []entities.User
}

type listAllUsersService struct {
	usersRepository repositories.UsersRepository
}

func NewListAllUsersService() ListAllUsersService {
	return &listAllUsersService{
		usersRepository: repositories.NewUsersRepository(),
	}
}

func (service *listAllUsersService) ListAllUsers() []entities.User {
	users := service.usersRepository.FindAll()
	return users
}
