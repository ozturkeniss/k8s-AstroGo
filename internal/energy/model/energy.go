package model

type TransactionType int32

const (
	TransactionTypeAddition    TransactionType = 0 // TRANSACTION_TYPE_ADDITION
	TransactionTypeConsumption TransactionType = 1 // TRANSACTION_TYPE_CONSUMPTION
)

type EnergyStock struct {
	ID     uint64  `gorm:"primaryKey"`
	Amount float64 `gorm:"not null"` // kWh
}

type EnergyTransaction struct {
	ID          uint64          `gorm:"primaryKey"`
	Amount      float64         `gorm:"not null"` // kWh
	Type        TransactionType `gorm:"not null"`
	Description string          `gorm:"not null"`
	CreatedAt   string          `gorm:"not null"`
}
