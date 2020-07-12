package services

import (
	"github.com/bernas1104/goexamples/api-gin-example/src/repositories"
)

type DeleteUserService interface {
	DeleteUser(email string)
}

type deleteUserService struct {
	usersRepository repositories.UsersRepository
}

func NewDeleteUserService() DeleteUserService {
	return &deleteUserService{
		usersRepository: repositories.NewUsersRepository(),
	}
}

func (service *deleteUserService) DeleteUser(email string) {
	service.usersRepository.Delete(email)
}
