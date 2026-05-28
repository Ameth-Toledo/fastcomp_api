package controllers

import (
	"fastcomp_api/features/auth/application"
	"net/http"

	"github.com/gin-gonic/gin"
)

type GetProfileController struct {
	useCase *application.GetUserByID
}

func NewGetProfileController(useCase *application.GetUserByID) *GetProfileController {
	return &GetProfileController{useCase: useCase}
}

func (gp *GetProfileController) Execute(c *gin.Context) {
	userID, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "No autenticado"})
		return
	}

	user, err := gp.useCase.Execute(userID.(int))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"user": application.ToUserResponse(user)})
}
