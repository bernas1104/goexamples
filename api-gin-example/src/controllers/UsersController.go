package controllers

import (
	"github.com/bernas1104/goexamples/api-gin-example/src/entities"
	"github.com/bernas1104/goexamples/api-gin-example/src/services"
	"github.com/gin-gonic/gin"
)

type UsersController interface {
	Index(ctx *gin.Context) []entities.User
	Create(ctx *gin.Context) entities.User
}

type usersController struct {
	listAllUsers services.ListAllUsersService
	createUser   services.CreateUserService
}

func NewUsersController(
	listAllUsersService services.ListAllUsersService,
	createUserService services.CreateUserService,
) UsersController {
	return &usersController{
		listAllUsers: listAllUsersService,
		createUser:   createUserService,
	}
}

func (controller *usersController) Index(ctx *gin.Context) []entities.User {
	users := controller.listAllUsers.ListAllUsers()
	return users
}

func (controller *usersController) Create(ctx *gin.Context) entities.User {
	var user entities.User

	err := ctx.ShouldBindJSON(&user)

	if err != nil {
		return entities.User{}
	}

	user = controller.createUser.CreateUser(user)

	return user
}
