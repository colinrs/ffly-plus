package api

import (
	"net/http"

	"ffly-plus/internal"
	"ffly-plus/internal/code"

	"github.com/gin-gonic/gin"
)

// GetUserV1 ...
// @Summary Get multiple article tags
// @Produce  json
// @Param name query string false "Name"
// @Param state query int false "State"
// @Success 200 {object} app.Response
// @Failure 500 {object} app.Response
// @Router /api/v1/user [get]
func GetUserV1(c *gin.Context) {
	gin := internal.NewGin(c)
	gin.Response(http.StatusOK, code.SUCCESS, "GetUserV1")
}

// AddUserV1 ...
// @Summary Get multiple article tags
// @Produce  json
// @Param name query string false "Name"
// @Param state query int false "State"
// @Success 200 {object} app.Response
// @Failure 500 {object} app.Response
// @Router /api/v1/user [post]
func AddUserV1(c *gin.Context) {
	gin := internal.NewGin(c)
	gin.Response(http.StatusOK, code.SUCCESS, "AddUserV1")
}

// EditUserV1 ...
// @Summary Get multiple article tags
// @Produce  json
// @Param name query string false "Name"
// @Param state query int false "State"
// @Success 200 {object} app.Response
// @Failure 500 {object} app.Response
// @Router /api/v1/user [put]
func EditUserV1(c *gin.Context) {
	gin := internal.NewGin(c)
	gin.Response(http.StatusOK, code.SUCCESS, "EditUserV1")
}

// DeleteUserV1 ...
// @Summary Get multiple article tags
// @Produce  json
// @Param name query string false "Name"
// @Param state query int false "State"
// @Success 200 {object} app.Response
// @Failure 500 {object} app.Response
// @Router /api/v1/user [delete]
func DeleteUserV1(c *gin.Context) {
	gin := internal.NewGin(c)
	gin.Response(http.StatusOK, code.SUCCESS, "DeleteUserV1")
}
