package handlers

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/nvlhnn/gommerce/internal/dtos"
	"github.com/nvlhnn/gommerce/internal/services"
	"github.com/nvlhnn/gommerce/internal/utils/response"
)

type CustomerHandlerInterface interface {
	Login(ctx *gin.Context)
	Register(ctx *gin.Context)
}

type CustomerHandler struct {
	custmerService services.CustomerServiceInterface
	jwtService  services.JWTServiceInterface
}

func NewCustomerHandler(custmerService services.CustomerServiceInterface, jwtService services.JWTServiceInterface) CustomerHandlerInterface {
	return &CustomerHandler{
		custmerService: custmerService,
		jwtService:  jwtService,
	}
}

func (c *CustomerHandler) Login(ctx *gin.Context) {

	var loginDTO dtos.LoginDTO
	errDTO := ctx.ShouldBind(&loginDTO)
	if errDTO != nil {
        res := response.ClientResponse(http.StatusBadRequest, "Invalid request", nil, errDTO.Error())
        ctx.JSON(http.StatusBadRequest, res)
		return
	}

	customer, err := c.custmerService.VerifyCredential(loginDTO.Email, loginDTO.Password)
	if err != nil {
        res := response.ClientResponse(http.StatusBadRequest, "Invalid request" , nil, err.Error())
        ctx.JSON(http.StatusBadRequest, res)
		return
	}
	
	token, err := c.jwtService.GenerateToken(*customer)
	if err != nil {
        res := response.ClientResponse(http.StatusInternalServerError, "Internal server error" , nil, err.Error())
        ctx.JSON(http.StatusInternalServerError, res)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"token": token})
}

func (c *CustomerHandler) Register(ctx *gin.Context) {

	var registerDTO dtos.RegisterDTO
	errDTO := ctx.ShouldBind(&registerDTO)
	if errDTO != nil {
        res := response.ClientResponse(http.StatusBadRequest, "Invalid request", nil, errDTO.Error())
        ctx.JSON(http.StatusBadRequest, res)
		return
	}

	if !c.custmerService.IsDuplicateEmail(registerDTO.Email) {
        res := response.ClientResponse(http.StatusConflict, "Email already exist" , nil, errors.New("email already exist").Error())
        ctx.JSON(http.StatusInternalServerError, res)
		return
	} 

	createdUser, err := c.custmerService.CreateUser(registerDTO)
	if err != nil {
        res := response.ClientResponse(http.StatusInternalServerError, "Internal server error" , nil, err.Error())
        ctx.JSON(http.StatusInternalServerError, res)
		return
	}

	token, err := c.jwtService.GenerateToken(*createdUser)
	if err != nil {
        res := response.ClientResponse(http.StatusInternalServerError, "Internal server error" , nil, err.Error())
        ctx.JSON(http.StatusInternalServerError, res)
		return
	}


	res := response.ClientResponse(http.StatusCreated, "User has been created", gin.H{"token": token}, nil)
	ctx.JSON(http.StatusCreated, res)
	
	
}