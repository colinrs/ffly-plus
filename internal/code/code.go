package code

const (
	// SUCCESS ...
	SUCCESS = 200
	// Error ...
	Error = 500
	// InvalidParams ...
	InvalidParams = 400

	// ErrorExistUser ...
	ErrorExistUser = 10001

	// ErrorAuthCheckTokenFail ...
	ErrorAuthCheckTokenFail = 20001
	// ErrorAuthCheckTokenTimeout ...
	ErrorAuthCheckTokenTimeout = 20002
	// ErrorAuthToken ...
	ErrorAuthToken = 20003
	// ErrorAuth ...
	ErrorAuth = 20004
)

// MsgFlags ...
var MsgFlags = map[int]string{
	SUCCESS:                    "ok",
	Error:                      "fail",
	InvalidParams:              "请求参数错误",
	ErrorExistUser:             "已存在该用户名称",
	ErrorAuthCheckTokenFail:    "Token鉴权失败",
	ErrorAuthCheckTokenTimeout: "Token已超时",
	ErrorAuthToken:             "Token生成失败",
	ErrorAuth:                  "Token错误",
}

// GetMsg get Error information based on Code
func GetMsg(code int) string {
	msg, ok := MsgFlags[code]
	if ok {
		return msg
	}

	return MsgFlags[Error]
}
