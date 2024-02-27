package repositories

import (
	"github.com/nvlhnn/gommerce/internal/models"
	"gorm.io/gorm"
)

// interface
type TransactionRepositoryInterface interface {
	CreateOrder(order *models.Transaction) error
    GetOrderByCustomerID(customerID uint) ([]models.Transaction, error)
	GetOrderByID(id uint) (*models.Transaction, error)
	UpdateOrder(order *models.Transaction) error
	DeleteOrder(order *models.Transaction) error
}

type TransactionRepository struct {
    db *gorm.DB
}

func NewTransactionRepository(db *gorm.DB) TransactionRepositoryInterface  {
    return &TransactionRepository{db}
}

func (r *TransactionRepository) CreateOrder(order *models.Transaction) error {
    result := r.db.Create(order)
    return result.Error
}

func (r *TransactionRepository) GetOrderByID(id uint) (*models.Transaction, error) {
    var order models.Transaction
    result := r.db.First(&order, id)
    if result.Error != nil {
        return nil, result.Error
    }
    return &order, nil
}

func (r *TransactionRepository) GetOrderByCustomerID(customerID uint) ([]models.Transaction, error) {
    var orders []models.Transaction
    result := r.db.Preload("Customer").Preload("Items.Product.Category").Where("customer_id = ?", customerID).Order("transaction_date DESC").Find(&orders)
    if result.Error != nil {
        return nil, result.Error
    }
    return orders, nil
}

func (r *TransactionRepository) UpdateOrder(order *models.Transaction) error {
    result := r.db.Save(order)
    return result.Error
}

func (r *TransactionRepository) DeleteOrder(order *models.Transaction) error {
    result := r.db.Delete(order)
    return result.Error
}
