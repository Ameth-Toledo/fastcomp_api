package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type VerifyTokenController struct{}

func NewVerifyTokenController() *VerifyTokenController {
	return &VerifyTokenController{}
}

func (vt *VerifyTokenController) Execute(c *gin.Context) {
	userID, _ := c.Get("user_id")
	email, _ := c.Get("email")
	roleID, _ := c.Get("role_id")

	c.JSON(http.StatusOK, gin.H{
		"authenticated": true,
		"user": gin.H{
			"id":     userID,
			"email":  email,
			"roleId": roleID,
		},
	})
}
