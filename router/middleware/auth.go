package middleware

import (
	"net/http"
	"time"

	"ffly-plus/pkg/token"

	"github.com/gin-gonic/gin"
)

// AuthMiddleware 认证中间件
func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Parse the json web token.

		ctx, err := token.ParseRequest(c)
		if err != nil {
			c.Abort()
			c.JSON(http.StatusOK, err.Error())
			return
		}

		if ctx.ExpirationTime > time.Now().Unix() {
			c.Abort()
			c.JSON(http.StatusOK, "Token Expiration")
			return
		}
		// set uid to context
		c.Set("uid", ctx.UserID)

		c.Next()
	}
}
