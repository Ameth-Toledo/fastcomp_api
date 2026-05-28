package controllers

import (
	"fastcomp_api/core/security"
	"fastcomp_api/features/auth/application"
	"net/http"

	"github.com/gin-gonic/gin"
)

type RefreshTokenController struct {
	useCase *application.GetUserByID
}

func NewRefreshTokenController(useCase *application.GetUserByID) *RefreshTokenController {
	return &RefreshTokenController{useCase: useCase}
}

func (rc *RefreshTokenController) Execute(c *gin.Context) {
	refreshToken, err := c.Cookie("refresh_token")
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Refresh token no encontrado"})
		return
	}

	claims, err := security.ValidateRefreshToken(refreshToken)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Refresh token inválido"})
		return
	}

	user, err := rc.useCase.Execute(claims.UserID)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Usuario no encontrado"})
		return
	}

	newAccessToken, err := security.GenerateJWT(user.ID, user.Email, user.RoleID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al generar token"})
		return
	}

	security.SetAuthCookie(c.Writer, newAccessToken)
	c.JSON(http.StatusOK, gin.H{"message": "Token renovado"})
}
