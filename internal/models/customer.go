package models

type Customer struct {
	ID        uint       `gorm:"primaryKey"`
	Name      string     `gorm:"not null"`
	Email     string     `gorm:"not null"`
	Password  string     `gorm:"not null"`
}