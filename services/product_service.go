package services

import (
	"kasir-api/models"
	"kasir-api/repositories"
)

type ProductService struct {
	repo *repositories.ProductRepository
}

func NewProductService(repo *repositories.ProductRepository) *ProductService {
	return &ProductService{repo: repo}
}

func (s *ProductService) GetAll() ([]models.Product, error) {
	return s.repo.GetAll()
}

func (s *ProductService) GetByID(id int) (models.Product, error) {
	return s.repo.GetByID(id)
}

func (s *ProductService) Create(product *models.Product) (*models.Product, error) {
	err := s.repo.Create(product)
	if err != nil {
		return nil, err
	}
	return product, nil
}

func (s *ProductService) Update(id int, product models.Product) (*models.Product, error) {
	product.ID = id
	err := s.repo.Update(&product)
	if err != nil {
		return nil, err
	}
	return &product, nil
}

func (s *ProductService) Delete(id int) error {
	return s.repo.Delete(id)
}
