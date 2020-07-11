package controllers

import (
	"github.com/bernas1104/goexamples/api-gin-example/src/dtos"
	"github.com/bernas1104/goexamples/api-gin-example/src/services"
	"github.com/gin-gonic/gin"
)

type SessionsController interface {
	Create(ctx *gin.Context) string
}

type sessionsController struct {
	userAuthentication services.UserAuthenticationService
	jwt                services.JWTService
}

func NewSessionsController(
	userAuthentication services.UserAuthenticationService,
	jwt services.JWTService,
) SessionsController {
	return &sessionsController{
		userAuthentication: userAuthentication,
		jwt:                jwt,
	}
}

func (controller *sessionsController) Create(ctx *gin.Context) string {
	var userAuth dtos.UserAuthenticationDTO

	err := ctx.ShouldBindJSON(&userAuth)

	if err != nil {
		return ""
	}

	authenticated := controller.userAuthentication.AuthenticateUser(
		userAuth.Email,
		userAuth.Password,
	)

	if !authenticated {
		return ""
	}

	token := controller.jwt.GenerateToken(userAuth.Email, true)

	return token
}
