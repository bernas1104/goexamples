package services

import (
	"time"

	"github.com/bernas1104/goexamples/api-gin-example/src/models"
)

// UsersService define the methods available to work with an User
type UsersService interface {
	Save(models.User) models.User
	FindAll() []models.User
}

type usersService struct {
	users []models.User
}

// New defines a initializer for a usersService struct
func New() UsersService {
	return &usersService{}
}

func (service *usersService) Save(user models.User) models.User {
	user.CreatedAt = time.Now()
	user.UpdatedAt = time.Now()
	service.users = append(service.users, user)

	return user
}

func (service *usersService) FindAll() []models.User {
	return service.users
}
