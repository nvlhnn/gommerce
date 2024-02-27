package models

type Cart struct {
	ID         uint    `gorm:"primaryKey"`
	CustomerID uint    `gorm:"uniqueIndex:idx_product_customer"`
	ProductID  uint    `gorm:"uniqueIndex:idx_product_customer"`
	Product    Product `gorm:"foreignKey:ProductID"`
	Quantity   uint
}