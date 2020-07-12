package routes

import (
	"github.com/bernas1104/goexamples/api-gin-example/src/controllers"
	"github.com/bernas1104/goexamples/api-gin-example/src/middlewares"
	"github.com/bernas1104/goexamples/api-gin-example/src/services"
	"github.com/gin-gonic/gin"
)

var (
	listAllUsers services.ListAllUsersService = services.NewListAllUsersService()
	createUser   services.CreateUserService   = services.NewCreateUserService()
	showUser     services.ShowUserService     = services.NewShowUserService()
	updateUser   services.UpdateUserService   = services.NewUpdateUserService()
	deleteUser   services.DeleteUserService   = services.NewDeleteUserService()

	usersController controllers.UsersController = controllers.NewUsersController(
		listAllUsers,
		createUser,
		showUser,
		updateUser,
		deleteUser,
	)
)

// setupUsersRoutes Defines routes for the User entity
func setupUsersRoutes(app *gin.Engine) {
	users := app.Group("v1/users")
	{
		users.GET("/", usersController.Index)

		users.GET("/:id", usersController.Show)

		users.POST("/", usersController.Create)

		users.PUT("/", middlewares.AuthorizeJWT(), usersController.Update)

		users.DELETE("/", middlewares.AuthorizeJWT(), usersController.Delete)
	}
}
