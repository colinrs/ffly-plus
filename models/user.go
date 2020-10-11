package models

import (
	"github.com/colinrs/pkgx/logger"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

// User 用户模型
type User struct {
	CustomeModel
	UserName       string `gorm:"type:varchar(30);not null;unique;index:user_name_idx"`
	PasswordDigest string `json:"password_digest"`
	Nickname       string `gorm:"type:varchar(30);not null;unique;index:nickname_idx"`
	Status         string `json:"status"`
	Avatar         string `gorm:"size:1000"`
}

const (
	// PassWordCost 密码加密难度
	PassWordCost = 12
	// Active 激活用户
	Active string = "active"
	// Inactive 未激活用户
	Inactive string = "inactive"
	// Suspend 被封禁用户
	Suspend string = "suspend"
	// UserTable ...
	UserTable string = "users"
)

// GetUser 用ID获取用户
func GetUser(ID interface{}) (User, error) {
	var user User
	result := DB.First(&user, ID)
	return user, result.Error
}

// CreateUser ...
func CreateUser(user *User) error {
	result := DB.Table(UserTable).Create(user)
	if result.Error != nil {
		logger.Error("created user:{%#v} err:%s", user, result.Error.Error())
		return result.Error
	}
	logger.Info("insert user:{%#v}", user)
	return nil
}

// SelectUser ...
func SelectUser(query *User) (*User, error) {
	var user *User
	result := DB.Table(UserTable).Where(query).First(user)
	if result.Error != nil {
		logger.Error("query user:{%#v} err:%s", query, result.Error.Error())
		return nil, result.Error
	}
	return user, nil
}

// UpdataUser ...
func UpdataUser(query *User) (*User, error) {
	var user *User
	result := DB.Table(UserTable).Model(query).Updates(
		User{Status: Suspend})
	if result.Error != nil {
		logger.Error("query user:{%#v} err:%s", query, result.Error.Error())
		return nil, result.Error
	}
	result.Take(user)
	return user, nil
}

// DeleteUser ...
func DeleteUser(query *User) (*User, error) {
	var user *User
	result := DB.Table(UserTable).Where(query).Delete(user)
	if result.Error != nil {
		logger.Error("query user:{%#v} err:%s", query, result.Error.Error())
		return nil, result.Error
	}
	return user, nil
}

// SetPassword 设置密码
func (user *User) SetPassword(password string) error {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), PassWordCost)
	if err != nil {
		return err
	}
	user.PasswordDigest = string(bytes)
	return nil
}

// CheckPassword 校验密码
func (user *User) CheckPassword(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(user.PasswordDigest), []byte(password))
	return err == nil
}

// BeforeCreate ...
func (user *User) BeforeCreate(tx *gorm.DB) error {
	logger.Info("Create user BeforeCreate")
	return nil
}

// BeforeUpdate ...
func (user *User) BeforeUpdate(tx *gorm.DB) error {
	logger.Info("Update user BeforeUpdate")
	return nil
}

// BeforeDelete ...
func (user *User) BeforeDelete(tx *gorm.DB) error {
	logger.Info("Delete user BeforeDelete")
	return nil
}
