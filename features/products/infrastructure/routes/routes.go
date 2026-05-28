package routes

import (
	"fastcomp_api/core/security"
	"fastcomp_api/features/products/infrastructure/controllers"

	"github.com/gin-gonic/gin"
)

func ConfigureProductRoutes(router *gin.Engine, ctrl *controllers.ProductController) {
	products := router.Group("/products")
	{
		products.GET("", ctrl.GetAll)
		products.GET("/:id", ctrl.GetByID)

		protected := products.Group("")
		protected.Use(security.JWTMiddleware())
		{
			protected.POST("", ctrl.Create)
			protected.PUT("/:id", ctrl.Update)
			protected.DELETE("/:id", ctrl.Delete)
		}
	}
}
