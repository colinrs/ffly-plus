package router

import (
	"ffly-plus/controller"

	"github.com/gin-gonic/gin"
)

// InitRouter ...
func InitRouter() *controller.Server {
	server := new(controller.Server)
	gin.SetMode(gin.DebugMode)
	server.GinEngine = gin.Default()
	// router
	server.GinEngine.GET("/", controller.Health)

	return server
}
