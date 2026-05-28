package main

import (
	"log"
	"os"

	authInfra "fastcomp_api/features/auth/infrastructure"
	authRoutes "fastcomp_api/features/auth/infrastructure/routes"
	productInfra "fastcomp_api/features/products/infrastructure"
	productRoutes "fastcomp_api/features/products/infrastructure/routes"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.Use(cors.New(cors.Config{
		AllowAllOrigins:  true,
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		AllowCredentials: false,
	}))

	auth := authInfra.InitAuth()
	authRoutes.ConfigureAuthRoutes(router, authRoutes.NewAuthDeps(
		auth.RegisterController,
		auth.LoginController,
		auth.LogoutController,
		auth.RefreshTokenController,
		auth.GetProfileController,
		auth.VerifyTokenController,
	))

	products := productInfra.InitProducts()
	productRoutes.ConfigureProductRoutes(router, products.ProductController)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	if err := router.Run(":" + port); err != nil {
		log.Fatal("Error al iniciar el servidor: ", err)
	}
}
