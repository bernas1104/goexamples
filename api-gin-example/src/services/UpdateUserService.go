package services

import (
	"fmt"
	"time"

	"github.com/bernas1104/goexamples/api-gin-example/src/dtos"
	"github.com/bernas1104/goexamples/api-gin-example/src/entities"
	"github.com/bernas1104/goexamples/api-gin-example/src/providers"
	"github.com/bernas1104/goexamples/api-gin-example/src/repositories"
)

type UpdateUserService interface {
	UpdateUser(email string, dto dtos.UpdateUserDTO) (entities.User, error)
}

type updateUserService struct {
	usersRepository repositories.UsersRepository
	hashProvider    providers.HashProvider
}

func NewUpdateUserService() UpdateUserService {
	return &updateUserService{
		usersRepository: repositories.NewUsersRepository(),
		hashProvider:    providers.NewBCryptHashProvider(),
	}
}

func (service *updateUserService) UpdateUser(
	email string,
	dto dtos.UpdateUserDTO,
) (entities.User, error) {
	user := service.usersRepository.FindByEmail(email)

	if user.ID == 0 {
		return user, fmt.Errorf("User account was deleted")
	}

	uniqueEmail := service.usersRepository.FindByEmail(dto.Email)

	if uniqueEmail.ID != 0 {
		return uniqueEmail, fmt.Errorf("E-mail already being used")
	}

	user = entities.User{
		ID:                   user.ID,
		Name:                 dto.Name,
		Age:                  dto.Age,
		Email:                dto.Email,
		Password:             service.hashProvider.GenerateHash(dto.NewPassword),
		PasswordConfirmation: "",
		CreatedAt:            user.CreatedAt,
		UpdatedAt:            time.Now(),
	}

	user = service.usersRepository.Save(user)

	return user, nil
}
