package api

import (
	"net/http"

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
	gin := internal.NewGin(c)
	gin.Response(http.StatusOK, code.OK, "GetUserV1")
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
	gin := internal.NewGin(c)
	gin.Response(http.StatusOK, code.OK, "EditUserV1")
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
	gin := internal.NewGin(c)
	gin.Response(http.StatusOK, code.OK, "DeleteUserV1")
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
	gin := internal.NewGin(c)

	err := c.ShouldBind(&service)
	if err != nil {
		gin.Response(http.StatusBadRequest, err, nil)
		return
	}

	err = service.Register()
	if err != nil {
		gin.Response(http.StatusInternalServerError, err, nil)
		return
	}

	gin.Response(http.StatusOK, code.OK, "EditUserV1")
	return

}

// UserLogin 用户登录接口
func UserLogin(c *gin.Context) {
	gin := internal.NewGin(c)
	var service service.UserLoginService
	err := c.ShouldBind(&service)
	if err != nil {
		gin.Response(http.StatusOK, code.ErrBind, nil)
		return
	}
	token, err := service.Login(c)

	gin.Response(http.StatusOK, err, token)

}
