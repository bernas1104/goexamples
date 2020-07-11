package routes

import (
	"github.com/bernas1104/goexamples/api-gin-example/src/middlewares"
	"github.com/gin-gonic/gin"

	"net/http"
)

// setupUsersRoutes Defines routes for the User entity
func setupUsersRoutes(app *gin.Engine) {
	users := app.Group("v1/users", middlewares.AuthorizeJWT())
	{
		users.GET("/", func(ctx *gin.Context) {
			ctx.JSON(http.StatusOK, gin.H{
				"message": "Users index route",
			})
		})
	}
}
