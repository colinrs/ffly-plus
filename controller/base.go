package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Server ...
type Server struct {
	GinEngine *gin.Engine
}

// Health ...
func Health(c *gin.Context) {
	c.String(http.StatusOK, "ok")
}
