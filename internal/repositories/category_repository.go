package repositories

import (
	"github.com/nvlhnn/gommerce/internal/models"
	"gorm.io/gorm"
)

type CategoryRepositoryInterface interface {
    Create(category *models.Category) error
    FindByID(id uint) (*models.Category, error)
    FindByName(name string) (*models.Category, error)
    Update(category *models.Category) error
    Delete(category *models.Category) error
}

type CategoryRepository struct {
    db *gorm.DB
}

func NewCategoryRepository(db *gorm.DB) CategoryRepositoryInterface {
    return &CategoryRepository{db}
}

func (r *CategoryRepository) Create(category *models.Category) error {
    result := r.db.Create(category)
    return result.Error
}

func (r *CategoryRepository) FindByID(id uint) (*models.Category, error) {
    var category models.Category
    result := r.db.First(&category, id)
    return &category, result.Error
}

func (r *CategoryRepository) FindByName(name string) (*models.Category, error) {
    var category models.Category
    result := r.db.Where("name = ?", name).First(&category)
    return &category, result.Error
}

func (r *CategoryRepository) Update(category *models.Category) error {
    result := r.db.Save(category)
    return result.Error
}

func (r *CategoryRepository) Delete(category *models.Category) error {
    result := r.db.Delete(category)
    return result.Error
}
