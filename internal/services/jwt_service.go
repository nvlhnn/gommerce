package services

import (
	"fmt"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/nvlhnn/gommerce/internal/models"
)

type JWTServiceInterface interface {
	GenerateToken(customer models.Customer) (string, error)
	ValidateToken(token string) (*jwt.Token, error)
}

type JWTClaim struct{
	jwt.StandardClaims
	CustomerID uint `json:"customer_id"`
}

type jwtService struct{
	secretKey string
	issuer string
}

func NewJWTService() JWTServiceInterface{
	return &jwtService{
		issuer: "nvlhnn",
		secretKey: getSecretKey(),
	}
}


func (s *jwtService) GenerateToken(customer models.Customer) (string, error){
	claims := &JWTClaim{
		CustomerID: customer.ID,
		StandardClaims: jwt.StandardClaims{
			Issuer:    s.issuer,
			IssuedAt: time.Now().Unix(),
			ExpiresAt: time.Now().Add(time.Hour * 24).Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256,claims)
	t, err := token.SignedString([]byte(s.secretKey))
	if err != nil{
		
		return "", err
	}

	return 	t, nil
} 


func (s *jwtService) ValidateToken(token string) (*jwt.Token, error) {

	return jwt.Parse(token, func(t_ *jwt.Token) (interface{}, error) {
		if _, ok := t_.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method %v", t_.Header["alg"])
		}
		return []byte(s.secretKey), nil
	})
}


func getSecretKey() string {
	secreKey := os.Getenv("JWT_SECRET")
	if secreKey== "" {
		secreKey = "nvlhnn"
	}
	return secreKey
}