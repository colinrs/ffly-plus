package middleware

import (
	"net/http"

	sentinelAPI "github.com/alibaba/sentinel-golang/api"
	"github.com/alibaba/sentinel-golang/core/base"
	"github.com/colinrs/pkgx/logger"
	"github.com/gin-gonic/gin"
)

// SentinelMiddleware 认证中间件
func SentinelMiddleware() gin.HandlerFunc {

	return func(c *gin.Context) {
		fullPath := c.FullPath()
		if fullPath == "" {
			fullPath = "/"
		}

		sentinelEntry, err := sentinelAPI.Entry(fullPath, sentinelAPI.WithTrafficType(base.Inbound))
		if err != nil {
			c.Abort()
			logger.Error("fullPath:%s err:(%#v)\n", fullPath, err)
			c.JSON(http.StatusOK, err.Error())
			return
		}
		sentinelEntry.Exit()
		c.Next()

	}
}
