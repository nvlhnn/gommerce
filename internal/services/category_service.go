package services

import (
	"github.com/nvlhnn/gommerce/internal/models"
	"github.com/nvlhnn/gommerce/internal/repositories"
)

type CategoryServiceInterface interface {
    CreateCategory(category *models.Category) error
    GetCategoryByID(id uint) (*models.Category, error)
    GetCategoryByName(name string) (*models.Category, error)
    UpdateCategory(category *models.Category) error
    DeleteCategory(category *models.Category) error
}

type CategoryService struct {
    categoryRepository repositories.CategoryRepositoryInterface
}

func NewCategoryService(categoryRepo repositories.CategoryRepositoryInterface) CategoryServiceInterface {
    return &CategoryService{categoryRepository: categoryRepo}
}

func (s *CategoryService) CreateCategory(category *models.Category) error {    
    return s.categoryRepository.Create(category)
}

func (s *CategoryService) GetCategoryByID(id uint) (*models.Category, error) {
    return s.categoryRepository.FindByID(id)
}

func (s *CategoryService) GetCategoryByName(name string) (*models.Category, error) {
    return s.categoryRepository.FindByName(name)
}

func (s *CategoryService) UpdateCategory(category *models.Category) error {    
    return s.categoryRepository.Update(category)
}

func (s *CategoryService) DeleteCategory(category *models.Category) error {    
    return s.categoryRepository.Delete(category)
}
