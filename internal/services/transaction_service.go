package services

import (
	"errors"
	"fmt"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/nvlhnn/gommerce/internal/dtos"
	"github.com/nvlhnn/gommerce/internal/models"
	"github.com/nvlhnn/gommerce/internal/repositories"
	"gorm.io/gorm"
)

type TransactionServiceInterface interface {
	CreateOrder(customerID uint) error
	GetCustomerOrders(customerID uint) ([]dtos.TransactionDTO, error)
	GetOrderByID(id uint) (*models.Transaction, error)
	UpdateOrder(order *models.Transaction) error
	DeleteOrder(order *models.Transaction) error
}

type TransactionService struct {
	transactionRepo repositories.TransactionRepositoryInterface
	cartRepo 	  repositories.CartRepositoryInterface
	db 		  *gorm.DB
	cache 	  *redis.Client
}

func NewTransactionService(db *gorm.DB, cache  *redis.Client, trxRepo repositories.TransactionRepositoryInterface, cartRepo repositories.CartRepositoryInterface) TransactionServiceInterface {
	return &TransactionService{transactionRepo: trxRepo, cartRepo: cartRepo, db: db}
}

func (s *TransactionService) CreateOrder(customerID uint) error {

	// create transaction
	tx := s.db.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	if err := tx.Error; err != nil {
		return err
	}

	// get all items in the cart
	carts, err := s.cartRepo.FindByCustomerID(customerID)
	if err != nil {
		return err
	}

	// check if cart empty
	if len(carts) == 0 {
		return errors.New("cart is empty")
	}

	// check for each product in the cart if the stock is enough
	for _, cart := range carts {
		if cart.Product.Stock < cart.Quantity {
			return errors.New("stock is not enough")
		}
	}

	// create a new order
	order := &models.Transaction{
		CustomerID: customerID,
		TransactionDate: time.Now(),
		Items:      make([]models.TransactionItem, len(carts)),
	}

	var totalAmount float64 = 0

	// copy items from the cart to the order
	for i, cart := range carts {
		order.Items[i] = models.TransactionItem{
			ProductID: cart.ProductID,
			Quantity:  cart.Quantity,
			UnitPrice:     cart.Product.Price,
		}
		totalAmount += cart.Product.Price * float64(cart.Quantity)
	}

	// create order
	order.TotalAmount = totalAmount
	err = s.transactionRepo.CreateOrder(order)
	if err != nil {
		tx.Rollback()
		return err
	}

	// delete all items in the cart
	for _, cart := range carts {
		err = s.cartRepo.Delete(&cart)
		if err != nil {
			tx.Rollback()
			return err
		}

		// reduce stock
		product := cart.Product
		product.Stock -= cart.Quantity
		err = s.db.Save(&product).Error
		if err != nil {
			tx.Rollback()
			return err
		}
	}

	// delete cache
	err = s.cache.HDel(s.cache.Context(), fmt.Sprint(customerID), "carts").Err()
	if err != nil && err != redis.Nil{
		return err
	}

	return tx.Commit().Error
}

func (s *TransactionService) GetCustomerOrders(customerID uint) ([]dtos.TransactionDTO, error) {

	orders, err := s.transactionRepo.GetOrderByCustomerID(customerID)
	if err != nil {
		return nil, err
	}

	// convert to dto
	var ordersDTO []dtos.TransactionDTO
	for _, order := range orders {
		var orderDTO dtos.TransactionDTO
		orderDTO.ModelToDto(order)
		ordersDTO = append(ordersDTO, orderDTO)
	}

	return ordersDTO, nil

}

func (s *TransactionService) GetOrderByID(id uint) (*models.Transaction, error) {
	return s.transactionRepo.GetOrderByID(id)
}


func (s *TransactionService) UpdateOrder(order *models.Transaction) error {
	return s.transactionRepo.UpdateOrder(order)
}

func (s *TransactionService) DeleteOrder(order *models.Transaction) error {
	return s.transactionRepo.DeleteOrder(order)
}


