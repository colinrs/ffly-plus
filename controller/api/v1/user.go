package api

import (
	"github.com/colinrs/ffly-plus/internal"
	"github.com/colinrs/ffly-plus/internal/code"
	"github.com/colinrs/ffly-plus/service"

	"github.com/gin-gonic/gin"
)

// GetUser ...
// @Summary GetUser
// @Produce  json
// @Success 200 {object} internal.Response
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
// @Summary EditUser
// @Produce json
// @Param user_info body service.UserUpdateService true "user info"
// @Success 200 {object} internal.Response
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
// @Summary DeleteUser
// @Success 200 {object} internal.Response
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
// @Summary UserRegister
// @Produce  json
// @Param user body service.UserRegisterService true "user info"
// @Success 200 {object} internal.Response
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
