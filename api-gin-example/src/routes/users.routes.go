package routes

import (
	"github.com/bernas1104/goexamples/api-gin-example/src/controllers"
	"github.com/bernas1104/goexamples/api-gin-example/src/services"
	"github.com/gin-gonic/gin"

	"net/http"
)

var (
	usersService    services.UsersService       = services.New()
	usersController controllers.UsersController = controllers.New(usersService)
)

// setupUsersRoutes Defines routes for the User entity
func setupUsersRoutes(app *gin.Engine) {
	users := app.Group("v1/users")
	{
		users.GET("/", func(ctx *gin.Context) {
			ctx.JSON(http.StatusOK, usersController.Index())
		})

		users.GET("/:id", func(ctx *gin.Context) {
			ctx.JSON(http.StatusOK, gin.H{
				"message": "User show route",
			})
		})

		users.POST("/", func(ctx *gin.Context) {
			ctx.JSON(http.StatusOK, usersController.Create(ctx))
		})

		users.PUT("/:id", func(ctx *gin.Context) {
			ctx.JSON(http.StatusOK, gin.H{
				"message": "User update route",
			})
		})

		users.DELETE("/:id", func(ctx *gin.Context) {
			ctx.JSON(http.StatusOK, gin.H{
				"message": "User delete route",
			})
		})
	}
}
