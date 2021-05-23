package api

import (
	"context"
	"fmt"

	"github.com/colinrs/ffly-plus/internal"
	"github.com/colinrs/ffly-plus/internal/code"
	"github.com/colinrs/ffly-plus/service"

	serverGin "github.com/colinrs/pkgx/server/gin"
	"github.com/gin-gonic/gin"
)

// UserController ...
type UserController struct {
}

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

// UserHelloResponse ...
type UserHelloResponse struct {
	Message string `json:"message"`
}

// Output ...
func (userHelloResponse UserHelloResponse) Output() {
	fmt.Print("run userHelloResponse Output\n")
}

// UserHelloRequest ...
type UserHelloRequest struct {
	UserMessage string `form:"user_message" json:"user_message"`
	Age         int    `form:"age" json:"age" binding:"ccless=10"`
}

// Validator ...
func (userHelloRequest UserHelloRequest) Validator(ctx context.Context) error {
	fmt.Print("run userHelloRequest Validator\n")
	return nil
}

// UserHello ...
func UserHello(ctx context.Context, request *UserHelloRequest) (UserHelloResponse, error) {
	res := UserHelloResponse{}
	res.Message = request.UserMessage
	return res, nil
}

// Init ...
func (uc *UserController) Init(engine *serverGin.Engine) error {
	fmt.Print("run UserController Init\n")
	engine.GET("/user_hello", UserHello)
	return nil
}
