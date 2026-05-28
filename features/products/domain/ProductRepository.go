package domain

import "fastcomp_api/features/products/domain/entities"

type ProductRepository interface {
	Save(product *entities.Product) (*entities.Product, error)
	GetByID(id string) (*entities.Product, error)
	GetAll() ([]*entities.Product, error)
	Update(product *entities.Product) error
	Delete(id string) error
	GetFeatured() ([]*entities.Product, error)
	GetByCategory(category entities.Category) ([]*entities.Product, error)
}
