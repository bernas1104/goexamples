package routes

import (
	"github.com/bernas1104/goexamples/api-gin-example/src/controllers"
	"github.com/bernas1104/goexamples/api-gin-example/src/services"
	"github.com/gin-gonic/gin"

	"net/http"
)

var (
	userAuthentication services.UserAuthenticationService = services.NewUserAuthenticationService()

	jwt services.JWTService = services.NewJWTService()

	sessionsController controllers.SessionsController = controllers.NewSessionsController(userAuthentication, jwt)
)

// setupSessionsRoutes Defines routes for creating a Session
func setupSessionsRoutes(app *gin.Engine) {
	sessions := app.Group("v1/sessions")
	{
		sessions.POST("/", func(ctx *gin.Context) {
			token := sessionsController.Create(ctx)

			if token != "" {
				ctx.JSON(http.StatusOK, gin.H{
					"token": token,
				})
			} else {
				ctx.JSON(http.StatusUnauthorized, gin.H{
					"error": "Invalid E-mail/Password combination",
				})
			}
		})
	}
}
