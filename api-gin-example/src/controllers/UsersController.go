package controllers

import (
	"strconv"

	"github.com/bernas1104/goexamples/api-gin-example/src/entities"
	"github.com/bernas1104/goexamples/api-gin-example/src/services"
	"github.com/gin-gonic/gin"
)

func hideUserPassword(user *entities.User) {
	user.Password = ""
	user.PasswordConfirmation = ""
}

type UsersController interface {
	Index(ctx *gin.Context) []entities.User
	Create(ctx *gin.Context) (entities.User, error)
	Show(ctx *gin.Context) (entities.User, error)
}

type usersController struct {
	listAllUsers services.ListAllUsersService
	createUser   services.CreateUserService
	showUser     services.ShowUserService
}

func NewUsersController(
	listAllUsersService services.ListAllUsersService,
	createUserService services.CreateUserService,
	showUserService services.ShowUserService,
) UsersController {
	return &usersController{
		listAllUsers: listAllUsersService,
		createUser:   createUserService,
		showUser:     showUserService,
	}
}

func (controller *usersController) Index(ctx *gin.Context) []entities.User {
	users := controller.listAllUsers.ListAllUsers()

	for i := range users {
		hideUserPassword(&users[i])
	}

	return users
}

func (controller *usersController) Create(ctx *gin.Context) (entities.User, error) {
	var user entities.User

	err := ctx.ShouldBindJSON(&user)

	if err != nil {
		return entities.User{}, err
	}

	user, err = controller.createUser.CreateUser(user)

	hideUserPassword(&user)

	return user, err
}

func (controller *usersController) Show(ctx *gin.Context) (entities.User, error) {
	var user entities.User

	param, err := strconv.Atoi(ctx.Param("id"))

	if err != nil {
		return user, err
	}

	id := uint64(param)
	user, err = controller.showUser.ShowService(id)

	if err != nil {
		return user, err
	}

	hideUserPassword(&user)

	return user, nil
}
