package middlewares

import (
	"errors"
	"log"
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/nvlhnn/gommerce/internal/services"
	"github.com/nvlhnn/gommerce/internal/utils/response"
)

//AuthorizeJWT validates the token user given, return 401 if not valid
func AuthorizeJWT(jwtService services.JWTServiceInterface, customerService services.CustomerServiceInterface) gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			res := response.ClientResponse(http.StatusUnauthorized, "Unauthorized", nil, errors.New("Unauthorized").Error())
			c.AbortWithStatusJSON(http.StatusUnauthorized, res)
			return
		}
		token, err := jwtService.ValidateToken(authHeader)
		if token.Valid {
			claims := token.Claims.(jwt.MapClaims)
			
			id, ok := claims["customer_id"].(float64)
			if !ok || id == 0 {
				res := response.ClientResponse(http.StatusUnauthorized, "error in retrieving id", nil, errors.New("error in retrieving id").Error())
				c.AbortWithStatusJSON(http.StatusUnauthorized, res)
				return
			}

			// check customer id exist
			_, err := customerService.FindByID(uint(id))
			if err != nil {
				res := response.ClientResponse(http.StatusUnauthorized, "invalid token", nil, err.Error())
				c.AbortWithStatusJSON(http.StatusUnauthorized, res)
				return
			}

			c.Set("customerId", uint(id))
			log.Print(claims)
		} else {
			res := response.ClientResponse(http.StatusUnauthorized, "Unauthorized", nil, err.Error())
			c.AbortWithStatusJSON(http.StatusUnauthorized, res)		
		}
	}
}