package seeders

import (
	"gorm.io/gorm"
)

// init all seeders
func InitSeeders(db *gorm.DB) {
	SeedCategories(db)
	SeedProducts(db)
}