package services

import (
	"net/url"
	"strconv"

	"github.com/nvlhnn/gommerce/internal/dtos"
	"github.com/nvlhnn/gommerce/internal/models"
	"github.com/nvlhnn/gommerce/internal/repositories"
)

// product interface
type ProductServiceInterface interface {
	CreateProduct(product *models.Product) error
	GetAllProducts() ([]models.Product, error)
	GetProductsByCategory(category_id uint, opts url.Values) ([]dtos.ProductDTO, dtos.Metadata, error)
	GetProductByID(id uint) (*models.Product, error)
	UpdateProduct(product *models.Product) error
	DeleteProduct(product *models.Product) error
}


// product service
type ProductService struct {
	productRepository repositories.ProductRepositoryInterface
	categoryRepo repositories.CategoryRepositoryInterface
}

// create new product service
func NewProductService(productRepo repositories.ProductRepositoryInterface, categoryRepo repositories.CategoryRepositoryInterface ) ProductServiceInterface {
	return &ProductService{productRepository: productRepo, categoryRepo: categoryRepo}
}

// create product
func (s *ProductService) CreateProduct(product *models.Product) error {
	return s.productRepository.Create(product)
}

// get all products
func (s *ProductService) GetAllProducts() ([]models.Product, error) {
	return s.productRepository.FindAll()
}

// get products by category
func (s *ProductService) GetProductsByCategory(category_id uint, options url.Values) ([]dtos.ProductDTO, dtos.Metadata, error) {

	metadata := dtos.Metadata{}

	// validate category id exist
	_, err := s.categoryRepo.FindByID(category_id)
	if err != nil {
		return nil, metadata, err
	}


	// get products by category
	products, total, err := s.productRepository.FindByCategory(category_id, options)
	if err != nil {
		return nil, metadata, err
	}

	// format products
	productDTOs := make([]dtos.ProductDTO, len(products))
	for i, product := range products {
		productDTOs[i].ModelToDto(product)
	}

	page, err := strconv.Atoi(options.Get("page"))
    if err != nil || page <= 0 {
        page = 1 
    }

	// create metadata
	metadata.TotalData = total
	metadata.LastPage = int((total + dtos.PageSize - 1) / dtos.PageSize)
	metadata.PerPage = dtos.PageSize
	metadata.CurrentPage = page

	return productDTOs, metadata, nil
}

// get product by id
func (s *ProductService) GetProductByID(id uint) (*models.Product, error) {
	return s.productRepository.FindByID(id)
}

// update product
func (s *ProductService) UpdateProduct(product *models.Product) error {
	return s.productRepository.Update(product)
}

// delete product
func (s *ProductService) DeleteProduct(product *models.Product) error {
	return s.productRepository.Delete(product)
}

