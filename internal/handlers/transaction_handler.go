package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/nvlhnn/gommerce/internal/services"
	"github.com/nvlhnn/gommerce/internal/utils/response"
)

type TransactionHandlerInterface interface {
	CreateOrder(ctx *gin.Context)
	GetOrders(ctx *gin.Context)
	// UpdateOrder(ctx *gin.Context)
	// DeleteOrder(ctx *gin.Context)
}

type TransactionHandler struct {
	transactionService services.TransactionServiceInterface
}

func NewTransactionHandler(transactionService services.TransactionServiceInterface) TransactionHandlerInterface {
	return &TransactionHandler{transactionService: transactionService}
}

func (t *TransactionHandler) CreateOrder(ctx *gin.Context) {

	customerID := ctx.MustGet("customerId").(uint)

	err := t.transactionService.CreateOrder(customerID)
	if err != nil {
        res := response.ClientResponse(http.StatusBadRequest , "failed to create order", nil, err.Error())
        ctx.JSON(http.StatusBadRequest, res)
		return
	}

	res := response.ClientResponse(http.StatusCreated, "order created", nil, nil)
	ctx.JSON(http.StatusCreated, res)

}


func (t *TransactionHandler) GetOrders(ctx *gin.Context) {
	
	customerID := ctx.MustGet("customerId").(uint)

	orders, err := t.transactionService.GetCustomerOrders(customerID)
	if err != nil {
		res := response.ClientResponse(http.StatusBadRequest , "failed to get orders", nil, err.Error())
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	res := response.ClientResponse(http.StatusOK, "orders found", orders, nil)
	ctx.JSON(http.StatusOK, res)

}

