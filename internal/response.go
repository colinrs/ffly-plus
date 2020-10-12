package internal

import (
	"net/http"

	"ffly-plus/internal/code"

	"github.com/gin-gonic/gin"
)

// Response ...
type Response struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

// APIResponse ....
func APIResponse(C *gin.Context, err error, data interface{}) {
	code, message := code.DecodeErr(err)
	C.JSON(http.StatusOK, Response{
		Code: code,
		Msg:  message,
		Data: data,
	})
	return
}
