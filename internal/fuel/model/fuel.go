package model

import (
	"time"
)

type TransactionType int32

const (
	TransactionTypeAdd     TransactionType = 1
	TransactionTypeConsume TransactionType = 2
)

type FuelStock struct {
	ID        uint64    `json:"id" gorm:"primaryKey"`
	Amount    float64   `json:"amount"` // Amount in liters
	UpdatedAt time.Time `json:"updated_at"`
}

// FuelTransaction represents fuel consumption or addition
type FuelTransaction struct {
	ID        uint64          `json:"id" gorm:"primaryKey"`
	Type      TransactionType `json:"type"`
	Amount    float64         `json:"amount"` // Positive for addition, negative for consumption
	Timestamp string          `json:"timestamp"`
}
