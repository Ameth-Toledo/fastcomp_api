package application

import (
	"crypto/rand"
	"encoding/hex"
	"errors"
	"fastcomp_api/features/products/domain"
	"fastcomp_api/features/products/domain/entities"
	"time"
)

type ProductService struct {
	repo domain.ProductRepository
}

func NewProductService(repo domain.ProductRepository) *ProductService {
	return &ProductService{repo: repo}
}

func generateID() string {
	b := make([]byte, 4)
	rand.Read(b)
	return "p" + hex.EncodeToString(b)
}

func (s *ProductService) Create(p *entities.Product) (*entities.Product, error) {
	p.ID = generateID()
	p.AddedAt = time.Now().Format("2006-01-02")
	if p.Specs == nil {
		p.Specs = []string{}
	}
	return s.repo.Save(p)
}

func (s *ProductService) GetByID(id string) (*entities.Product, error) {
	product, err := s.repo.GetByID(id)
	if err != nil {
		return nil, err
	}
	if product == nil {
		return nil, errors.New("producto no encontrado")
	}
	return product, nil
}

func (s *ProductService) GetAll() ([]*entities.Product, error) {
	return s.repo.GetAll()
}

func (s *ProductService) Update(id string, p *entities.Product) (*entities.Product, error) {
	existing, err := s.repo.GetByID(id)
	if err != nil || existing == nil {
		return nil, errors.New("producto no encontrado")
	}
	p.ID = id
	p.AddedAt = existing.AddedAt
	if p.Specs == nil {
		p.Specs = []string{}
	}
	if err := s.repo.Update(p); err != nil {
		return nil, err
	}
	return p, nil
}

func (s *ProductService) Delete(id string) error {
	existing, err := s.repo.GetByID(id)
	if err != nil || existing == nil {
		return errors.New("producto no encontrado")
	}
	return s.repo.Delete(id)
}

func (s *ProductService) GetFeatured() ([]*entities.Product, error) {
	return s.repo.GetFeatured()
}

func (s *ProductService) GetByCategory(category entities.Category) ([]*entities.Product, error) {
	return s.repo.GetByCategory(category)
}
