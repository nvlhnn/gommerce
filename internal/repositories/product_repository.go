package repositories

import (
	"net/url"
	"strconv"
	"strings"

	"github.com/nvlhnn/gommerce/internal/models"
	"gorm.io/gorm"
)

type ProductRepositoryInterface interface {
    Create(product *models.Product) error
    FindAll() ([]models.Product, error)
    FindByCategory(category_id uint, opts url.Values) ([]models.Product, int64, error)
	FindByID(id uint) (*models.Product, error)
    Update(product *models.Product) error
    Delete(product *models.Product) error
}

type ProductRepository struct {
    db *gorm.DB
}

func NewProductRepository(db *gorm.DB) ProductRepositoryInterface {
    return &ProductRepository{db}
}

func (r *ProductRepository) Create(product *models.Product) error {
    result := r.db.Create(product)
    return result.Error
}

func (r *ProductRepository) FindAll() ([]models.Product, error) {
    var products []models.Product
    result := r.db.Find(&products)
    return products, result.Error
}

func (r *ProductRepository) FindByCategory(category_id uint, options url.Values) ([]models.Product, int64, error) {
    
	var products []models.Product
	var total int64

	query := r.db.Where("category_id = ?", category_id).Where("stock > ?", 0).Limit(10).Preload("Category")

	if options.Get("sort") != "" {
		query.Order(options.Get("sort"))
	}

	if options.Get("search") != "" {
		query.Where("LOWER(name) LIKE ?", "%"+strings.ToLower(options.Get("search"))+"%")
	}

	// get total count and total pages
	query.Model(&models.Product{}).Count(&total)

	// pagination
	if options.Get("page") != "" {
		page, _ := strconv.Atoi(options.Get("page"))
		query.Offset((page - 1) * 10)
	}

	err := query.Find(&products).Error

    return products, total, err
}

func (r *ProductRepository) FindByID(id uint) (*models.Product, error) {
	var product models.Product
	result := r.db.Preload("Category").First(&product, id)
	return &product, result.Error
}

func (r *ProductRepository) Update(product *models.Product) error {
    result := r.db.Save(product)
    return result.Error
}

func (r *ProductRepository) Delete(product *models.Product) error {
    result := r.db.Delete(product)
    return result.Error
}
