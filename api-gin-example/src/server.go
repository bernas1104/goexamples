package main

import (
	"github.com/bernas1104/goexamples/api-gin-example/src/routes"
	"github.com/gin-gonic/gin"
)

// Server defines the API main file
var Server *gin.Engine

func main() {
	Server := gin.Default()

	routes.SetupRoutes(Server)

	Server.Run(":3333")
}
