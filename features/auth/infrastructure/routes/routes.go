package routes

import (
	"fastcomp_api/core/security"
	"fastcomp_api/features/auth/infrastructure/controllers"

	"github.com/gin-gonic/gin"
)

func ConfigureAuthRoutes(router *gin.Engine, deps *authDeps) {
	auth := router.Group("/auth")
	{
		auth.POST("/register", deps.register.Execute)
		auth.POST("/login", deps.login.Execute)
		auth.POST("/logout", deps.logout.Execute)
		auth.POST("/refresh", deps.refresh.Execute)
		auth.GET("/profile", security.JWTMiddleware(), deps.profile.Execute)
		auth.GET("/verify", security.JWTMiddleware(), deps.verify.Execute)
	}
}

type authDeps struct {
	register *controllers.RegisterController
	login    *controllers.LoginController
	logout   *controllers.LogoutController
	refresh  *controllers.RefreshTokenController
	profile  *controllers.GetProfileController
	verify   *controllers.VerifyTokenController
}

func NewAuthDeps(
	register *controllers.RegisterController,
	login *controllers.LoginController,
	logout *controllers.LogoutController,
	refresh *controllers.RefreshTokenController,
	profile *controllers.GetProfileController,
	verify *controllers.VerifyTokenController,
) *authDeps {
	return &authDeps{register, login, logout, refresh, profile, verify}
}
