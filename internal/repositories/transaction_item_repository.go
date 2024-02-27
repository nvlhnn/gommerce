package repositories

import (
	"github.com/nvlhnn/gommerce/internal/models"
	"gorm.io/gorm"
)


type TransactionItemRepositoryInterface interface {
    Create(transactionItem *models.TransactionItem) error
    FindByID(id uint) (*models.TransactionItem, error)
    Update(transactionItem *models.TransactionItem) error
    Delete(transactionItem *models.TransactionItem) error
}

type TransactionItemRepository struct {
    db *gorm.DB
}

func NewTransactionItemRepository(db *gorm.DB) TransactionItemRepositoryInterface {
    return &TransactionItemRepository{db}
}

func (r *TransactionItemRepository) Create(transactionItem *models.TransactionItem) error {
    result := r.db.Create(transactionItem)
    return result.Error
}

func (r *TransactionItemRepository) FindByID(id uint) (*models.TransactionItem, error) {
    var transactionItem models.TransactionItem
    result := r.db.First(&transactionItem, id)
    return &transactionItem, result.Error
}

func (r *TransactionItemRepository) Update(transactionItem *models.TransactionItem) error {
    result := r.db.Save(transactionItem)
    return result.Error
}

func (r *TransactionItemRepository) Delete(transactionItem *models.TransactionItem) error {
    result := r.db.Delete(transactionItem)
    return result.Error
}

