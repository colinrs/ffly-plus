package router

import (
	"ffly-plus/controller"
	"ffly-plus/router/api"

	"github.com/gin-gonic/gin"
)

// Server ...
type Server struct {
	GinEngine *gin.Engine
}

// InitRouter ...
func InitRouter() *Server {
	server := new(Server)
	gin.SetMode(gin.DebugMode)
	server.GinEngine = gin.Default()
	server.GinEngine.Use(gin.Recovery())
	server.GinEngine.Use(gin.Logger())

	registerBaseAPI(server)
	apiGroupV1 := server.GinEngine.Group("/api/v1")
	api.RegisterAPIV1(apiGroupV1)

	return server
}

// registerBaseAPI ...
func registerBaseAPI(server *Server) {
	server.GinEngine.GET("/", controller.Health)
}
