package internal

import (
	"ffly-plus/internal/code"

	"github.com/gin-gonic/gin"
)

// Gin ...
type Gin struct {
	C *gin.Context
}

// Response ...
type Response struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

// NewGin ...
func NewGin(C *gin.Context) *Gin {
	return &Gin{
		C: C,
	}
}

// Response setting gin.JSON
func (g *Gin) Response(httpCode int, err error, data interface{}) {
	code, message := code.DecodeErr(err)
	g.C.JSON(httpCode, Response{
		Code: code,
		Msg:  message,
		Data: data,
	})
	return
}
