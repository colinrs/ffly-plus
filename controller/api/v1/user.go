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
	if uid, ok := c.Keys["uid"]; ok {
		query := map[string]interface{}{
			"id": uid,
		}
		user, err := service.SelectUser(query)
		internal.APIResponse(c, err, user)
		return
	}
	internal.APIResponse(c, code.ErrToken, nil)

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

	if uid, ok := c.Keys["uid"]; ok {
		query := map[string]interface{}{
			"id": uid,
		}
		var updateService service.UserUpdateService
		err := c.ShouldBind(&updateService)
		if err != nil {
			internal.APIResponse(c, err, nil)
			return
		}

		user, err := updateService.UserUpdate(query)
		internal.APIResponse(c, err, user)
		return
	}
	internal.APIResponse(c, code.ErrToken, nil)
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

	if uid, ok := c.Keys["uid"]; ok {
		query := map[string]interface{}{
			"id": uid,
		}
		user, err := service.DeletetUser(query)
		internal.APIResponse(c, err, user)
		return
	}
	internal.APIResponse(c, code.ErrToken, nil)
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
