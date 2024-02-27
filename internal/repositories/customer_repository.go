package repositories

import (
	"github.com/nvlhnn/gommerce/internal/models"
	"gorm.io/gorm"
)

type CustomerRepositoryInterface interface {
    Create(customer *models.Customer) error
    FindByID(id uint) (*models.Customer, error)
    FindByEmail(email string) (*models.Customer, error)
    Update(customer *models.Customer) error
    Delete(customer *models.Customer) error
}

type CustomerRepository struct {
    db *gorm.DB
}

func NewCustomerRepository(db *gorm.DB) CustomerRepositoryInterface {
    return &CustomerRepository{db}
}

func (r *CustomerRepository) Create(customer *models.Customer) error {
    result := r.db.Create(customer)
    return result.Error
}

func (r *CustomerRepository) FindByID(id uint) (*models.Customer, error) {
    var customer models.Customer
    result := r.db.First(&customer, id)
    return &customer, result.Error
}

func (r *CustomerRepository) FindByEmail(email string) (*models.Customer, error) {
    var customer models.Customer
    result := r.db.Where("email = ?", email).First(&customer)
    return &customer, result.Error
}

func (r *CustomerRepository) Update(customer *models.Customer) error {
    result := r.db.Save(customer)
    return result.Error
}

func (r *CustomerRepository) Delete(customer *models.Customer) error {
    result := r.db.Delete(customer)
    return result.Error
}
