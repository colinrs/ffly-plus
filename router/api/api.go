package api

import (
	apiV1 "ffly-plus/controller/api/v1"

	"github.com/gin-gonic/gin"
)

// RegisterAPIV1 ...
func RegisterAPIV1(apiGroup *gin.RouterGroup) {

	apiUserGroup := apiGroup.Group("/user")
	//apiGroup.Use(jwt.JWT())
	registerUserAPIV1(apiUserGroup)
}

func registerUserAPIV1(apiUserGroup *gin.RouterGroup) {

	//获取标签列表
	apiUserGroup.GET("/", apiV1.GetUserV1)
	//新建标签
	apiUserGroup.POST("/", apiV1.AddUserV1)
	//更新指定标签
	apiUserGroup.PUT("/", apiV1.EditUserV1)
	//删除指定标签
	apiUserGroup.DELETE("/", apiV1.DeleteUserV1)
	//导出标签
}
