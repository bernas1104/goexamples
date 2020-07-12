package routes

import (
	"github.com/bernas1104/goexamples/api-gin-example/src/controllers"
	"github.com/bernas1104/goexamples/api-gin-example/src/middlewares"
	"github.com/bernas1104/goexamples/api-gin-example/src/services"
	"github.com/gin-gonic/gin"

	"net/http"
)

var (
	listAllUsers services.ListAllUsersService = services.NewListAllUsersService()
	createUser   services.CreateUserService   = services.NewCreateUserService()
	showUser     services.ShowUserService     = services.NewShowUserService()

	usersController controllers.UsersController = controllers.NewUsersController(
		listAllUsers,
		createUser,
		showUser,
	)
)

// setupUsersRoutes Defines routes for the User entity
func setupUsersRoutes(app *gin.Engine) {
	users := app.Group("v1/users", middlewares.AuthorizeJWT())
	{
		users.GET("/", func(ctx *gin.Context) {
			ctx.JSON(http.StatusOK, usersController.Index(ctx))
		})

		users.GET("/:id", func(ctx *gin.Context) {
			user, err := usersController.Show(ctx)

			if err != nil {
				ctx.JSON(http.StatusBadRequest, gin.H{
					"error": err.Error(),
				})
			} else {
				ctx.JSON(http.StatusOK, user)
			}
		})

		users.POST("/", func(ctx *gin.Context) {
			user, err := usersController.Create(ctx)

			if err != nil {
				ctx.JSON(http.StatusBadRequest, gin.H{
					"error": err.Error(),
				})
			} else {
				ctx.JSON(http.StatusOK, user)
			}
		})
	}
}
