package routes

import (
	"github.com/bernas1104/goexamples/api-gin-example/src/controllers"
	"github.com/bernas1104/goexamples/api-gin-example/src/middlewares"
	"github.com/bernas1104/goexamples/api-gin-example/src/services"
	"github.com/gin-gonic/gin"

	"net/http"
)

var (
	listAllUsers    services.ListAllUsersService = services.NewListAllUsersService()
	createUser      services.CreateUserService   = services.NewCreateUserService()
	usersController controllers.UsersController  = controllers.NewUsersController(
		listAllUsers,
		createUser,
	)
)

// setupUsersRoutes Defines routes for the User entity
func setupUsersRoutes(app *gin.Engine) {
	users := app.Group("v1/users", middlewares.AuthorizeJWT())
	{
		users.GET("/", func(ctx *gin.Context) {
			ctx.JSON(http.StatusOK, usersController.Index(ctx))
		})

		users.POST("/", func(ctx *gin.Context) {
			user := usersController.Create(ctx)

			if user.ID == 0 {
				ctx.JSON(http.StatusBadRequest, gin.H{
					"error": "Error while creating new User",
				})
			}

			ctx.JSON(http.StatusOK, user)
		})
	}
}
