package controllers

import (
	"fastcomp_api/core/security"
	"net/http"

	"github.com/gin-gonic/gin"
)

type LogoutController struct{}

func NewLogoutController() *LogoutController {
	return &LogoutController{}
}

func (lc *LogoutController) Execute(c *gin.Context) {
	security.ClearAuthCookies(c.Writer)
	c.JSON(http.StatusOK, gin.H{"message": "Logout exitoso"})
}
