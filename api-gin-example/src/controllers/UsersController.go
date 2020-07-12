package controllers

import (
	"net/http"
	"strconv"

	"github.com/bernas1104/goexamples/api-gin-example/src/dtos"

	"github.com/bernas1104/goexamples/api-gin-example/src/entities"
	"github.com/bernas1104/goexamples/api-gin-example/src/services"
	"github.com/gin-gonic/gin"
)

func hideUserPassword(user *entities.User) {
	user.Password = ""
	user.PasswordConfirmation = ""
}

type UsersController interface {
	Index(ctx *gin.Context)
	Create(ctx *gin.Context)
	Show(ctx *gin.Context)
	Update(ctx *gin.Context)
	Delete(ctx *gin.Context)
}

type usersController struct {
	listAllUsers services.ListAllUsersService
	createUser   services.CreateUserService
	showUser     services.ShowUserService
	updateUser   services.UpdateUserService
	deleteUSer   services.DeleteUserService
}

func NewUsersController(
	listAllUsersService services.ListAllUsersService,
	createUserService services.CreateUserService,
	showUserService services.ShowUserService,
	updateUserService services.UpdateUserService,
	deleteUserService services.DeleteUserService,
) UsersController {
	return &usersController{
		listAllUsers: listAllUsersService,
		createUser:   createUserService,
		showUser:     showUserService,
		updateUser:   updateUserService,
		deleteUSer:   deleteUserService,
	}
}

func (controller *usersController) Index(ctx *gin.Context) {
	users := controller.listAllUsers.ListAllUsers()

	for i := range users {
		hideUserPassword(&users[i])
	}

	ctx.JSON(http.StatusOK, users)
}

func (controller *usersController) Create(ctx *gin.Context) {
	var user entities.User

	err := ctx.ShouldBindJSON(&user)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	user, err = controller.createUser.CreateUser(user)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	hideUserPassword(&user)

	ctx.JSON(http.StatusOK, user)
}

func (controller *usersController) Show(ctx *gin.Context) {
	var user entities.User

	param, err := strconv.Atoi(ctx.Param("id"))

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	id := uint64(param)
	user, err = controller.showUser.ShowService(id)

	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"error": err.Error(),
		})
		return
	}

	hideUserPassword(&user)

	ctx.JSON(http.StatusOK, user)
}

func (controller *usersController) Update(ctx *gin.Context) {
	var userUpdate dtos.UpdateUserDTO

	err := ctx.ShouldBindJSON(&userUpdate)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})

		return
	}

	auth, _ := ctx.Get("Email")
	email := auth.(string)

	user, err := controller.updateUser.UpdateUser(email, userUpdate)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})

		return
	}

	hideUserPassword(&user)

	ctx.JSON(http.StatusOK, user)
}

func (controller *usersController) Delete(ctx *gin.Context) {
	auth, _ := ctx.Get("Email")

	email := auth.(string)

	controller.deleteUSer.DeleteUser(email)

	ctx.JSON(http.StatusNoContent, struct{}{})
}
