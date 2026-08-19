package dtos

import "time"

// USER DTOs
type CreateUserDTO struct {
	Name          string `json:"name" binding:"required"`
	Description   string `json:"description"`
	Active        bool   `json:"active"`
	DepartmentIDs []int  `json:"department_ids"`
}

type UpdateUserDTO struct {
	Name          *string `json:"name,omitempty"`
	Description   *string `json:"description,omitempty"`
	Active        *bool   `json:"active,omitempty"`
	DepartmentIDs []int   `json:"department_ids,omitempty"`
}

type UserResponseDTO struct {
	ID          int                     `json:"id"`
	Name        string                  `json:"name"`
	Description string                  `json:"description"`
	Active      bool                    `json:"active"`
	Departments []DepartmentResponseDTO `json:"departments"`
	CreatedAt   time.Time               `json:"created_at"`
	UpdatedAt   time.Time               `json:"updated_at"`
}
