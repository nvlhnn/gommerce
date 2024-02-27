package seeders

import (
	"errors"
	"log"

	"github.com/nvlhnn/gommerce/internal/models"
	"gorm.io/gorm"
)

// seed prducts
func SeedProducts(db *gorm.DB) {

	if db.Migrator().HasTable(&models.Product{}){
		if err := db.First(&models.Product{}).Error; errors.Is(err, gorm.ErrRecordNotFound) {
			days := []models.Product{
				{
					Name:        "Iphone 12",
					Price:       12000000,
					Stock:       10,
					CategoryID:  1,
				},
				{
					Name:        "Macbook Pro",
					Price:       30000000,
					Stock:       5,
					CategoryID:  1,
				},
				{
					Name:        "T-shirt",
					Price:       200000,
					Stock:       20,
					CategoryID:  2,
				},
				{
					Name:        "Jeans",
					Price:       300000,
					Stock:       15,
					CategoryID:  2,
				},
				{
					Name:        "Sneakers",
					Price:       500000,
					Stock:       10,
					CategoryID:  2,
				},
				{
					Name:        "Iphone 11",
					Price:       10000000,
					Stock:       10,
					CategoryID:  1,
				},
				{
					Name:        "Iphone 11 Pro",
					Price:       15000000,
					Stock:       10,
					CategoryID:  1,
				},
				{
					Name:        "Vitamin C",
					Price:       50000,
					Stock:       50,
					CategoryID:  3,
				},

				{
					Name:        "Vitamin D",
					Price:       60000,
					Stock:       50,
					CategoryID:  3,
				},
			}
		
		
			err := db.Create(&days).Error
			if err != nil {
				log.Println("[seeding error] : ",err )
			}
		}
	}

}