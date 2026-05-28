package infrastructure

import (
	"fastcomp_api/core"
	"fastcomp_api/features/auth/application"
	"fastcomp_api/features/auth/infrastructure/adapters"
	"fastcomp_api/features/auth/infrastructure/controllers"
)

type DependenciesAuth struct {
	RegisterController     *controllers.RegisterController
	LoginController        *controllers.LoginController
	LogoutController       *controllers.LogoutController
	RefreshTokenController *controllers.RefreshTokenController
	GetProfileController   *controllers.GetProfileController
	VerifyTokenController  *controllers.VerifyTokenController
}

func InitAuth() *DependenciesAuth {
	conn := core.GetDBPool()
	userRepo := adapters.NewPostgreSQL(conn.DB)

	login := application.NewLogin(userRepo)
	register := application.NewRegister(userRepo)
	getUserByID := application.NewGetUserByID(userRepo)

	return &DependenciesAuth{
		RegisterController:     controllers.NewRegisterController(register),
		LoginController:        controllers.NewLoginController(login),
		LogoutController:       controllers.NewLogoutController(),
		RefreshTokenController: controllers.NewRefreshTokenController(getUserByID),
		GetProfileController:   controllers.NewGetProfileController(getUserByID),
		VerifyTokenController:  controllers.NewVerifyTokenController(),
	}
}
