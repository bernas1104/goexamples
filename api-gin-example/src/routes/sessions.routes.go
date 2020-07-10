package routes

import (
	"github.com/gin-gonic/gin"

	"net/http"
)

// setupSessionsRoutes Defines routes for creating a Session
func setupSessionsRoutes(app *gin.Engine) {
	sessions := app.Group("v1/sessions")
	{
		sessions.POST("/", func(ctx *gin.Context) {
			ctx.JSON(http.StatusOK, gin.H{
				"message": "Login to API",
			})
		})
	}
}
