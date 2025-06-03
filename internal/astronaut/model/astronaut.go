package model

import (
	"time"
)

type AvailabilityStatus string

const (
	Available   AvailabilityStatus = "AVAILABLE"
	Unavailable AvailabilityStatus = "UNAVAILABLE"
)

type Astronaut struct {
	ID               uint64             `json:"id" gorm:"primaryKey"`
	Name             string             `json:"name"`
	Status           AvailabilityStatus `json:"status"`
	CurrentMissionID *uint64            `json:"current_mission_id,omitempty"`
	CreatedAt        time.Time          `json:"created_at"`
	UpdatedAt        time.Time          `json:"updated_at"`
}
