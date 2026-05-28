package security

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func JWTMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		var tokenString string

		authHeader := c.GetHeader("Authorization")
		if strings.HasPrefix(authHeader, "Bearer ") {
			tokenString = strings.TrimPrefix(authHeader, "Bearer ")
		} else {
			cookie, err := c.Cookie("access_token")
			if err != nil || cookie == "" {
				c.JSON(http.StatusUnauthorized, gin.H{"error": "No autenticado"})
				c.Abort()
				return
			}
			tokenString = cookie
		}

		claims, err := ValidateJWT(tokenString)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Token inválido o expirado"})
			c.Abort()
			return
		}

		c.Set("user_id", claims.UserID)
		c.Set("userID", claims.UserID)
		c.Set("email", claims.Email)
		c.Set("role_id", claims.RoleID)
		c.Next()
	}
}

func RequireRole(requiredRoleID int) gin.HandlerFunc {
	return func(c *gin.Context) {
		roleID, exists := c.Get("role_id")
		if !exists {
			c.JSON(http.StatusForbidden, gin.H{"error": "Rol no encontrado en el token"})
			c.Abort()
			return
		}
		userRoleID, ok := roleID.(int)
		if !ok {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al procesar el rol"})
			c.Abort()
			return
		}
		if userRoleID != requiredRoleID {
			c.JSON(http.StatusForbidden, gin.H{"error": "No tienes permisos"})
			c.Abort()
			return
		}
		c.Next()
	}
}

func RequireAnyRole(allowedRoleIDs ...int) gin.HandlerFunc {
	return func(c *gin.Context) {
		roleID, exists := c.Get("role_id")
		if !exists {
			c.JSON(http.StatusForbidden, gin.H{"error": "Rol no encontrado en el token"})
			c.Abort()
			return
		}
		userRoleID, ok := roleID.(int)
		if !ok {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al procesar el rol"})
			c.Abort()
			return
		}
		for _, allowedRole := range allowedRoleIDs {
			if userRoleID == allowedRole {
				c.Next()
				return
			}
		}
		c.JSON(http.StatusForbidden, gin.H{"error": "No tienes permisos"})
		c.Abort()
	}
}

func OptionalJWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		var tokenString string

		authHeader := c.GetHeader("Authorization")
		if strings.HasPrefix(authHeader, "Bearer ") {
			tokenString = strings.TrimPrefix(authHeader, "Bearer ")
		} else {
			cookie, err := c.Cookie("access_token")
			if err != nil || cookie == "" {
				c.Next()
				return
			}
			tokenString = cookie
		}

		claims, err := ValidateJWT(tokenString)
		if err != nil {
			c.Next()
			return
		}

		c.Set("user_id", claims.UserID)
		c.Set("userID", claims.UserID)
		c.Set("email", claims.Email)
		c.Set("role_id", claims.RoleID)
		c.Next()
	}
}
