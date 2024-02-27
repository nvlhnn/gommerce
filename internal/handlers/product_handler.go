package handlers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/nvlhnn/gommerce/internal/services"
	"github.com/nvlhnn/gommerce/internal/utils/response"
)

type ProductHandler interface {
	List(ctx *gin.Context)
}

type productHandler struct {
	productService services.ProductServiceInterface
}

func NewProductHandler(productService services.ProductServiceInterface) ProductHandler {
	return &productHandler{productService: productService}
}

func (h *productHandler) List(ctx *gin.Context) {

	// get category id
	categoryID := ctx.Param("category_id")
	categoryIDUint, err := strconv.ParseUint(categoryID, 10, 64)
	if err != nil {
        res := response.ClientResponse(http.StatusBadRequest, "Invalid request", nil, err.Error())
        ctx.JSON(http.StatusBadRequest, res)
		return
	}

	// get query string
	query := ctx.Request.URL.Query()


	// get products by category
	products, metadata, err := h.productService.GetProductsByCategory(uint(categoryIDUint), query)

	if err != nil {
        res := response.ClientResponse(http.StatusInternalServerError, "Internal server error" , nil, err.Error())
        ctx.JSON(http.StatusInternalServerError, res)
		return
	}

	// response
	res := response.ClientResponse(http.StatusOK, "Success", gin.H{
		"products": products,
		"metadata": metadata,
	}, nil)
	ctx.JSON(http.StatusOK, res)
	
}


	




