package services

import (
	"encoding/json"
	"errors"
	"fmt"

	"github.com/go-redis/redis/v8"
	"github.com/nvlhnn/gommerce/internal/dtos"
	"github.com/nvlhnn/gommerce/internal/models"
	"github.com/nvlhnn/gommerce/internal/repositories"
)

// CartService is a service that provides operations for shopping carts.
type CartServiceInterface interface {
	AddCart(customerID uint, cartDto dtos.AddCartRequest) ( dtos.CartDTO, error)
	ListCarts(customerID uint) ([]dtos.CartDTO, error)
	DeleteCart(productID uint, customerID uint) error
}

type CartService struct {
	cartRepo repositories.CartRepositoryInterface
	productRepo repositories.ProductRepositoryInterface

	// add redis cache
	cache *redis.Client
}

// NewCartService creates a new CartService.
func NewCartService(cartRepo repositories.CartRepositoryInterface, productRepo repositories.ProductRepositoryInterface, cache *redis.Client ) CartServiceInterface {
	return &CartService{cartRepo: cartRepo, productRepo: productRepo, cache: cache}
}


// AddCart adds a new cart to the database.
func (s *CartService) AddCart(customerID uint, cartDto dtos.AddCartRequest) (dtos.CartDTO, error) {

	cartRes := dtos.CartDTO{}

	// check if product exist
	product, err := s.productRepo.FindByID(cartDto.ProductID)
	if err != nil {
		return cartRes, err
	}

	if product.Stock < cartDto.Quantity {
		return cartRes, errors.New("product is out of stock")
	}
	
	// check if cart exist
	cartInDB, err := s.cartRepo.FindByProductID(cartDto.ProductID, customerID)
	if err != nil && err.Error() != "record not found"{
		return cartRes, err
	}

	if err == nil {
		cartInDB.Quantity = cartDto.Quantity
		err = s.cartRepo.Update(cartInDB)
		if err != nil {
			return cartRes, err
		}

		cartRes.ModelToDto(*cartInDB)
		return cartRes, nil
		
	}

	cart := &models.Cart{
		ProductID:  cartDto.ProductID,
		Quantity:   cartDto.Quantity,
		CustomerID: customerID,
	}
	err = s.cartRepo.Create(cart)

	if err != nil {
		return cartRes, err
	}
	cartRes.ModelToDto(*cart)

	// reset cache by customer id
	err = s.cache.HDel(s.cache.Context(), fmt.Sprint(customerID), "carts").Err()
	if err != nil && err != redis.Nil{
		return cartRes, err
	}

	return cartRes, nil

}

// ListCarts returns a list of carts in the database.
func (s *CartService) ListCarts(customerID uint) ([]dtos.CartDTO, error) {

	// get cache
	getSerielizedCarts, err := s.cache.HGet(s.cache.Context(), fmt.Sprint(customerID), "carts" ).Result()
	if err == nil {
		cardCached := make([]dtos.CartDTO, 0)
		err = json.Unmarshal([]byte(getSerielizedCarts), &cardCached)
		if err != nil {
			return nil, err
		}
		return cardCached, nil
	}


	carts, err := s.cartRepo.FindByCustomerID(customerID)
	if err != nil {
		return nil, err
	}

	cartsDto := make([]dtos.CartDTO, len(carts))
	for i, cart := range carts {
		cartsDto[i].ModelToDto(cart)
	}

	serializedCarts, err := json.Marshal(cartsDto)
	if err != nil {
		return nil, err
	}

	// set cache
	err = s.cache.HSet(s.cache.Context(), fmt.Sprint(customerID), "carts", serializedCarts).Err()
	if err != nil {
		return nil, err
	}

	return cartsDto, nil
}

// DeleteCart deletes a cart from the database.
func (s *CartService) DeleteCart(productID uint, customerID uint) error {
	cart, err := s.cartRepo.FindByProductID(productID, customerID)
	if err != nil {
		return err
	}

	err = s.cache.HDel(s.cache.Context(), fmt.Sprint(customerID), "carts").Err()
	if err != nil && err != redis.Nil{
		return err
	}

	return s.cartRepo.Delete(cart)
}

