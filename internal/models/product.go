package models

type Product struct {
	ID         uint     `gorm:"primaryKey"`
	Name       string   `gorm:"not null"`
	Price      float64  `gorm:"not null"`
	Stock      uint     `gorm:"not null"`
	CategoryID uint     `gorm:"not null"`
	Category   Category `gorm:"foreignKey:CategoryID"`
}
