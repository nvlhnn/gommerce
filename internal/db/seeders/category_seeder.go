package seeders

import (
	"errors"
	"log"

	"github.com/nvlhnn/gommerce/internal/models"
	"gorm.io/gorm"
)

func SeedCategories(db *gorm.DB) {

	if db.Migrator().HasTable(&models.Category{}){
		if err := db.First(&models.Category{}).Error; errors.Is(err, gorm.ErrRecordNotFound) {
			days := []models.Category{
				{
					Name: "Electronics",
				},
				{
					Name: "Fashion",
				},
				{
					Name: "Health",
				},
			}
		
			err := db.Create(&days).Error
			if err != nil {
				log.Println("[seeding error] : ",err )
			}
		}
	}
}

