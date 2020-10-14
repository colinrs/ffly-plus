package api

import (
	"time"

	apiV1 "ffly-plus/controller/api/v1"
	"ffly-plus/router/middleware"

	"github.com/gin-contrib/cache"
	"github.com/gin-contrib/cache/persistence"
	"github.com/gin-gonic/gin"
)

// RegisterAPIV1 ...
func RegisterAPIV1(apiGroup *gin.RouterGroup) {

	apiUserGroup := apiGroup.Group("/user")
	// 用户登入
	apiGroup.POST("/login", apiV1.UserLogin)
	// 用户注册
	apiGroup.POST("/register", apiV1.UserRegister)
	apiGroup.GET("/logout", apiV1.UserRegister)

	registerUserAPIV1(apiUserGroup)
}

func registerUserAPIV1(apiUserGroup *gin.RouterGroup) {
	store := persistence.NewInMemoryStore(time.Minute)
	apiUserGroup.Use(middleware.AuthMiddleware())
	//获取自己信息
	apiUserGroup.GET("/", cache.CachePage(store, time.Minute, apiV1.GetUser))
	//apiUserGroup.GET("/", apiV1.GetUser)
	//更新用户信息
	apiUserGroup.PUT("/", apiV1.EditUser)
	//注销用户
	apiUserGroup.DELETE("/", apiV1.DeleteUser)
}
