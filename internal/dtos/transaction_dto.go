package dtos

import (
	"github.com/nvlhnn/gommerce/internal/models"
)

type TransactionDTO struct {
	ID      			uint      			`json:"id"`
	TotalAmount 		float64   			`json:"total_amount"`
	TransactionDate 	string		`json:"transaction_date"`
	Items       		[]TransactionItemDTO `json:"items"`
}

type TransactionItemDTO struct {
	Product 	ProductDTO    `json:"product"`
	Quantity  	uint    `json:"quantity"`
	UnitPrice 	float64 `json:"unit_price"`
}

func (t *TransactionDTO) ModelToDto(transaction models.Transaction) {
	t.ID = transaction.ID
	t.TotalAmount = transaction.TotalAmount
	t.TransactionDate = transaction.TransactionDate.Format("2006-01-02 15:04:05")
	t.Items = make([]TransactionItemDTO, len(transaction.Items))
	for i, item := range transaction.Items {
		t.Items[i].ModelToDto(item)
	}
}

func (t *TransactionItemDTO) ModelToDto(item models.TransactionItem) {
	t.Product.ModelToDto(item.Product)
	t.Quantity = item.Quantity
	t.UnitPrice = item.UnitPrice
}
