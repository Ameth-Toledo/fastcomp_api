package infrastructure

import (
	"fastcomp_api/core"
	"fastcomp_api/features/products/application"
	"fastcomp_api/features/products/infrastructure/adapters"
	"fastcomp_api/features/products/infrastructure/controllers"
)

type DependenciesProducts struct {
	ProductController *controllers.ProductController
}

func InitProducts() *DependenciesProducts {
	conn := core.GetDBPool()
	repo := adapters.NewPostgreSQL(conn.DB)
	service := application.NewProductService(repo)

	return &DependenciesProducts{
		ProductController: controllers.NewProductController(service),
	}
}
