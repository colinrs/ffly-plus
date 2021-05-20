package router

import (
	"github.com/colinrs/ffly-plus/controller"
	"github.com/colinrs/ffly-plus/router/api"
	"github.com/colinrs/ffly-plus/router/middleware"

	//nolint: golint
	_ "github.com/colinrs/ffly-plus/docs"

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
	//server.GinEngine.Use(gin.Logger())
	server.GinEngine.Use(middleware.AcclogSetUp())
	server.GinEngine.Use(middleware.SentinelMiddleware())

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
