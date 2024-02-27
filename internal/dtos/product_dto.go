package dtos

import "github.com/nvlhnn/gommerce/internal/models"

type ProductDTO struct {
	ID       uint   `json:"id"`
	Name     string `json:"name"`
	Price	float64 `json:"price"`
	Stock	uint    `json:"stock"`
	Category CategoryDTO `json:"category"`
}

func (dto *ProductDTO) ModelToDto(model models.Product) {
	dto.ID = model.ID
	dto.Name = model.Name
	dto.Price = model.Price
	dto.Stock = model.Stock
	dto.Category.ID = model.CategoryID
	dto.Category.Name = model.Category.Name
}