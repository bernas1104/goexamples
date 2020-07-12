package services

import (
	"fmt"

	"github.com/bernas1104/goexamples/api-gin-example/src/entities"
	"github.com/bernas1104/goexamples/api-gin-example/src/repositories"
)

type ShowUserService interface {
	ShowService(id uint64) (entities.User, error)
}

type showUserService struct {
	usersRepository repositories.UsersRepository
}

func NewShowUserService() ShowUserService {
	return &showUserService{
		usersRepository: repositories.NewUsersRepository(),
	}
}

func (service *showUserService) ShowService(id uint64) (entities.User, error) {
	user := service.usersRepository.FindByID(id)

	if user.ID == 0 {
		return user, fmt.Errorf("User does not exist")
	}

	return user, nil
}
