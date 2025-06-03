package repository

import (
	"astrogo/internal/fuel/model"
	"fmt"
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type FuelRepository struct {
	db *gorm.DB
}

func NewFuelRepository() (*FuelRepository, error) {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_PORT"),
	)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Printf("Failed to connect to database: %v", err)
		return nil, err
	}

	// Auto migrate the schema
	err = db.AutoMigrate(&model.FuelStock{}, &model.FuelTransaction{})
	if err != nil {
		log.Printf("Failed to migrate database: %v", err)
		return nil, err
	}

	// Initialize stock if not exists
	var count int64
	db.Model(&model.FuelStock{}).Count(&count)
	if count == 0 {
		stock := &model.FuelStock{
			Amount: 0,
		}
		if err := db.Create(stock).Error; err != nil {
			return nil, err
		}
	}

	return &FuelRepository{db: db}, nil
}

func (r *FuelRepository) GetCurrentStock() (float64, error) {
	var stock model.FuelStock
	err := r.db.First(&stock).Error
	if err != nil {
		return 0, err
	}
	return stock.Amount, nil
}

func (r *FuelRepository) AddTransaction(transaction *model.FuelTransaction) error {
	return r.db.Transaction(func(tx *gorm.DB) error {
		// Create transaction record
		if err := tx.Create(transaction).Error; err != nil {
			return err
		}

		// Update stock
		var stock model.FuelStock
		if err := tx.First(&stock).Error; err != nil {
			return err
		}

		if transaction.Type == model.TransactionTypeAdd {
			stock.Amount += transaction.Amount
		} else if transaction.Type == model.TransactionTypeConsume {
			stock.Amount -= transaction.Amount
		}

		return tx.Save(&stock).Error
	})
}

func (r *FuelRepository) GetTransactionHistory() ([]model.FuelTransaction, error) {
	var transactions []model.FuelTransaction
	err := r.db.Order("timestamp desc").Find(&transactions).Error
	return transactions, err
}
