package repository

import (
	"astrogo/internal/energy/model"
	"fmt"
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type EnergyRepository struct {
	db *gorm.DB
}

func NewEnergyRepository() (*EnergyRepository, error) {
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
	err = db.AutoMigrate(&model.EnergyStock{}, &model.EnergyTransaction{})
	if err != nil {
		log.Printf("Failed to migrate database: %v", err)
		return nil, err
	}

	// Initialize stock if not exists
	var count int64
	db.Model(&model.EnergyStock{}).Count(&count)
	if count == 0 {
		stock := &model.EnergyStock{
			Amount: 0,
		}
		if err := db.Create(stock).Error; err != nil {
			return nil, err
		}
	}

	return &EnergyRepository{db: db}, nil
}

func (r *EnergyRepository) GetStock() (*model.EnergyStock, error) {
	var stock model.EnergyStock
	if err := r.db.First(&stock).Error; err != nil {
		return nil, err
	}
	return &stock, nil
}

func (r *EnergyRepository) UpdateStock(amount float64) error {
	return r.db.Model(&model.EnergyStock{}).Where("id = ?", 1).Update("amount", amount).Error
}

func (r *EnergyRepository) ConsumeEnergy(amount float64) error {
	return r.db.Transaction(func(tx *gorm.DB) error {
		var stock model.EnergyStock
		if err := tx.First(&stock).Error; err != nil {
			return err
		}

		if stock.Amount < amount {
			return gorm.ErrRecordNotFound
		}

		stock.Amount -= amount
		return tx.Save(&stock).Error
	})
}

func (r *EnergyRepository) CreateTransaction(transaction *model.EnergyTransaction) error {
	return r.db.Create(transaction).Error
}

func (r *EnergyRepository) GetTransactions() ([]model.EnergyTransaction, error) {
	var transactions []model.EnergyTransaction
	if err := r.db.Find(&transactions).Error; err != nil {
		return nil, err
	}
	return transactions, nil
}

func (r *EnergyRepository) GetCurrentStock() (*model.EnergyStock, error) {
	var stock model.EnergyStock
	err := r.db.First(&stock).Error
	if err != nil {
		return nil, err
	}
	return &stock, nil
}

func (r *EnergyRepository) AddEnergy(amount float64, description string) error {
	if amount <= 0 {
		return nil
	}
	return r.UpdateStock(amount)
}

func (r *EnergyRepository) GetTransactionHistory() ([]model.EnergyTransaction, error) {
	var transactions []model.EnergyTransaction
	err := r.db.Order("created_at desc").Find(&transactions).Error
	return transactions, err
}

func (r *EnergyRepository) CheckAvailability(amount float64) (bool, error) {
	stock, err := r.GetCurrentStock()
	if err != nil {
		return false, err
	}
	return stock.Amount >= amount, nil
}
