package db

import (
	"fmt"

	"github.com/nvlhnn/gommerce/internal/config"
	"github.com/nvlhnn/gommerce/internal/db/seeders"
	"github.com/nvlhnn/gommerce/internal/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectDatabase(cfg config.Config) (*gorm.DB, error) {
	psqlInfo := fmt.Sprintf("host=%s user=%s dbname=%s port=%s password=%s", cfg.DBHost, cfg.DBUser, cfg.DBName, cfg.DBPort, cfg.DBPassword)
	db, dbErr := gorm.Open(postgres.Open(psqlInfo), &gorm.Config{SkipDefaultTransaction: true})

	err := db.AutoMigrate(&models.Customer{},&models.Category{}, &models.Product{}, &models.Category{}, &models.Cart{}, &models.Transaction{}, &models.TransactionItem{})
	if  err != nil {
		return db, err
	}

	if cfg.IsSeeding {
		seeders.InitSeeders(db)	
	}

	return db, dbErr
}

