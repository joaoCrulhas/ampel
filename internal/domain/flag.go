package domain

import "time"

type TargetRules struct {
	PercentageRollout int          `json:"percentage_rollout"` // Ex: 25%
	Departments       []Department `json:"departments"`        // Ex: ["RH", "TI"]
}

type Flag struct {
	ID                int       `json:"id"`
	Key               string    `json:"key"`
	Name              string    `json:"name"`
	Description       string    `json:"description"`
	Enabled           bool      `json:"enabled"`
	PercentageRollout int       `json:"percentage_rollout"`
	TargetDepartments []string  `json:"target_departments"`
	CreatedAt         time.Time `json:"created_at"`
	UpdatedAt         time.Time `json:"updated_at"`
}
