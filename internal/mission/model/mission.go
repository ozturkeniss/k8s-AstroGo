package model

type MissionType int32

const (
	MissionTypeExploration MissionType = iota
	MissionTypeResearch
	MissionTypeMaintenance
	MissionTypeTransport
)

type MissionStatus int32

const (
	MissionStatusPending MissionStatus = iota
	MissionStatusInProgress
	MissionStatusCompleted
	MissionStatusFailed
)

type Mission struct {
	ID            uint64        `gorm:"primaryKey" json:"id"`
	Name          string        `gorm:"not null" json:"name"`
	Description   string        `gorm:"not null" json:"description"`
	Type          MissionType   `gorm:"not null" json:"type"`
	Status        MissionStatus `gorm:"not null" json:"status"`
	EnergyReq     float64       `gorm:"not null" json:"energy_req"`
	CreatedAt     string        `gorm:"not null" json:"created_at"`
	UpdatedAt     string        `gorm:"not null" json:"updated_at"`
	StartedAt     string        `json:"started_at"`
	CompletedAt   string        `json:"completed_at"`
	FailedAt      string        `json:"failed_at"`
	FailureReason string        `json:"failure_reason"`
	AstronautID   uint64        `json:"astronaut_id"`
}

// MissionResource represents the resource requirements for a mission
type MissionResource struct {
	EnergyAmount float64
	FuelAmount   float64
}
