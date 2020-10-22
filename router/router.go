package router

import (
	"ffly-plus/controller"
	"ffly-plus/router/api"

	_ "ffly-plus/docs"

	"github.com/gin-contrib/pprof"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
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
	pprof.Register(server.GinEngine)
	server.GinEngine.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
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
	server.GinEngine.GET("/version", controller.Version)
}
