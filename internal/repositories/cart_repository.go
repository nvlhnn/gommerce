package repositories

import (
	// using gorm to interact with the database

	"github.com/nvlhnn/gommerce/internal/models"
	"gorm.io/gorm"
)

type CartRepositoryInterface interface {
	Create(cart *models.Cart)  error
	FindByCustomerID(customerID uint) ([]models.Cart, error)
	Delete(cart *models.Cart) error
	FindByProductID(productID uint, customerID uint) (*models.Cart, error)
	Update(cart *models.Cart) (error)
}

type CartRepository struct {
    db *gorm.DB
}

func NewCartRepository(db *gorm.DB) CartRepositoryInterface {
	return &CartRepository{db}
}

func (r *CartRepository) Create(cart *models.Cart) error {
	result := r.db.Preload("Product.Category").Create(cart).Find(cart)
	return result.Error
}

func (r *CartRepository) FindByCustomerID(customerID uint) ([]models.Cart, error) {
	var carts []models.Cart
	result := r.db.Preload("Product.Category").Where("customer_id = ?", customerID).Find(&carts)
	return carts, result.Error
}

func (r *CartRepository) FindByProductID(productID uint, customerID uint) (*models.Cart, error) {
	var cart models.Cart
	result := r.db.Preload("Product.Category").Where("product_id = ?", productID).Where("customer_id = ?", customerID).First(&cart)
	return &cart, result.Error
}

func (r *CartRepository) Update(cart *models.Cart) ( error) {
	result := r.db.Preload("Product.Category").Save(cart)
	return result.Error
}

func (r *CartRepository) Delete(cart *models.Cart) error {
	result := r.db.Delete(cart)
	return result.Error
}