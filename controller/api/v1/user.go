package api

import (
	"ffly-plus/internal"
	"ffly-plus/internal/code"
	"ffly-plus/service"

	"github.com/gin-gonic/gin"
)

// GetUser ...
// @Summary Get multiple article tags
// @Produce  json
// @Param name query string false "Name"
// @Param state query int false "State"
// @Success 200 {object} app.Response
// @Failure 500 {object} app.Response
// @Router /api/v1/user [get]
func GetUser(c *gin.Context) {
	internal.APIResponse(c, code.OK, "GetUserV1")
}

// EditUser ...
// @Summary Get multiple article tags
// @Produce  json
// @Param name query string false "Name"
// @Param state query int false "State"
// @Success 200 {object} app.Response
// @Failure 500 {object} app.Response
// @Router /api/v1/user [put]
func EditUser(c *gin.Context) {
	internal.APIResponse(c, code.OK, "EditUserV1")
}

// DeleteUser ...
// @Summary Get multiple article tags
// @Produce  json
// @Param name query string false "Name"
// @Param state query int false "State"
// @Success 200 {object} app.Response
// @Failure 500 {object} app.Response
// @Router /api/v1/user [delete]
func DeleteUser(c *gin.Context) {

	internal.APIResponse(c, code.OK, "DeleteUserV1")
}

// UserRegister 用户注册接口
// @Summary Get multiple article tags
// @Produce  json
// @Param name query string false "Name"
// @Param state query int false "State"
// @Success 200 {object} app.Response
// @Failure 500 {object} app.Response
// @Router /api/v1/user/register [post]
func UserRegister(c *gin.Context) {
	var service service.UserRegisterService

	err := c.ShouldBind(&service)
	if err != nil {
		internal.APIResponse(c, err, nil)
		return
	}

	err = service.Register()
	if err != nil {
		internal.APIResponse(c, err, nil)
		return
	}

	internal.APIResponse(c, code.OK, service)
	return

}

// UserLogin 用户登录接口
func UserLogin(c *gin.Context) {
	var service service.UserLoginService
	err := c.ShouldBind(&service)
	if err != nil {
		internal.APIResponse(c, code.ErrBind, err.Error())
		return
	}
	token, err := service.Login(c)

	internal.APIResponse(c, err, token)

}
