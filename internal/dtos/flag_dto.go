package dtos

import "time"

type CreateFlagDTO struct {
	Key               string   `json:"key" binding:"required"`
	Name              string   `json:"name" binding:"required"`
	Description       string   `json:"description"`
	Enabled           bool     `json:"enabled"`
	PercentageRollout int      `json:"percentage_rollout" binding:"gte=0,lte=100"`
	TargetDepartments []string `json:"target_departments"`
}

type UpdateFlagDTO struct {
	Name              *string  `json:"name,omitempty"`
	Description       *string  `json:"description,omitempty"`
	Enabled           *bool    `json:"enabled,omitempty"`
	PercentageRollout *int     `json:"percentage_rollout,omitempty" binding:"omitempty,gte=0,lte=100"`
	TargetDepartments []string `json:"target_departments,omitempty"`
}

type FlagResponseDTO struct {
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
