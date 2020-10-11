package service

import (
	"ffly-plus/internal/code"
	"ffly-plus/models"

	"github.com/gin-gonic/gin"
)

// UserRegisterService 管理用户注册服务
type UserRegisterService struct {
	Nickname        string `form:"nickname" json:"nickname" binding:"required,min=2,max=30"`
	UserName        string `form:"user_name" json:"user_name" binding:"required,min=5,max=30"`
	Password        string `form:"password" json:"password" binding:"required,min=8,max=40"`
	PasswordConfirm string `form:"password_confirm" json:"password_confirm" binding:"required,min=8,max=40"`
}

// valid 验证表单
func (service *UserRegisterService) valid() error {
	if service.PasswordConfirm != service.Password {
		return code.ErrPasswordIncorrect
	}

	query := models.User{
		UserName: service.UserName,
		Nickname: service.Nickname,
	}
	user, err := models.SelectUser(&query)
	if err != nil {
		return err
	}
	if user.UserName != "" {
		return code.ErrUserNotFound
	}
	return nil
}

// Register 用户注册
func (service *UserRegisterService) Register() error {
	user := models.User{
		Nickname: service.Nickname,
		UserName: service.UserName,
		Status:   models.Active,
	}

	// 表单验证
	if err := service.valid(); err != nil {
		return code.ErrParam
	}

	// 加密密码
	if err := user.SetPassword(service.Password); err != nil {
		return code.ErrEncrypt
	}

	// 创建用户
	if err := models.CreateUser(&user); err != nil {
		return code.ErrUserNotFound
	}

	return nil
}

// UserLoginService 管理用户登录的服务
type UserLoginService struct {
	UserName string `form:"user_name" json:"user_name" binding:"required,min=5,max=30"`
	Password string `form:"password" json:"password" binding:"required,min=8,max=40"`
}

// Login 用户登录函数
func (service *UserLoginService) Login(c *gin.Context) (string, error) {
	var user *models.User

	query := models.User{
		UserName: service.UserName,
	}
	user, err := models.SelectUser(&query)
	if err != nil {
		return "", err
	}

	if user.CheckPassword(service.Password) == false {
		return "", code.ErrEmailOrPassword
	}

	return "token", nil
}
