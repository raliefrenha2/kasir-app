package services

import (
	"kasir-api/models"
	"kasir-api/repositories"
)

type CategoryService struct {
	repo *repositories.CategoryRepository
}

func NewCategoryService(repo *repositories.CategoryRepository) *CategoryService {
	return &CategoryService{repo: repo}
}

func (s *CategoryService) GetAll() ([]models.Category, error) {
	return s.repo.GetAll()
}

func (s *CategoryService) GetByID(id int) (models.Category, error) {
	return s.repo.GetByID(id)
}

func (s *CategoryService) Create(category models.Category) (models.Category, error) {
	return s.repo.Create(category)
}

func (s *CategoryService) Update(id int, category models.Category) (models.Category, error) {
	category.ID = id
	err := s.repo.Update(category)
	if err != nil {
		return models.Category{}, err
	}
	return category, nil
}

func (s *CategoryService) Delete(id int) error {
	return s.repo.Delete(id)
}
