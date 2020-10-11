package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Health ...
func Health(c *gin.Context) {
	c.String(http.StatusOK, "ok")
}
