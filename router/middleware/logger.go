package middleware

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"time"

	"github.com/colinrs/ffly-plus/internal"

	"github.com/gin-gonic/gin"
)

var defaultWriter io.Writer = os.Stdout

type bodyLogWriter struct {
	gin.ResponseWriter
	body *bytes.Buffer
}

func (w bodyLogWriter) Write(b []byte) (int, error) {
	w.body.Write(b)
	return w.ResponseWriter.Write(b)

}

func (w bodyLogWriter) WriteString(s string) (int, error) {

	w.body.WriteString(s)
	return w.ResponseWriter.WriteString(s)

}

type line struct {
	requestStime    time.Time
	requestMethod   string
	requestURI      string
	requestProto    string
	requestUA       string
	requestReferer  string
	requestPOSTData string
	requestClientIP string
	responseCode    int
	responseMsg     string
	responseData    interface{}
	requestEtime    time.Time
	CostTime        string
}

func logFormatter(line *line) string {
	return fmt.Sprintf("[GIN] %s|%s|%s|%s|%s|%s|%s|%d|%s|%v|%s|%s\n",
		line.requestStime.Format("2006/01/02 - 15:04:05"),
		line.requestMethod,
		line.requestURI,
		line.requestUA,
		line.requestReferer,
		line.requestPOSTData,
		line.requestClientIP,
		line.responseCode,
		line.responseMsg,
		line.responseData,
		line.requestEtime.Format("2006/01/02 - 15:04:05"),
		line.CostTime)
}

// AcclogSetUp ...
func AcclogSetUp() gin.HandlerFunc {

	return func(c *gin.Context) {

		bodyLogWriter := &bodyLogWriter{body: bytes.NewBufferString(""), ResponseWriter: c.Writer}
		c.Writer = bodyLogWriter
		//开始时间
		startTime := time.Now()
		//处理请求
		c.Next()
		responseBody := bodyLogWriter.body.String()
		var responseCode int
		var responseMsg string
		var responseData interface{}
		if responseBody != "" {
			response := internal.Response{}
			err := json.Unmarshal([]byte(responseBody), &response)
			if err == nil {
				responseCode = response.Code
				responseMsg = response.Msg
				responseData = response.Data

			}
		}
		// 结束时间
		endTime := time.Now()
		if c.Request.Method == "POST" {
			c.Request.ParseForm()
		}
		//日志格式
		lineLog := new(line)
		lineLog.requestStime = startTime
		lineLog.requestMethod = c.Request.Method
		lineLog.requestURI = c.Request.RequestURI
		lineLog.requestProto = c.Request.Proto
		lineLog.requestUA = c.Request.UserAgent()
		lineLog.requestReferer = c.Request.Referer()
		lineLog.requestReferer = c.Request.PostForm.Encode()
		lineLog.requestReferer = c.ClientIP()
		lineLog.requestEtime = endTime
		lineLog.responseCode = responseCode
		lineLog.responseMsg = responseMsg
		lineLog.responseData = responseData
		lineLog.CostTime = fmt.Sprintf("%d µs", (endTime.UnixNano()-startTime.UnixNano())/1e3)
		fmt.Fprint(defaultWriter, logFormatter(lineLog))
	}
}
