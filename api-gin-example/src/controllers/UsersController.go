package controllers

import (
	"github.com/bernas1104/goexamples/api-gin-example/src/models"
	"github.com/bernas1104/goexamples/api-gin-example/src/services"
	"github.com/gin-gonic/gin"
)

// UsersController defines the methods of an UsersController
type UsersController interface {
	Index() []models.User
	Create(ctx *gin.Context) models.User
}

type controller struct {
	service services.UsersService
}

// New Creates an instance of controller and injects its dependencies
func New(service services.UsersService) UsersController {
	return &controller{
		service: service,
	}
}

func (c *controller) Index() []models.User {
	return c.service.FindAll()
}

func (c *controller) Create(ctx *gin.Context) models.User {
	var user models.User
	ctx.BindJSON(&user)
	user = c.service.Save(user)
	return user
}
