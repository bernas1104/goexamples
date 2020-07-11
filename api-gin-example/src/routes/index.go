package routes

import "github.com/gin-gonic/gin"

// SetupRoutes Defines routes to the application
func SetupRoutes(app *gin.Engine) {
	// setupUsersRoutes(app)
	setupSessionsRoutes(app)
}
