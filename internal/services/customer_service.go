package services

import (
	"errors"
	"log"

	"github.com/nvlhnn/gommerce/internal/dtos"
	"github.com/nvlhnn/gommerce/internal/models"
	"github.com/nvlhnn/gommerce/internal/repositories"

	"golang.org/x/crypto/bcrypt"
)

type CustomerServiceInterface interface {
	VerifyCredential(email string, password string) (*models.Customer, error)
	CreateUser(customerDto dtos.RegisterDTO) (*models.Customer, error)
	FindByEmail(email string) (*models.Customer, error)
	IsDuplicateEmail(email string) bool
	FindByID(ID uint) (*models.Customer, error)

}

type CustomerService struct {
	customerRepository repositories.CustomerRepositoryInterface
}

//NewAuthService creates a new instance of AuthService
func NewCustomerService(userRep repositories.CustomerRepositoryInterface) CustomerServiceInterface {
	return &CustomerService{
		customerRepository: userRep,
	}
}

func (service *CustomerService) VerifyCredential(email string, password string) (*models.Customer, error) {

	res, err := service.customerRepository.FindByEmail(email)
	if err != nil {
		if err == errors.New("record not found"){
			return res, errors.New("invalid email or password")
		}
		return res, err
	}
	
	comparedPassword := comparePassword(res.Password, []byte(password))
	if res.Email == email && comparedPassword {
		return res, nil
	}

	return res, errors.New("invalid email or password")
}

func (service *CustomerService) FindByID(ID uint) (*models.Customer, error) {
	return service.customerRepository.FindByID(ID)
}

func (service *CustomerService) CreateUser(customerDto dtos.RegisterDTO) (*models.Customer, error) {

	hashedPass := hashAndSalt([]byte(customerDto.Password))

	customer := models.Customer{
		Name: customerDto.Name,
		Email: customerDto.Email,
		Password: hashedPass,
	}
	err := service.customerRepository.Create(&customer)
	return &customer, err
}

func (service *CustomerService) FindByEmail(email string)( *models.Customer, error ){
	return service.customerRepository.FindByEmail(email)
}

func (service *CustomerService) IsDuplicateEmail(email string) bool {
	_, err := service.customerRepository.FindByEmail(email)
	return !(err == nil)
}

func comparePassword(hashedPwd string, plainPassword []byte) bool {
	byteHash := []byte(hashedPwd)
	err := bcrypt.CompareHashAndPassword(byteHash, plainPassword)
	if err != nil {
		log.Println(err)
		return false
	}
	return true
}

func hashAndSalt(pwd []byte) string {
	hash, err := bcrypt.GenerateFromPassword(pwd, bcrypt.MinCost)
	if err != nil {
		log.Println(err)
	}
	return string(hash)
}


