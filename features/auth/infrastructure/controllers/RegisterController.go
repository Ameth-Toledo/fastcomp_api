package controllers

import (
	"fastcomp_api/features/auth/application"
	"fastcomp_api/features/auth/domain/dto"
	"fastcomp_api/features/auth/domain/entities"
	"net/http"

	"github.com/gin-gonic/gin"
)

type RegisterController struct {
	useCase *application.Register
}

func NewRegisterController(useCase *application.Register) *RegisterController {
	return &RegisterController{useCase: useCase}
}

func (rc *RegisterController) Execute(c *gin.Context) {
	var req dto.CreateUserRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user := &entities.User{
		FirstName:    req.FirstName,
		LastName:     req.LastName,
		BusinessName: req.BusinessName,
		Email:        req.Email,
		Password:     &req.Password,
		Phone:        req.Phone,
		Website:      req.Website,
		ProfilePhoto: req.ProfilePhoto,
	}

	created, err := rc.useCase.Execute(user)
	if err != nil {
		c.JSON(http.StatusConflict, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "Usuario registrado exitosamente",
		"user":    application.ToUserResponse(created),
	})
}
