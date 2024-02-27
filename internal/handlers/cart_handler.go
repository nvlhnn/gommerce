package handlers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/nvlhnn/gommerce/internal/dtos"
	"github.com/nvlhnn/gommerce/internal/services"
	"github.com/nvlhnn/gommerce/internal/utils/response"
)

type CartHandlerInterface interface {
	AddCart(ctx *gin.Context)
	ListCarts(ctx *gin.Context)
	DeleteCart(ctx *gin.Context)
}

type CartHanlder struct {
    cartService services.CartServiceInterface
}

func NewCartHandler(cartService services.CartServiceInterface) CartHandlerInterface {
    return &CartHanlder{cartService: cartService}
}

func (c *CartHanlder) AddCart(ctx *gin.Context) {
	customerID := ctx.MustGet("customerId").(uint)

	var cardRequest dtos.AddCartRequest
    if err := ctx.BindJSON(&cardRequest); err != nil {
        res := response.ClientResponse(http.StatusBadRequest, "Invalid request", nil, err.Error())
        ctx.JSON(http.StatusBadRequest, res)
        return
    }

    cart, err := c.cartService.AddCart(customerID, cardRequest)
    if err != nil {
        res := response.ClientResponse(http.StatusBadRequest, "fields provided are in wrong format" , nil, err.Error())
        ctx.JSON(http.StatusBadRequest, res)
        return
    }

    res := response.ClientResponse(http.StatusCreated, "Cart added", cart, nil)
    ctx.JSON(http.StatusCreated, res)
}


func (c *CartHanlder) ListCarts(ctx *gin.Context) {

    // get customerId from context
    customerID := ctx.MustGet("customerId").(uint)

    carts, err := c.cartService.ListCarts(customerID)
    if err != nil {
        res := response.ClientResponse(http.StatusBadRequest, "could not retrieve records" , nil, err.Error())
        ctx.JSON(http.StatusBadRequest, res)
        return
    }

    res := response.ClientResponse(http.StatusOK, "Success", carts, nil)
    ctx.JSON(http.StatusOK, res)
}

func (c *CartHanlder) DeleteCart(ctx *gin.Context) {
    productIdString := ctx.Param("product_id")
    customerID := ctx.MustGet("customerId").(uint)

    productID, err := strconv.ParseUint(productIdString, 10, 64)
    if err != nil {
        res := response.ClientResponse(http.StatusBadRequest, "Invalid product id", nil, err.Error())
        ctx.JSON(http.StatusBadRequest, res)
        return
    }

    err = c.cartService.DeleteCart(uint(productID), customerID)
    if err != nil {
        res := response.ClientResponse(http.StatusBadRequest, "fields provided are in wrong format" , nil, err.Error())
        ctx.JSON(http.StatusBadRequest, res)
        return
    }

    res := response.ClientResponse(http.StatusOK, "Cart deleted", nil, nil)
    ctx.JSON(http.StatusOK, res)
}
