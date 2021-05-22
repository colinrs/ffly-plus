package router

import (
	"github.com/colinrs/ffly-plus/controller"
	apiV1 "github.com/colinrs/ffly-plus/controller/api/v1"
	"github.com/colinrs/ffly-plus/router/api"
	"github.com/colinrs/ffly-plus/router/middleware"

	//nolint: golint
	_ "github.com/colinrs/ffly-plus/docs"

	"github.com/colinrs/pkgx/server/gin"
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
	server.GinEngine.Engine.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	server.GinEngine.Use(middleware.AcclogSetUp())
	server.GinEngine.Use(middleware.SentinelMiddleware())

	registerBaseAPI(server)
	apiGroupV1 := server.GinEngine.Group("/api/v1")
	api.RegisterAPIV1(apiGroupV1)
	var models []gin.Model
	models = append(models, new(apiV1.UserController))
	server.GinEngine.RegisterModel(models)
	return server
}

// registerBaseAPI ...
func registerBaseAPI(server *Server) {
	server.GinEngine.Engine.GET("/", controller.Health)
	server.GinEngine.Engine.GET("/version", controller.Version)
}
