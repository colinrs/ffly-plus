package controller

import (
	"net/http"

	"github.com/colinrs/ffly-plus/internal"
	"github.com/colinrs/ffly-plus/internal/version"

	"github.com/gin-gonic/gin"
)

// Health ...
func Health(c *gin.Context) {
	c.String(http.StatusOK, "ok")
}

// Version ...
func Version(c *gin.Context) {
	internal.APIResponse(c, nil, version.Get())
}
