package api

import (
	apiV1 "ffly-plus/controller/api/v1"

	"github.com/gin-gonic/gin"
)

// RegisterAPIV1 ...
func RegisterAPIV1(apiGroup *gin.RouterGroup) {

	apiUserGroup := apiGroup.Group("/user")
	// 用户登入
	apiUserGroup.POST("/login", apiV1.UserLogin)
	// 用户注册
	apiUserGroup.POST("/register", apiV1.UserRegister)
	apiUserGroup.GET("/logout", apiV1.UserRegister)

	//apiGroup.Use(jwt.JWT())
	registerUserAPIV1(apiUserGroup)
}

func registerUserAPIV1(apiUserGroup *gin.RouterGroup) {

	//获取自己信息
	apiUserGroup.GET("/", apiV1.GetUser)
	//更新用户信息
	apiUserGroup.PUT("/", apiV1.EditUser)
	//注销用户
	apiUserGroup.DELETE("/", apiV1.DeleteUser)
}
