package kafka

// Event types
const (
	MissionCreatedEvent    = "mission.created"
	MissionStartedEvent    = "mission.started"
	MissionCompletedEvent  = "mission.completed"
	MissionFailedEvent     = "mission.failed"
	AstronautAssignedEvent = "astronaut.assigned"
	EnergyConsumedEvent    = "energy.consumed"
	FuelConsumedEvent      = "fuel.consumed"
)

// MissionCreated represents the event when a new mission is created
type MissionCreated struct {
	MissionID uint64  `json:"mission_id"`
	Name      string  `json:"name"`
	Type      int32   `json:"type"`
	EnergyReq float64 `json:"energy_req"`
	FuelReq   float64 `json:"fuel_req"`
	CreatedAt string  `json:"created_at"`
}

// MissionStarted represents the event when a mission starts
type MissionStarted struct {
	MissionID uint64 `json:"mission_id"`
	StartedAt string `json:"started_at"`
}

// MissionCompleted represents the event when a mission is completed
type MissionCompleted struct {
	MissionID   uint64 `json:"mission_id"`
	CompletedAt string `json:"completed_at"`
}

// MissionFailed represents the event when a mission fails
type MissionFailed struct {
	MissionID     uint64 `json:"mission_id"`
	FailedAt      string `json:"failed_at"`
	FailureReason string `json:"failure_reason"`
}

// AstronautAssigned represents the event when an astronaut is assigned to a mission
type AstronautAssigned struct {
	MissionID   uint64 `json:"mission_id"`
	AstronautID uint64 `json:"astronaut_id"`
	AssignedAt  string `json:"assigned_at"`
}

// EnergyConsumed represents the event when energy is consumed
type EnergyConsumed struct {
	MissionID  uint64  `json:"mission_id"`
	Amount     float64 `json:"amount"`
	ConsumedAt string  `json:"consumed_at"`
}

// FuelConsumed represents the event when fuel is consumed
type FuelConsumed struct {
	MissionID  uint64  `json:"mission_id"`
	Amount     float64 `json:"amount"`
	ConsumedAt string  `json:"consumed_at"`
}
