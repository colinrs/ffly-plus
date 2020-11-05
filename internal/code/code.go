package code

import "fmt"

//nolint: golint
var (
	// Common errors
	OK                  = &Errno{Code: 0, Message: "OK"}
	InternalServerError = &Errno{Code: 10001, Message: "Internal server error"}
	ErrBind             = &Errno{Code: 10002, Message: "Error occurred while binding the request body to the struct."}
	ErrParam            = &Errno{Code: 10003, Message: "参数有误"}
	ErrSignParam        = &Errno{Code: 10004, Message: "签名参数有误"}

	ErrValidation         = &Errno{Code: 20001, Message: "Validation failed."}
	ErrDatabase           = &Errno{Code: 20002, Message: "Database error."}
	ErrToken              = &Errno{Code: 20003, Message: "Error occurred while signing the JSON web token."}
	ErrInvalidTransaction = &Errno{Code: 20004, Message: "invalid transaction."}

	// user errors
	ErrEncrypt               = &Errno{Code: 20101, Message: "密码加密错误"}
	ErrUserNotFound          = &Errno{Code: 20102, Message: "用户不存在"}
	ErrTokenInvalid          = &Errno{Code: 20103, Message: "Token错误"}
	ErrPasswordIncorrect     = &Errno{Code: 20104, Message: "密码错误"}
	ErrUserExistBefor        = &Errno{Code: 20105, Message: "用户已存在"}
	ErrUserCreate            = &Errno{Code: 20105, Message: "用户创建错误"}
	ErrSendSMSTooMany        = &Errno{Code: 20109, Message: "已超出当日限制，请明天再试"}
	ErrVerifyCode            = &Errno{Code: 20110, Message: "验证码错误"}
	ErrEmailOrPassword       = &Errno{Code: 20111, Message: "邮箱或密码错误"}
	ErrTwicePasswordNotMatch = &Errno{Code: 20112, Message: "两次密码输入不一致"}
	ErrRegisterFailed        = &Errno{Code: 20113, Message: "注册失败"}
	ErrCreatedUser           = &Errno{Code: 20114, Message: "用户创建失败"}
)

// Errno ...
type Errno struct {
	Code    int
	Message string
}

func (err Errno) Error() string {
	return err.Message
}

// Err represents an error
type Err struct {
	Code    int
	Message string
	Err     error
}

func (err *Err) Error() string {
	return fmt.Sprintf("Err - code: %d, message: %s, error: %s", err.Code, err.Message, err.Err)
}

// DecodeErr ...
func DecodeErr(err error) (int, string) {
	if err == nil {
		return OK.Code, OK.Message
	}

	switch typed := err.(type) {
	case *Err:
		return typed.Code, typed.Message
	case *Errno:
		return typed.Code, typed.Message
	default:
	}

	return InternalServerError.Code, err.Error()
}
