package types

// Common types used across the project

// Status represents an operational status
type Status string

const (
	// StatusOK indicates healthy status
	StatusOK Status = "ok"
	// StatusDegraded indicates degraded performance
	StatusDegraded Status = "degraded"
	// StatusDown indicates service is down
	StatusDown Status = "down"
)

// HealthCheck represents a health check result
type HealthCheck struct {
	Status  Status            `json:"status"`
	Version string            `json:"version,omitempty"`
	Checks  map[string]string `json:"checks,omitempty"`
}
