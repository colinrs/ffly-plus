package service

import (
	"context"
	"fmt"

	"ffly-plus/internal/code"
	"ffly-plus/internal/config"
	"ffly-plus/models"
	"ffly-plus/pkg/token"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
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

	query := map[string]interface{}{
		"user_name": service.UserName,
		"nickname":  service.Nickname,
	}
	_, err := models.SelectUser(query)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil
		}
		return err
	}
	return code.ErrUserExistBefor
}

// Register 用户注册
func (service *UserRegisterService) Register() error {
	user := models.User{
		Nickname: service.Nickname,
		UserName: service.UserName,
		Status:   fmt.Sprintf("%s--status", service.Nickname),
		Avatar:   models.Active,
	}

	// 表单验证
	if err := service.valid(); err != nil {
		return err
	}

	// 加密密码
	if err := user.SetPassword(service.Password); err != nil {
		return code.ErrEncrypt
	}

	// 创建用户
	if err := models.CreateUser(&user); err != nil {
		return code.ErrUserCreate
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

	query := map[string]interface{}{
		"user_name": service.UserName,
	}
	user, err := models.SelectUser(query)
	if err != nil {
		return "", err
	}

	if !user.CheckPassword(service.Password) {
		return "", code.ErrEmailOrPassword
	}
	tokenContext := token.Context{
		UserID:         uint64(user.ID),
		Username:       user.UserName,
		ExpirationTime: int64(config.Conf.App.JwtExpirationTime),
	}
	tokenString, err := token.Sign(context.Background(), tokenContext, config.Conf.App.JwtSecret)
	return tokenString, err
}

// SelectUser ....
func SelectUser(query map[string]interface{}) (*models.User, error) {
	return models.SelectUser(query)
}

// DeletetUser ....
func DeletetUser(query map[string]interface{}) (*models.User, error) {
	return models.DeleteUser(query)
}

// UserUpdateService 管理用户更新务
type UserUpdateService struct {
	Nickname string `form:"nickname" json:"nickname"`
	Password string `form:"password" json:"password"`
}

// UserUpdate ...
func (service *UserUpdateService) UserUpdate(query map[string]interface{}) (*models.User, error) {

	update := map[string]interface{}{}
	if service.Password != "" {
		update["password"] = service.Password
	}
	if service.Nickname != "" {
		update["nickname"] = service.Nickname
	}
	if len(update) == 0 {
		return nil, nil
	}
	return models.UpdataUser(query, update)
}
