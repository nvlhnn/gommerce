package dtos

import "github.com/nvlhnn/gommerce/internal/models"

type AddCartRequest struct {

	ProductID uint `json:"product_id" form:"product_id" binding:"required,number"`
	Quantity uint `json:"quantity" form:"quantity" binding:"required,min=1,number"`
}

type CartDTO struct {
	ID       uint `json:"id"`
	Product  ProductDTO `json:"product"`
	Quantity uint `json:"quantity"`
}

// card model to dto
func (c *CartDTO) ModelToDto(cart models.Cart) {
	c.ID = cart.ID
	c.Product.ModelToDto(cart.Product)
	c.Quantity = cart.Quantity

}
