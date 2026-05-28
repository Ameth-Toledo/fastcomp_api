package controllers

import (
	"fastcomp_api/core/security"
	"fastcomp_api/features/auth/application"
	"fastcomp_api/features/auth/domain/dto"
	"net/http"

	"github.com/gin-gonic/gin"
)

type LoginController struct {
	useCase *application.Login
}

func NewLoginController(useCase *application.Login) *LoginController {
	return &LoginController{useCase: useCase}
}

func (lc *LoginController) Execute(c *gin.Context) {
	var req dto.LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err := lc.useCase.Execute(req.Email, req.Password)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	accessToken, err := security.GenerateJWT(user.ID, user.Email, user.RoleID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al generar token"})
		return
	}

	refreshToken, err := security.GenerateRefreshToken(user.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al generar refresh token"})
		return
	}

	security.SetAuthCookie(c.Writer, accessToken)
	security.SetRefreshCookie(c.Writer, refreshToken)

	c.JSON(http.StatusOK, gin.H{
		"message": "Login exitoso",
		"user":    application.ToUserResponse(user),
		"token":   accessToken,
	})
}
