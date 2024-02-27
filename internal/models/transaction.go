package models

import (
	"time"
)

type Transaction struct {
	ID      			uint      			`gorm:"primaryKey"`
	CustomerID  		uint      			`gorm:"not null"`
	Customer    		Customer  			`gorm:"foreignKey:CustomerID"`
	TotalAmount 		float64   			`gorm:"not null"`
	TransactionDate 	time.Time 			`gorm:"not null"`
	Items       		[]TransactionItem
}