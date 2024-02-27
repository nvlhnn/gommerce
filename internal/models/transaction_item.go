package models

type TransactionItem struct {
	ID            uint    `gorm:"primaryKey"`
	TransactionID uint    `gorm:"not null"`
	ProductID     uint    `gorm:"not null"`
	Quantity      uint     `gorm:"not null"`
	UnitPrice     float64 `gorm:"not null"`
	Product       Product `gorm:"foreignKey:ProductID"`
}