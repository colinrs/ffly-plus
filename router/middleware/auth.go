package middleware

import (
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
			//fmt.Print("xxxxxxxxxxxxxxxxxxxx\n")
			return
		}

		// set uid to context
		c.Set("uid", ctx.UserID)

		c.Next()
	}
}
